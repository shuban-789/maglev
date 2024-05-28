package main

import "fmt"
import "net"
import "os"
import "os/exec"
import "os/user"
import "strings"
import "bufio"

func handleError(err error) int {
	if err != nil {
		return 1
	}
	return 0
}

func spawnShell(conn net.Conn) {
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
			} else {
				dir, _ = os.Getwd()
			}
			continue
		}

		dir, err = os.Getwd()
		if handleError(err) == 1 {
			fmt.Printf("🔴 [ERROR] Could not update directory: %v\n", err)
		}

		cmd := exec.Command("/bin/bash", "-c", command)
		cmd.Dir = dir
		cmd.Stdout = conn
		cmd.Stderr = conn
		if err := cmd.Run(); handleError(err) == 1 {
			fmt.Fprintf(conn, "🔴 [ERROR] Unable to execute commands: %v\n", err)
		}
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
		go spawnShell(conn)
	}
}

func connect(IP string, PORT string) {
    conn, err := net.Dial("tcp", IP+":"+PORT)
    if handleError(err) == 1 {
        fmt.Printf("🔴 [ERROR] Unable to connect to %v on port %v: %v\n", IP, PORT, err)
        return
    }
    defer conn.Close()

    reader := bufio.NewReader(os.Stdin)
    serverReader := bufio.NewReader(conn)

    for {
        prompt, err := serverReader.ReadString('$')
        if handleError(err) == 1 {
            fmt.Printf("🔴 [ERROR] Could not read server prompt: %v\n", err)
            return
        }

		prompt += string(' ')
        fmt.Print(prompt)

        input, err := reader.ReadString('\n')
        if handleError(err) == 1 {
            fmt.Printf("🔴 [ERROR] Could not read input: %v\n", err)
            continue
        }

        input = strings.TrimSpace(input)
        if input == "exit" {
            return
        }

        _, err = conn.Write([]byte(input + "\n"))
        if handleError(err) == 1 {
            fmt.Printf("🔴 [ERROR] Could not send command: %v\n", err)
            continue
        }
    }
}

func help() {
	fmt.Printf("Usage: ./bluefox [OPTION] [ARGUMENT]\n")
	fmt.Printf("\nOptions:\n")
	fmt.Printf("	-h, Shows help menu for this command\n")
	fmt.Printf("	-l, Sets up listener for a specified port\n")
	fmt.Printf("	-c, Connects to a server based on a specified address and port\n")
	fmt.Printf("\nFormat:\n")
	fmt.Printf("	./bluefox -h\n")
	fmt.Printf("	./bluefox -l <PORT>\n")
	fmt.Printf("	./bluefox -c <IP> <PORT>\n")
	fmt.Printf("\nExamples:\n")
	fmt.Printf("	./bluefox -l 1234\n")
	fmt.Printf("	./bluefox -c 127.0.0.1 1234\n")
}

func main() {
	if len(os.Args) > 1 {
		if strings.Compare(os.Args[1], "-l") == 0 {
			listen(os.Args[2])
		} else if strings.Compare(os.Args[1], "-c") == 0 {
			var ipAddr string
			if strings.Compare(os.Args[2], "localhost") == 0 {
				ipAddr = "127.0.0.1"
			} else {
				ipAddr = os.Args[2]
			}
			Port := os.Args[3]
			connect(ipAddr, Port)
		} else if strings.Compare(os.Args[1], "-h") == 0 {
			help()
		}
	} else {
		help()
	}
}
