# Usage

```
Usage: ./bluefox [OPTION1] [ARGUMENT1] ... [OPTIONn] [ARGUMENTn]

Options:
        -h, Shows help menu for this command
        -l, Sets up listener for a specified port
                --shell, spawns a specified shell supporting the -c argument
        -c, Connects to a server based on a specified address and port

Format:
        ./bluefox -h
        ./bluefox -l <PORT>
        ./bluefox -l <PORT> --shell <SHELL>
        ./bluefox -c <IP> <PORT>

Examples:
        ./bluefox -l 1234
        ./bluefox -l 1234 --shell /usr/bin/python3
        ./bluefox -c 127.0.0.1 1234
```
