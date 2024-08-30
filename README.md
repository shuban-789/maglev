# Usage

```
Usage: ./bluefox [OPTION1] [ARGUMENT1] ... [OPTIONn] [ARGUMENTn]

Options:
        -h, Shows help menu for this command
        -l, Sets up listener for a specified port
                --shell, spawns a specified shell supporting the -c argument
                --tls, use Transport Layer Security (TLS) protection
        -c, Connects to a device based on a specified address and port
                --payload, spawns a specified shell supporting the -c argument

Format:
        ./bluefox -h
        ./bluefox -l <PORT> --tls <KEY> <CERT>
        ./bluefox -l <PORT> --shell <SHELL>
        ./bluefox -l <PORT> --shell <SHELL> --tls <KEY> <CERT>
        ./bluefox -c <IP> <PORT>

Examples:
        ./bluefox -l 1234
        ./bluefox -l 1234 --shell /usr/bin/python3
        ./bluefox -c 127.0.0.1 1234
        ./bluefox -c 127.0.0.1 1234 --payload "ls -l"
```
