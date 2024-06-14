package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"bufio"
	"bytes"
	"strconv"
)

func handleError(err error) int {
	if err != nil {
		return 1
	}
	return 0
}

func spawnShell(conn net.Conn, shell string) {
	addrInfo := strings.Split(conn.RemoteAddr().String(), ":")
	ip := addrInfo[0]

	defer conn.Close()
	currentUser, err := user.Current()
	if handleError(err) == 1 {
		fmt.Fprintf(conn, "🔴 [ERROR] Unable to get current user: %v\n", err)
		return
	}

	username := currentUser.Username
	hostname, err := os.Hostname()

	if handleError(err) == 1 {
		return
	}

	fmt.Printf("🟢 [SUCCESS] Received connection from %v\n", ip)
	conn.Write([]byte("🦊 Connection established!\n"))
	conn.Write([]byte("⚙️ SHELL: " + shell + "\n"))
	conn.Write([]byte("⚙️ USER: " + username + "\n"))
	conn.Write([]byte("⚙️ HOSTNAME: " + hostname + "\n"))

	dir, err := os.Getwd()
	if handleError(err) == 1 {
		fmt.Fprintf(conn, "🔴 [ERROR] Unable to get current directory: %v\n", err)
		return
	}

	for {
		prompt := fmt.Sprintf("%s@%s:%s$ ", username, hostname, dir)
		conn.Write([]byte(prompt))
		input := make([]byte, 1024)
		n, err := conn.Read(input)

		if handleError(err) == 1 {
			fmt.Printf("🔴 [ERROR] Could not read input from client: %v\n", err)
			return
		}

		command := strings.TrimSpace(string(input[:n]))

		if command == "exit" {
			conn.Write([]byte("👋 Bye!\n"))
			fmt.Printf("🟢 [SUCCESS] Connection from %v successfully closed\n", ip)
			return
		}

		if strings.HasPrefix(command, "cd ") {
			path := strings.TrimSpace(command[3:])
			err := os.Chdir(path)
			if handleError(err) == 1 {
				fmt.Fprintf(conn, "🔴 [ERROR] Unable to change directory: %v\n", err)
				fmt.Printf("🔴 [ERROR] Client is unnable to change directory: %v\n", err)
			} else {
				dir, _ = os.Getwd()
			}
			continue
		}

		dir, err = os.Getwd()
		if handleError(err) == 1 {
			fmt.Printf("🔴 [ERROR] Could not update directory: %v\n", err)
		}

		cmd := exec.Command(shell, "-c", command)
		cmd.Dir = dir
		cmd.Stdout = conn
		cmd.Stderr = conn
		if err := cmd.Run(); handleError(err) == 1 {
			fmt.Fprintf(conn, "🔴 [ERROR] Unable to execute commands: %v\n", err)
			fmt.Printf("🔴 [ERROR] Client is unnable to execute commands: %v\n", err)
		}
	}
}

func spawnComm(conn net.Conn) {
	addrInfo := strings.Split(conn.RemoteAddr().String(), ":")
	ip := addrInfo[0]
	defer conn.Close()

	fmt.Printf("🟢 [SUCCESS] Received connection from %v\n", ip)
	conn.Write([]byte("🦊 Connection established!\n"))

	for {
		input := make([]byte, 1024)
		n, err := conn.Read(input)
		if handleError(err) == 1 {
			fmt.Printf("🔴 [ERROR] Could not read input from client: %v\n", err)
			return
		}

		msg := strings.TrimSpace(string(input[:n]))
		fmt.Printf(msg + "\n")

		if msg == "exit" {
			conn.Write([]byte("👋 Bye!\n"))
			fmt.Printf("🟢 [SUCCESS] Connection from %v successfully closed\n", ip)
			return
		}
	}
}

func listenShell(PORT string, SHELL string) {
	ln, err := net.Listen("tcp", ":"+PORT)
	if handleError(err) == 1 {
		fmt.Printf("🔴 [ERROR] Unable to listen on specified port: %v\n", err)
		return
	} else {
		fmt.Printf("🟡 [IDLE] Listening on port %s\n", PORT)
	}

	for {
		conn, err := ln.Accept()
		if handleError(err) == 1 {
			fmt.Printf("🔴 [ERROR] Unable to establish connection: %v\n", err)
		} else {
			fmt.Printf("🟢 [SUCCESS] Connection established\n")
		}
		go spawnShell(conn, SHELL)
	}
}

func listen(PORT string) {
	ln, err := net.Listen("tcp", ":"+PORT)
	if handleError(err) == 1 {
		fmt.Printf("🔴 [ERROR] Unable to listen on specified port: %v\n", err)
		return
	} else {
		fmt.Printf("🟡 [IDLE] Listening on port %s\n", PORT)
	}

	for {
		conn, err := ln.Accept()
		if handleError(err) == 1 {
			fmt.Printf("🔴 [ERROR] Unable to establish connection: %v\n", err)
		} else {
			fmt.Printf("🟢 [SUCCESS] Connection established\n")
		}
		go spawnComm(conn)
	}
}

func connectPayloadFile(IP string, PORT string) {
	conn, err := net.Dial("tcp", IP+":"+PORT)
	if err != nil {
		fmt.Printf("🔴 [ERROR] Unable to connect to %v on port %v: %v\n", IP, PORT, err)
		return
	}
	defer conn.Close()

	fmt.Printf("🟢 [CONNECTED] Connected to %v on port %v\n", IP, PORT)

	go func() {
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("🔴 [ERROR] Failed to read message: %v\n", err)
				return
			}
			fmt.Printf("🟡 [RECEIVED] %s", message)
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("🔴 [ERROR] Could not read input: %v\n", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "exit" {
			fmt.Printf("🟢 [SUCCESS] Client has disconnected")
			return
		}

		_, err = conn.Write([]byte(input + "\n"))
		if err != nil {
			fmt.Printf("🔴 [ERROR] Could not send input: %v\n", err)
			continue
		}
	}
}

func connectPayload(IP string, PORT string) {
	conn, err := net.Dial("tcp", IP+":"+PORT)
	if err != nil {
		fmt.Printf("🔴 [ERROR] Unable to connect to %v on port %v: %v\n", IP, PORT, err)
		return
	}
	defer conn.Close()

	fmt.Printf("🟢 [CONNECTED] Connected to %v on port %v\n", IP, PORT)

	go func() {
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("🔴 [ERROR] Failed to read message: %v\n", err)
				return
			}
			fmt.Printf("🟡 [RECEIVED] %s", message)
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("🔴 [ERROR] Could not read input: %v\n", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "exit" {
			fmt.Printf("🟢 [SUCCESS] Client has disconnected")
			return
		}

		_, err = conn.Write([]byte(input + "\n"))
		if err != nil {
			fmt.Printf("🔴 [ERROR] Could not send input: %v\n", err)
			continue
		}
	}
}

func connect(IP string, PORT string) {
	conn, err := net.Dial("tcp", IP+":"+PORT)
	if err != nil {
		fmt.Printf("🔴 [ERROR] Unable to connect to %v on port %v: %v\n", IP, PORT, err)
		return
	}
	defer conn.Close()

	fmt.Printf("🟢 [CONNECTED] Connected to %v on port %v\n", IP, PORT)

	go func() {
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("🔴 [ERROR] Failed to read message: %v\n", err)
				return
			}
			fmt.Printf("🟡 [RECEIVED] %s", message)
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("🔴 [ERROR] Could not read input: %v\n", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "exit" {
			fmt.Printf("🟢 [SUCCESS] Client has disconnected")
			return
		}

		_, err = conn.Write([]byte(input + "\n"))
		if err != nil {
			fmt.Printf("🔴 [ERROR] Could not send input: %v\n", err)
			continue
		}
	}
}

func help() {
	fmt.Printf("Usage: ./bluefox [OPTION1] [ARGUMENT1] ... [OPTIONn] [ARGUMENTn]\n")
	fmt.Printf("\nOptions:\n")
	fmt.Printf("	-h, Shows help menu for this command\n")
	fmt.Printf("	-l, Sets up listener for a specified port\n")
	fmt.Printf("		--shell, spawns a specified shell supporting the -c argument\n")
	fmt.Printf("	-c, Connects to a device based on a specified address and port\n")
	fmt.Printf("\nFormat:\n")
	fmt.Printf("	./bluefox -h\n")
	fmt.Printf("	./bluefox -l <PORT>\n")
	fmt.Printf("	./bluefox -l <PORT> --shell <SHELL>\n")
	fmt.Printf("	./bluefox -c <IP> <PORT>\n")
	fmt.Printf("\nExamples:\n")
	fmt.Printf("	./bluefox -l 1234\n")
	fmt.Printf("	./bluefox -l 1234 --shell /usr/bin/python3\n")
	fmt.Printf("	./bluefox -c 127.0.0.1 1234\n")
}

func main() {
	if len(os.Args) > 1 {
		if strings.Compare(os.Args[1], "-l") == 0 {
			if len(os.Args) > 3 && strings.Compare(os.Args[3], "--shell") == 0 {
				shell := os.Args[4]
				listenShell(os.Args[2], shell)
			} else {
				listen(os.Args[2])
			}
		} else if len(os.Args) > 2 && strings.Compare(os.Args[1], "-c") == 0 {
			if len(os.Args) > 3 && strings.Compare(os.Args[3], "--payload") == 0 {
				shell := os.Args[4]
				listenShell(os.Args[2], shell)
			} else {
				var ipAddr string
				if strings.Compare(os.Args[2], "localhost") == 0 {
					ipAddr = "127.0.0.1"
				} else {
					ipAddr = os.Args[2]
				}
				Port := os.Args[3]
				connect(ipAddr, Port)
			}
		} else if strings.Compare(os.Args[1], "-h") == 0 {
			help()
		}
	} else {
		help()
	}
}
