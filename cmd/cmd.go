package cmd

import (
	"os"
	"os/exec"
)

func Clear() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}
