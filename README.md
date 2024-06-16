# Usage

```
Usage: ./gocat [OPTION1] [ARGUMENT1] ... [OPTIONn] [ARGUMENTn]

Options:
        -h, Shows help menu for this command
        -l, Sets up listener for a specified port
                --shell, spawns a specified shell supporting the -c argument
        -c, Connects to a device based on a specified address and port
                --payload, spawns a specified shell supporting the -c argument

Format:
        ./gocat -h
        ./gocat -l <PORT>
        ./gocat -l <PORT> --shell <SHELL>
        ./gocat -c <IP> <PORT>

Examples:
        ./gocat -l 1234
        ./gocat -l 1234 --shell /usr/bin/python3
        ./gocat -c 127.0.0.1 1234
        ./gocat -c 127.0.0.1 1234 --payload "ls -l"
```
