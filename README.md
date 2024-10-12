# Maglev [![Go Report Card](https://goreportcard.com/badge/github.com/shuban-789/maglev)](https://goreportcard.com/report/github.com/shuban-789/maglev)

Maglev (`maglev`) is a versatile remote connection Swiss Army Knife programmed in Go. Maglev can currently successfully act as a client and a listener, accept the transfer of commands as payload, and support TLS connections. This project is still in development for more features.

# Usage
```
Usage: ./maglev [OPTION1] [ARGUMENT1] ... [OPTIONn] [ARGUMENTn]

Options:
        -h, Shows help menu for this command
        -l, Sets up listener for a specified port
                --shell, spawns a specified shell supporting the -c argument
                --tls, use Transport Layer Security (TLS) protection
        -c, Connects to a device based on a specified address and port
                --payload, spawns a specified shell supporting the -c argument

Format:
        ./maglev -h
        ./maglev -l <PORT> --tls <KEY> <CERT>
        ./maglev -l <PORT> --shell <SHELL>
        ./maglev -l <PORT> --shell <SHELL> --tls <KEY> <CERT>
        ./maglev -c <IP> <PORT>

Examples:
        ./maglev -l 1234
        ./maglev -l 1234 --shell /usr/bin/python3
        ./maglev -c 127.0.0.1 1234
        ./maglev -c 127.0.0.1 1234 --payload "whoami"
```
