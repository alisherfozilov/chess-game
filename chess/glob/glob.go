package glob

import "github.com/alisherfozilov/chess-game/chess/glob/Xy"

var Letters = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'}
var Numbers = []int{8, 7, 6, 5, 4, 3, 2, 1}
var Black = []int{1, 2, 3, 5, 10, 99}
var White = []int{11, 21, 31, 51, 101, 990}
var Selection = Xy.Xy{-1, -1}

const SIZE = 8

var Table = [SIZE][SIZE]int{
	{51, 21, 31, 101, 990, 31, 21, 51},
	{11, 11, 11, 11, 11, 11, 11, 11},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{1, 1, 1, 1, 1, 1, 1, 1},
	{5, 2, 3, 10, 99, 3, 2, 5},
}

var TableStandartPosition = [SIZE][SIZE]int{
	{51, 21, 31, 101, 990, 31, 21, 51},
	{11, 11, 11, 11, 11, 11, 11, 11},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{1, 1, 1, 1, 1, 1, 1, 1},
	{5, 2, 3, 10, 99, 3, 2, 5},
}

var Counter = 0

type ComputerMove struct {
	FromX int
	FromY int
	ToX   int
	ToY   int
}

var Magic = false
var BadMagic = false
