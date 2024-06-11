# Usage

```
Usage: ./bluefox [OPTION1] [ARGUMENT1] ... [OPTIONn] [ARGUMENTn]

Options:
        -h, Shows help menu for this command
        -l, Sets up listener for a specified port
                --shell, Spawns a specified shell
        -c, Connects to a server based on a specified address and port
        -s, List all nll online processes cnnected to the network
                --proto, Outputs the specified protocol

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

# What to Expect

- More features
- Formatting improvements in source code
