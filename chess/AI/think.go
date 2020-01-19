package AI

import (
	"fmt"
	"log"

	"temp/chess"
	gb "temp/chess/glob"
	XyPack "temp/chess/glob/Xy"
	"temp/chess/glob/stackXy"
)

type stack = stackXy.Stack
type Xy = XyPack.Xy

func Think(recur int) gb.ComputerMove { // such like this!
	var bestMove gb.ComputerMove
	var whiteSide bool = true
	var isFirst bool = true

	helpThink(recur, whiteSide, &bestMove, 0.0, isFirst)
	//fmt.Println("Price: ", price)
	return bestMove
}
func helpThink(recStep int, whiteSide bool, bestMove *gb.ComputerMove, tempScore float64, isFirst bool) float64 {

	var finalScore float64 = 0.0
	var moves stack
	var figures []Xy
	var score, bestScore float64
	var isLast = true

	if whiteSide {
		bestScore = -1000000.0
	} else {
		bestScore = 1000000.0
	}

	chess.Scan(&figures, whiteSide) // get figures for this side

	for i := range figures { // for every figure
		chess.GetFiguresMoves(figures[i], &moves)

		for !moves.Empty() { // for every move
			move, _ := moves.Top() // get move
			_ = moves.Pop()// delete figure

			// make move
			whoYouKilled := gb.Table[move.Y][move.X]
			gb.Table[move.Y][move.X] = gb.Table[figures[i].Y][figures[i].X]
			gb.Table[figures[i].Y][figures[i].X] = 0

			score = rank(whoYouKilled) + positionScore(move, whiteSide)
			score += tempScore

			if recStep > 0 {
				isLast = false
				finalScore = helpThink(recStep - 1, !whiteSide, bestMove, score,false)
			}

			if !isLast {
				score = finalScore
			}

			gb.Table[figures[i].Y][figures[i].X] = gb.Table[move.Y][move.X]
			gb.Table[move.Y][move.X] = whoYouKilled
			/*{
				printTable()
				fmt.Println()
				fmt.Println("recurStep : ", recStep)
				fmt.Println("score : ", score)
				fmt.Println("bestScore : ", bestScore)
				fmt.Println("figures[i].Y, figures[i].X :  move.Y, move.X")
				fmt.Println(figures[i].Y, figures[i].X, " : ", move.Y, move.X)
				fmt.Println()
			}*/
			if compare(bestScore, score, whiteSide) {
				bestScore = score
				//fmt.Println("recStep: ", recStep,"BestScore: ", bestScore)
				if isFirst {
					bestMove.FromY = figures[i].Y
					bestMove.FromX = figures[i].X
					bestMove.ToY = move.Y
					bestMove.ToX = move.X
				}
			}
		}

	}
	return bestScore
}
func compare(bestScore, score float64, whiteSide bool) bool {
	if whiteSide {
		return bestScore < score
	} else {
		return bestScore > score
	}
}

func rank(fig int) float64 { // thanks to my stupidness this func returns rank only for white side :D
	switch fig {
	case 0:
		return 0;

	case 11:
		return -1;
	case 1:
		return 1;

	case 21:
		return -2.7;
	case 2:
		return 2.7;

	case 31:
		return -3;
	case 3:
		return 3;

	case 51:
		return -5;
	case 5:
		return 5;

	case 101:
		return -13;
	case 10:
		return 13;

	case 990:
		return -9990;
	case 99:
		return 1000;

	default:
		fmt.Println("RANKer")
	}
	return -9876543.21
}
func positionScore(p Xy, whiteColor bool) float64 {
	var st stack
	chess.GetFiguresMoves(p, &st)
	var temp Xy
	var answer float64 = 0.0

	for !st.Empty() {
		temp, _ = st.Top()
		_ = st.Pop()
		switch gb.Table[temp.Y][temp.X] {
		case 0:
			answer += 0;

		case 11:
			answer += -0.00001;
		case 1:
			answer += 0.00001;

		case 21:
			answer += -0.00003;
		case 2:
			answer += 0.00003;

		case 31:
			answer += -0.00003;
		case 3:
			answer += 0.00003;

		case 51:
			answer += -0.00005;
		case 5:
			answer += 0.00005;

		case 101:
			answer += -0.0001;
		case 10:
			answer += 0.0001;

		case 990:
			answer += -0.001;
		case 99:
			answer += 0.001;

		default:
			log.Fatal("rankER_POSIT")
		}
	}
	//answer += isChecked(p)
	if whiteColor {
		answer += cofficientWhite[p.Y][p.X]
	} else {
		answer -= cofficientBlack[p.Y][p.X]
	}

	return answer
}
/*var cofficientWhite = [][]float64 {
	{0.071, 0.071, 0.071, 0.071, 0.071, 0.071, 0.071, 0.071},
	{0.061, 0.062, 0.062, 0.062, 0.062, 0.062, 0.062, 0.061},
	{0.051, 0.052, 0.053, 0.053, 0.053, 0.053, 0.052, 0.051},
	{0.041, 0.042, 0.043, 0.044, 0.044, 0.043, 0.042, 0.041},
	{0.031, 0.032, 0.033, 0.033, 0.033, 0.033, 0.032, 0.031},
	{0.021, 0.022, 0.023, 0.023, 0.023, 0.023, 0.022, 0.021},
	{0.011, 0.012, 0.012, 0.012, 0.012, 0.012, 0.012, 0.012},
	{0.001, 0.001, 0.001, 0.001, 0.001, 0.001, 0.001, 0.001},
}
var cofficientBlack = [][]float64 {
	{0.001, 0.001, 0.001, 0.001, 0.001, 0.001, 0.001, 0.001},
	{0.011, 0.012, 0.012, 0.012, 0.012, 0.012, 0.012, 0.011},
	{0.021, 0.022, 0.023, 0.023, 0.023, 0.023, 0.022, 0.021},
	{0.031, 0.032, 0.033, 0.033, 0.033, 0.033, 0.032, 0.031},
	{0.041, 0.042, 0.043, 0.044, 0.044, 0.043, 0.042, 0.041},
	{0.051, 0.052, 0.053, 0.053, 0.053, 0.053, 0.052, 0.051},
	{0.061, 0.062, 0.062, 0.062, 0.062, 0.062, 0.062, 0.061},
	{0.071, 0.071, 0.071, 0.071, 0.071, 0.071, 0.071, 0.071},
}*/
var cofficientWhite = [][]float64 {
	{0.0000071, 0.0000071, 0.0000071, 0.0000071, 0.0000071, 0.0000071, 0.0000071, 0.0000071},
	{0.0000061, 0.0000062, 0.0000062, 0.0000062, 0.0000062, 0.0000062, 0.0000062, 0.0000061},
	{0.0000051, 0.0000052, 0.0000053, 0.0000053, 0.0000053, 0.0000053, 0.0000052, 0.0000051},
	{0.0000041, 0.0000042, 0.0000043, 0.0000044, 0.0000044, 0.0000043, 0.0000042, 0.0000041},
	{0.0000031, 0.0000032, 0.0000033, 0.0000034, 0.0000034, 0.0000033, 0.0000032, 0.0000031},
	{0.0000021, 0.0000022, 0.0000023, 0.0000023, 0.0000023, 0.0000023, 0.0000022, 0.0000021},
	{0.0000011, 0.0000012, 0.0000012, 0.0000012, 0.0000012, 0.0000012, 0.0000012, 0.0000011},
	{0.0000001, 0.0000001, 0.0000001, 0.0000001, 0.0000001, 0.0000001, 0.0000001, 0.0000001},
}
var cofficientBlack = [][]float64 {
	{0.0000001, 0.0000001, 0.0000001, 0.0000001, 0.0000001, 0.0000001, 0.0000001, 0.0000001},
	{0.0000011, 0.0000012, 0.0000012, 0.0000012, 0.0000012, 0.0000012, 0.0000012, 0.0000011},
	{0.0000021, 0.0000022, 0.0000023, 0.0000023, 0.0000023, 0.0000023, 0.0000022, 0.0000021},
	{0.0000031, 0.0000032, 0.0000033, 0.0000034, 0.0000034, 0.0000033, 0.0000032, 0.0000031},
	{0.0000041, 0.0000042, 0.0000043, 0.0000044, 0.0000044, 0.0000043, 0.0000042, 0.0000041},
	{0.0000051, 0.0000052, 0.0000053, 0.0000053, 0.0000053, 0.0000053, 0.0000052, 0.0000051},
	{0.0000061, 0.0000062, 0.0000062, 0.0000062, 0.0000062, 0.0000062, 0.0000062, 0.0000061},
	{0.0000071, 0.0000071, 0.0000071, 0.0000071, 0.0000071, 0.0000071, 0.0000071, 0.0000071},
}
