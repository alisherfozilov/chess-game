package main

import (
	"fmt"
  
	"github.com/alisherfozilov/chess-game/chess"
	"github.com/alisherfozilov/chess-game/cmd"
	menuPack "github.com/alisherfozilov/chess-game/menu"
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
