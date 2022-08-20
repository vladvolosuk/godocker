package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var args []string

	wslExePath, _ := exec.LookPath("wsl")

	args = append(args, wslExePath, "docker")
	args = append(args, os.Args[1:]...)

	wslCmd := &exec.Cmd{
		Path:   wslExePath,
		Args:   args,
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	if err := wslCmd.Run(); err != nil {
		fmt.Println("Error:", err)
	}

}
