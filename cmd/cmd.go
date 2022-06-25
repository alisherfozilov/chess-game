package cmd

import (
	"os"
	"os/exec"
	"runtime"
)

func Clear() {
	cmd := getCmd()
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

func getCmd() *exec.Cmd {
	switch runtime.GOOS {
	case "windows":
		return exec.Command("cmd", "/c", "cls")
	default:
		return exec.Command("clear")
	}
}
