package game

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/alisherfozilov/chess-game/chess"
	"github.com/alisherfozilov/chess-game/chess/AI"
	gb "github.com/alisherfozilov/chess-game/chess/glob"
	"github.com/alisherfozilov/chess-game/cmd"
)

func init() {
	settings.sound = true
	settings.thinkLevel = 3
	settings.timer = -1
	settings.boardType = BoardTypeStandard
}

func Play() {
	var whiteWinner bool
	gb.Magic = false
	gb.BadMagic = false
	if gb.Table != gb.TableStandartPosition {
		chess.Label()
		fmt.Println("Вы хотите продолжить предыдущую игру?")
		fmt.Printf("1.Да\t\t\t2.Нет\n")

		var move int
		_, err := fmt.Scan(&move)

		if err != nil {
			fmt.Println("ERROR", err)
			return
		}
		if move == 2 {
			switch settings.boardType {
			case BoardTypeStandard:
				gb.Table = gb.TableStandartPosition
			case BoardTypeRandom:
				gb.Table = randomTablePosition()
			}
			gb.PlayerTime = settings.timer
		}
	} else {
		gb.PlayerTime = settings.timer
	}

	var move string
	cmd.Clear()
	fmt.Println("Для выхода в меню введите \"EXIT\" ")
	for {
		chess.Label()
		startMoveTime := time.Now()
		_, err := fmt.Scan(&move)
		if err != nil {
			fmt.Println("ERROR", err)
			return
		}
		moveTime := time.Since(startMoveTime)
		if gb.PlayerTime != -1 {
			gb.PlayerTime -= moveTime
		}

		if strings.ToUpper(move) == "EXIT" {
			return
		}

		invalid, keep := chess.MakeMove(move)
		for invalid || keep {
			cmd.Clear()
			fmt.Println()
			chess.Label()
			if !keep {
				fmt.Println("Некорректный ход. Повторите попытку:")
			}
			_, err := fmt.Scan(&move)
			if err != nil {
				fmt.Println("ERROR", err)
				return
			}
			if move == "EXIT" {
				return
			}
			invalid, keep = chess.MakeMove(move)
		}

		cmd.Clear()
		if chess.IsWhiteWin() {
			whiteWinner = true
			break
		} else if chess.IsBlackWin() {
			whiteWinner = false
			break
		}

		fmt.Println()
		chess.Label()
		computerMove := AI.Think(settings.thinkLevel) // here's a place for 'thinking' algorithm!
		if settings.sound {
			fmt.Printf("\a")
		}
		cmd.Clear()
		gb.Table[computerMove.ToY][computerMove.ToX] = gb.Table[computerMove.FromY][computerMove.FromX]
		gb.Table[computerMove.FromY][computerMove.FromX] = 0

		fmt.Printf("Computer's move | from  %q%v  to  %q%v\n", gb.Letters[computerMove.FromX], gb.Numbers[computerMove.FromY], gb.Letters[computerMove.ToX], gb.Numbers[computerMove.ToY])

		if chess.IsWhiteWin() {
			whiteWinner = true
			break
		} else if chess.IsBlackWin() {
			whiteWinner = false
			break
		}
	}
	cmd.Clear()
	chess.Label()
	chess.PrintTable()
	if whiteWinner {
		fmt.Println("Компьютер тебя уделал, но ты не сдавайся!")
		fmt.Println("(с) Конфуций")
	} else {
		fmt.Println("Ого, ты победил! Мои поздравления!")
		fmt.Println("Спасибо за то, что поиграл.")
		fmt.Println("Надеюсь было весело:D")
	}
	fmt.Println("\n\n1.Вернуться в меню")
	var temp string
	_, _ = fmt.Scan(&temp)
}

func randomTablePosition() [8][8]int {
	board := gb.TableStandartPosition
	rand.Shuffle(len(board[0]), func(i, j int) {
		board[0][i], board[0][j] = board[0][j], board[0][i]
		board[7][i], board[7][j] = board[7][j], board[7][i]
	})
	return board
}

func init() {
	rand.Seed(time.Now().Unix())
}

type Settings struct {
	choice     int
	thinkLevel int
	sound      bool
	timer      time.Duration
	boardType  BoardType
}

type BoardType int

const (
	BoardTypeStandard = iota
	BoardTypeRandom
)

var settings Settings

func (s *Settings) print() {
	fmt.Printf("\t\tНастройки\n\n")
	fmt.Println("1.Сложность")
	fmt.Println("2.Звук")
	fmt.Println("3.Танцевать при взятии фигуры")
	fmt.Println("4.Таймер")
	fmt.Println("5.Расположение фигур")
	fmt.Println("6.Выйти в меню")
}
func (s *Settings) switchChoice() bool {
	cmd.Clear()
	fmt.Println()
	chess.Label()
	fmt.Printf("\t\tНастройки\n\n")
	switch s.choice {
	case 1:
		s.setThinkLevel()
	case 2:
		s.setSound()
	case 3:
		s.dance()
	case 4:
		s.setTimer()
	case 5:
		s.setBoardType()
	case 6:
		return false
	}
	return true
}
func (s *Settings) inputChoice() {
	_, err := fmt.Scan(&s.choice)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
}
func (s *Settings) setThinkLevel() {
	fmt.Println("Самый лучший вариант - 4. Он установлен по умолчанию,")
	fmt.Println("а уровень 5 немного медлит(ага, немного)")
	fmt.Println("3 - просто. 2 - слишком просто.")
	fmt.Println("Если вы попугай, смело ставьте 1\n")
metka:
	fmt.Printf("Введите уровень сложности от 1 до 6 включительно:\n")
	s.inputChoice()
	if s.choice > 0 && s.choice < 7 {
		if s.choice == 6 {
			fmt.Println("На этом режиме сложности можете спокойно идти пить кофе,")
			fmt.Println("в ожидании хода компьютера. Если вы запустили ALICHESS")
			fmt.Println("с калькулятора, можете не возвращаться")
		} else {
			s.thinkLevel = s.choice - 1
			return
		}
	} else if s.choice < 1 {
		fmt.Println("Я понимаю Ваше желание выиграть у компьютера,")
		fmt.Println("но не до такой же степени!")
	} else {
		fmt.Println("Вижу Вам нравится хардкор. Мне тоже. Но у Вашего ноута не хватит")
		fmt.Println("вычислительной мощности. Да и вряд ли у кого-то хватит")
		fmt.Println("Но это пока что...")
	}
	goto metka
}
func (s *Settings) setSound() {
metka:
	fmt.Println("1.Вкл")
	fmt.Println("2.Выкл")
	s.inputChoice()
	if s.choice == 1 {
		s.sound = true
		return
	} else if s.choice == 2 {
		s.sound = false
		return
	} else {
		fmt.Println("Вы сделали неверный выбор. Конец человечеству.")
	}
	goto metka
}
func (s *Settings) dance() {
	fmt.Println("1.Выкл")
	fmt.Println("P.S. компьютер непротив того, чтобы Вы сами потанцевали:D")
	s.inputChoice()
}

func (s *Settings) setTimer() {
	fmt.Println("0. Выкл.")
	fmt.Println("1. 5 Минут")
	fmt.Println("2. 10 Минут")
	fmt.Println("3. 20 Минут")
	fmt.Println("4. 30 Минут")
	fmt.Println("5. 1 час")
	fmt.Println("6. 3 часа")
	s.inputChoice()
	switch s.choice {
	case 0:
		settings.timer = -1
	case 1:
		settings.timer = 5 * time.Minute
	case 2:
		settings.timer = 10 * time.Minute
	case 3:
		settings.timer = 20 * time.Minute
	case 4:
		settings.timer = 30 * time.Minute
	case 5:
		settings.timer = time.Hour
	case 6:
		settings.timer = 3 * time.Hour
	}
}

func (s *Settings) setBoardType() {
	fmt.Println("1. Классическая")
	fmt.Println("2. Куча мала")
	s.inputChoice()
	switch s.choice {
	case 1:
		s.boardType = BoardTypeStandard
	case 2:
		s.boardType = BoardTypeRandom
	}
}
func SettingsFunc() {
	flag := true
	for flag {
		cmd.Clear()
		fmt.Println()
		chess.Label()
		settings.print()
		settings.inputChoice()
		flag = settings.switchChoice()
	}
}
