package main

import (
	"fmt"

	"temp/chess"
	"temp/cmd"
	menuPack "temp/menu"
)

type Menu = menuPack.Menu

func main() {
	flag := true
	var menu Menu

	for flag {
		cmd.Clear()
		fmt.Println()
		chess.Label()
		menu.Print()
		menu.InputChoice()
		flag = menu.SwitchChoice()
	}
}
