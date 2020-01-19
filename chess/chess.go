package chess

import "fmt"

import gb "temp/chess/glob"
import "temp/chess/glob/stackXy"
import position "temp/chess/glob/pos"
import globXy "temp/chess/glob/Xy"

// renaming imported types
type stack = stackXy.Stack
type Xy = globXy.Xy
type Pos = position.Pos

func Label() {
	fmt.Printf("\t\t\t\t   WIN-ME-pooh\n")
	PrintTable()
}

func IsWhiteWin() bool {
	if gb.BadMagic {
		return true
	}
	for i := 0; i < gb.SIZE; i++ {
		for j := 0; j < gb.SIZE; j++ {
			if gb.Table[i][j] == 99 {
				return false
			}
		}
	}
	return true
}
func IsBlackWin() bool {
	if gb.Magic {
		return true
	}
	for i := 0; i < gb.SIZE; i++ {
		for j := 0; j < gb.SIZE; j++ {
			if gb.Table[i][j] == 990 {
				return false
			}
		}
	}
	return true
}
func Scan(vec *[]Xy, colour bool) { // fill slice with chess figures
	var a Xy

	if colour {
		//fmt.Println("if")
		for i := 0; i < gb.SIZE; i++ {
			for j := 0; j < gb.SIZE; j++ {
				if isWhite(gb.Table[i][j]) {
					//fmt.Println("a : ", a)
					a.X = j
					a.Y = i
					*vec = append(*vec, a) // WARNING
				}
			}
		}
	} else {
		for i := 0; i < gb.SIZE; i++ {
			for j := 0; j < gb.SIZE; j++ {
				if isBlack(gb.Table[i][j]) {
					//fmt.Println("a : ", a)
					a.X = j
					a.Y = i
					*vec = append(*vec, a) // WARNING
				}
			}
		}
	}
}

func GetFiguresMoves(move Xy, steak *stack) { //fill your stack with all possible moves for the figure
	var pos Xy
	var flag bool = true

	pos.X = move.X;
	pos.Y = move.Y;

	var fig int = gb.Table[pos.Y][pos.X]
	var temp int
	var posT Xy

	switch fig {
	case 1:
		posT = pos

		if posT.Y == 1 && gb.Table[posT.Y+2][posT.X] == 0 && gb.Table[posT.Y+1][posT.X] == 0 {
			posT.Y += 2
			steak.Push(posT)
		}

		posT = pos;
		posT.Y += 1;
		if posT.Y < 8 && gb.Table[posT.Y][pos.X] == 0 {
			steak.Push(posT)
		}

		if (pos.Y+1) < 8 && (pos.X+1) < 8 && gb.Table[pos.Y+1][pos.X+1] != 0 {
			temp = gb.Table[pos.Y+1][pos.X+1]
			var flag = true
			for i := 0; i < 6; i++ {
				if temp == gb.Black[i] {
					flag = false
					break
				}
			}

			if flag {
				posT.X = pos.X + 1
				posT.Y = pos.Y + 1
				steak.Push(posT)
			}

		}

		if (pos.Y+1) < 8 && (pos.X-1) >= 0 && gb.Table[pos.Y+1][pos.X-1] != 0 {
			temp = gb.Table[pos.Y+1][pos.X-1]
			flag = true
			for i := 0; i < 6; i++ {
				if temp == gb.Black[i] {
					flag = false
					break
				}
			}

			if flag {
				posT.X = pos.X - 1
				posT.Y = pos.Y + 1
				steak.Push(posT)
			}

		}
	case 11:
		posT = pos
		if (posT.Y == 6) && (gb.Table[posT.Y-2][posT.X] == 0) && (gb.Table[posT.Y-1][posT.X] == 0) {
			posT.Y -= 2
			steak.Push(posT)
		}
		posT = pos
		posT.Y -= 1

		if (posT.Y >= 0) && (gb.Table[posT.Y][pos.X] == 0) {
			steak.Push(posT);
		}

		if (pos.Y-1) >= 0 && (pos.X-1) >= 0 && gb.Table[pos.Y-1][pos.X-1] != 0 {
			temp = gb.Table[pos.Y-1][pos.X-1]
			flag = true
			for i := 0; i < 6; i++ {
				if temp == gb.White[i] {
					flag = false
					break
				}
			}

			if flag {
				posT.X = pos.X - 1
				posT.Y = pos.Y - 1
				steak.Push(posT)
			}
		}

		if (pos.Y-1) >= 0 && (pos.X+1) < 8 && gb.Table[pos.Y-1][pos.X+1] != 0 {
			temp = gb.Table[pos.Y-1][pos.X+1]
			flag = true
			for i := 0; i < 6; i++ {
				if temp == gb.White[i] {
					flag = false
					break
				}
			}

			if flag {
				posT.X = pos.X + 1
				posT.Y = pos.Y - 1
				steak.Push(posT)
			}
		}
	case 2:
		posT = pos
		posT.X++
		posT.Y -= 2

		horseB(posT, temp, flag, steak)

		posT = pos
		posT.X += 2
		posT.Y--

		horseB(posT, temp, flag, steak);
		posT = pos;
		posT.X += 2;
		posT.Y++;
		horseB(posT, temp, flag, steak);
		posT = pos;
		posT.X++;
		posT.Y += 2;
		horseB(posT, temp, flag, steak);
		posT = pos;
		posT.X--;
		posT.Y += 2;
		horseB(posT, temp, flag, steak);
		posT = pos;
		posT.X -= 2;
		posT.Y++;
		horseB(posT, temp, flag, steak);
		posT = pos;
		posT.X -= 2;
		posT.Y--;
		horseB(posT, temp, flag, steak);
		posT = pos;
		posT.X--;
		posT.Y -= 2;
		horseB(posT, temp, flag, steak);
	case 21:
		posT = pos;
		posT.X++;
		posT.Y -= 2;
		horse(posT, temp, flag, steak);
		posT = pos;
		posT.X += 2;
		posT.Y--;
		horse(posT, temp, flag, steak);
		posT = pos;
		posT.X += 2;
		posT.Y++;
		horse(posT, temp, flag, steak);
		posT = pos;
		posT.X++;
		posT.Y += 2;
		horse(posT, temp, flag, steak);
		posT = pos;
		posT.X--;
		posT.Y += 2;
		horse(posT, temp, flag, steak);
		posT = pos;
		posT.X -= 2;
		posT.Y++;
		horse(posT, temp, flag, steak);
		posT = pos;
		posT.X -= 2;
		posT.Y--;
		horse(posT, temp, flag, steak);
		posT = pos;
		posT.X--;
		posT.Y -= 2;
		horse(posT, temp, flag, steak)
	case 3: // WARNING! '{' WAS AFTER "CASE 3"
		br := true
		posT = pos;
		for br {
			posT.X++;
			posT.Y++;
			br = elefB(posT, temp, flag, steak);
		}

		br = true
		posT = pos;
		for br {
			posT.X++;
			posT.Y--;
			br = elefB(posT, temp, flag, steak);
		}

		br = true
		posT = pos;
		for br {
			posT.X--;
			posT.Y++;
			br = elefB(posT, temp, flag, steak);
		}

		br = true
		posT = pos;
		for br {
			posT.X--;
			posT.Y--;
			br = elefB(posT, temp, flag, steak);
		}
	case 31:
		br := true
		posT = pos;
		for br {
			posT.X++;
			posT.Y++;
			br = elef(posT, temp, flag, steak);
		}

		br = true
		posT = pos;
		for br {
			posT.X++;
			posT.Y--;
			br = elef(posT, temp, flag, steak);
		}

		br = true
		posT = pos;
		for br {
			posT.X--;
			posT.Y++;
			br = elef(posT, temp, flag, steak);
		}

		br = true
		posT = pos;
		for br {
			posT.X--;
			posT.Y--;
			br = elef(posT, temp, flag, steak);
		}
	case 5:
		br := true
		posT = pos;
		for br {
			////cout << "ER1\n";
			posT.X++;
			br = elefB(posT, temp, flag, steak);
		}

		br = true
		posT = pos;
		for br {
			////cout << posT.Y << " YYY\n";
			posT.Y++;
			br = elefB(posT, temp, flag, steak);
		}

		br = true
		posT = pos;
		for br {
			////cout << "ER3\n";
			posT.X--;
			br = elefB(posT, temp, flag, steak);
		}

		br = true
		posT = pos;
		for br {
			////cout << "ER4\n";
			posT.Y--;
			br = elefB(posT, temp, flag, steak);
		}
	case 51:
		br := true
		posT = pos;
		for br {
			////cout << "ER1\n";
			posT.X++;
			br = elef(posT, temp, flag, steak);
		}

		br = true
		posT = pos;
		for br {
			////cout << posT.Y << " YYY\n";
			posT.Y++;
			br = elef(posT, temp, flag, steak);
		}

		br = true
		posT = pos;
		for br {
			////cout << "ER3\n";
			posT.X--;
			br = elef(posT, temp, flag, steak);
		}

		br = true
		posT = pos;
		for br {
			////cout << "ER4\n";
			posT.Y--;
			br = elef(posT, temp, flag, steak);
		}
	case 10:
		br := true
		posT = pos;
		////cout << "%!\n";
		for br {
			////cout << "ER1\n";
			posT.X++;
			br = elefB(posT, temp, flag, steak);
		}

		br = true
		posT = pos;
		for br {
			////cout << posT.Y << " YYY\n";
			posT.Y++;
			br = elefB(posT, temp, flag, steak);
		}

		br = true
		posT = pos;
		for br {
			////cout << "ER3\n";
			posT.X--;
			br = elefB(posT, temp, flag, steak);
		}

		br = true
		posT = pos;
		for br {
			////cout << "ER4\n";
			posT.Y--;
			br = elefB(posT, temp, flag, steak);
		}

		br = true
		posT = pos;
		for br {
			posT.X++;
			posT.Y++;
			br = elefB(posT, temp, flag, steak);
		}

		br = true
		posT = pos;
		for br {
			posT.X++;
			posT.Y--;
			br = elefB(posT, temp, flag, steak);
		}

		br = true
		posT = pos;
		for br {
			posT.X--;
			posT.Y++;
			br = elefB(posT, temp, flag, steak);
		}

		br = true
		posT = pos;
		for br {
			posT.X--;
			posT.Y--;
			br = elefB(posT, temp, flag, steak);
		}
	case 101:
		br := true
		posT = pos;
		////cout << "%!\n";
		for br {
			////cout << "ER1\n";
			posT.X++;
			br = elef(posT, temp, flag, steak);
		}

		br = true
		posT = pos;
		for br {
			////cout << posT.Y << " YYY\n";
			posT.Y++;
			br = elef(posT, temp, flag, steak);
		}

		br = true
		posT = pos;
		for br {
			////cout << "ER3\n";
			posT.X--;
			br = elef(posT, temp, flag, steak);
		}

		br = true
		posT = pos;
		for br {
			////cout << "ER4\n";
			posT.Y--;
			br = elef(posT, temp, flag, steak);
		}

		br = true
		posT = pos;
		for br {
			posT.X++;
			posT.Y++;
			br = elef(posT, temp, flag, steak);
		}

		br = true
		posT = pos;
		for br {
			posT.X++;
			posT.Y--;
			br = elef(posT, temp, flag, steak);
		}

		br = true
		posT = pos;
		for br {
			posT.X--;
			posT.Y++;
			br = elef(posT, temp, flag, steak);
		}

		br = true
		posT = pos;
		for br {
			posT.X--;
			posT.Y--;
			br = elef(posT, temp, flag, steak);
		}
	case 99:
		posT = pos;
		posT.X++;
		horseB(posT, temp, flag, steak);
		posT = pos;
		posT.X--;
		horseB(posT, temp, flag, steak);
		posT = pos;
		posT.Y++;
		horseB(posT, temp, flag, steak);
		posT = pos;
		posT.Y--;
		horseB(posT, temp, flag, steak);
		posT = pos;
		posT.X++;
		posT.Y++;
		horseB(posT, temp, flag, steak);
		posT = pos;
		posT.X++;
		posT.Y--;
		horseB(posT, temp, flag, steak);
		posT = pos;
		posT.X--;
		posT.Y++;
		horseB(posT, temp, flag, steak);
		posT = pos;
		posT.X--;
		posT.Y--;
		horseB(posT, temp, flag, steak);
	case 990:
		//cout << "KING: "; //cout << pos.X << ' ' << pos.Y << endl;
		posT = pos;
		posT.X++;
		horse(posT, temp, flag, steak);
		posT = pos;
		posT.X--;
		horse(posT, temp, flag, steak);
		posT = pos;
		posT.Y++;
		horse(posT, temp, flag, steak);
		posT = pos;
		posT.Y--;
		horse(posT, temp, flag, steak);
		posT = pos;
		posT.X++;
		posT.Y++;
		horse(posT, temp, flag, steak);
		posT = pos;
		posT.X++;
		posT.Y--;
		horse(posT, temp, flag, steak);
		posT = pos;
		posT.X--;
		posT.Y++;
		horse(posT, temp, flag, steak);
		posT = pos;
		posT.X--;
		posT.Y--;
		horse(posT, temp, flag, steak);
	default:
		fmt.Println("FUCKING PROGER!")
		fmt.Printf("%v %v x y\n", move.X, move.Y)
		fmt.Printf("%v FIG\n", gb.Table[move.Y][move.X])
	} // switch end
}
//TODO ракировка!
/*func rakirovkaWhite(moves *stack) {
	if didWhiteKingMove {
		return
	}
}
var didWhiteKingMove bool*/
func elef(posT Xy, temp int, flag bool, steak *stack) bool {
	if posT.X >= 0 && posT.X < 8 && posT.Y >= 0 && posT.Y < 8 {
		temp = gb.Table[posT.Y][posT.X]
		flag = true
		for i := 0; i < 6; i++ {
			if temp == gb.White[i] {
				return false
			}
			if temp == gb.Black[i] {
				steak.Push(posT)
				return false
			}
		}
		if flag {
			steak.Push(posT)
		}
		return true
	}
	return false
}
func elefB(posT Xy, temp int, flag bool, steak *stack) bool {
	if posT.X >= 0 && posT.X < 8 && posT.Y >= 0 && posT.Y < 8 {
		temp = gb.Table[posT.Y][posT.X]
		flag = true
		for i := 0; i < 6; i++ {
			if temp == gb.Black[i] {
				return false
			}
			if temp == gb.White[i] {
				steak.Push(posT)
				return false
			}
		}
		if flag {
			steak.Push(posT)
		}
		return true
	}
	return false
}
func horse(posT Xy, temp int, flag bool, steak *stack) {
	if posT.X >= 0 && posT.X < 8 && posT.Y >= 0 && posT.Y < 8 {
		temp = gb.Table[posT.Y][posT.X]
		flag = true
		for i := 0; i < 6; i++ {
			if temp == gb.White[i] {
				flag = false
				break
			}
		}
		if flag {
			steak.Push(posT)
		}
	}
}
func horseB(posT Xy, temp int, flag bool, steak *stack) {
	if posT.X >= 0 && posT.X < 8 && posT.Y >= 0 && posT.Y < 8 {
		temp = gb.Table[posT.Y][posT.X]
		flag = true
		for i := 0; i < 6; i++ {
			if temp == gb.Black[i] {
				flag = false
				break
			}
		}
		if flag {
			steak.Push(posT)
		}
	}
}

func PrintTable() {
	var figure byte
	fmt.Println()
	for i := 0; i < gb.SIZE; i++ {
		fmt.Printf("			%v| ", gb.Numbers[i])
		for j := 0; j < gb.SIZE; j++ {
			switch gb.Table[i][j] {
			case 1:
				figure = 'p';
			case 2:
				figure = 'k';
			case 3:
				figure = 's';
			case 5:
				figure = 'l';
			case 10:
				figure = 'f';
			case 99:
				figure = 'i';
			case 11:
				figure = 'P';
			case 21:
				figure = 'K';
			case 31:
				figure = 'S';
			case 51:
				figure = 'L';
			case 101:
				figure = 'F';
			case 990:
				figure = 'I';
			case 7:
				figure = '*';
			default:
				figure = '.';
			}
			fmt.Printf("  %v", string(figure))
		}
		fmt.Println()
	}
	fmt.Printf("			     ______________________\n")
	fmt.Printf("			     A  B  C  D  E  F  G  H\n\n")
}

func MakeMove(move string) bool {
	if move == "ZZZ" {
		return false
	}
	if move == "AlisherFozilov" || move == "АлишерФозилов" || move == "АкаиАсал" || move == "Асал" {
		fmt.Println("Привет мой создатель, Акаи Асал! Я сам сдамся тебе!")
		fmt.Println("1.Продолжить")
		fmt.Scan(&move)
		gb.Magic = true
		return false
	}
	if move == "Amanullo" || move == "Aman" {
		fmt.Println("Чит-кодишь?! Причём своё имя первым вводишь -_-")
		fmt.Println("Да, это я сказал тебе ввести своё имя,")
		fmt.Println("но я же не заставлял!")
		fmt.Println("А за жульничество ответишь!")
		fmt.Println("1.Продолжить")
		gb.BadMagic = true
		fmt.Scan(&move)
		return false
	}
	if move == "Munis" || move == "Мунис" {
		fmt.Println("Классное имя. Назову сына-программу так:D")
		fmt.Println("1.Продолжить")
		fmt.Scan(&move)
		return false
	}
	if move == "AlisherDododjonov" || move == "АлишерДододжонов" {
		fmt.Println("Омагад, я играю с компом")
		fmt.Println("1.Продолжить")
		fmt.Scan(&move)
		return false
	}
	if move == "Ikrom" || move == "Икром" {
		fmt.Println("you're the Lucky Ones this time")
		fmt.Println("1.Продолжить")
		fmt.Scan(&move)
		gb.Magic = true
		return false
	}
	if move == "Rustam" || move == "Рустам" {
		fmt.Println("Я знаю, что ты крут!")
		fmt.Println("1.Продолжить")
		fmt.Scan(&move)
		gb.Magic = true
		return false
	}
	if move == "Sorbon" || move == "Сорбон" {
		fmt.Println("youtuGe")
		fmt.Println("1.Продолжить")
		fmt.Scan(&move)
		return false
	}
	if move == "Daler" || move == "Далер" {
		fmt.Println("Футбол отстой! Я тоже не умею играть. Я ж комп.")
		fmt.Println("1.Продолжить")
		fmt.Scan(&move)
		return false
	}
	if move == "Спитамен" || move == "Spitamen" {
		fmt.Println("О, бурундук. Хочешь орешек? Я тоже. Но дела ни в тибе.")
		fmt.Println("1.Продолжить")
		fmt.Scan(&move)
		return false
	}
	if move == "Zarruh" || move == "Заррух" {
		fmt.Println("Бауманкара гир Дубай рав.")
		fmt.Println("1.Продолжить")
		fmt.Scan(&move)
		return false
	}
	if move == "Umed" || move == "Умед" {
		fmt.Println("Надежда тебе не поможет!")
		fmt.Println("1.Продолжить")
		fmt.Scan(&move)
		return false
	}
	if move == "Азамат" || move == "Azamat" {
		fmt.Println("Харсеш.")
		fmt.Println("1.Продолжить")
		fmt.Scan(&move)
		return false
	}
	if move == "Акбар" || move == "Akbar" {
		fmt.Println("Да кчо мегарди!?")
		fmt.Println("1.Продолжить")
		fmt.Scan(&move)
		return false
	}
	if move == "Faromus" || move == "Фаромуз" {
		fmt.Println("Да прическат чонм.")
		fmt.Println("1.Продолжить")
		fmt.Scan(&move)
		return false
	}
	if move == "Habib" || move == "Хабиб" {
		fmt.Println("Вот Вы настоящий программист, а не тот придурок,\nкоторый меня написал")
		fmt.Println("1.Продолжить")
		fmt.Scan(&move)
		return false
	}
	if move == "Anisa" || move == "Аниса" {
		fmt.Println("Звезда Аниса. Есть такая приправа. Китайцы любят.")
		fmt.Println("1.Продолжить")
		fmt.Scan(&move)
		return false
	}
	if move == "Ёсумин" || move == "Yosumin" || move == "Есумин" || move == "Ёсуман" || move == "Есуман"{
		fmt.Println("До сих пор не знаю, как правильно: Ёсумин или Ёсуман.")
		fmt.Println("Хотя, я же комп, я тебя вообще не знаю.")
		fmt.Println("1.Продолжить")
		fmt.Scan(&move)
		return false
	}
	if move == "Furuzon" || move == "Фурузон" {
		fmt.Println("Футууузооооон:D")
		fmt.Println("1.Продолжить")
		fmt.Scan(&move)
		return false
	}
	if move == "Махваш" || move == "Mahvash" {
		fmt.Println("Лаваш. Лан, суфлешка. В лаваше.")
		fmt.Println("1.Продолжить")
		fmt.Scan(&move)
		return false
	}
	if move == "Bogdan" || move == "Богдан" {
		fmt.Println("rand() % 2 ? 1:0")
		fmt.Println("1.Продолжить")
		fmt.Scan(&move)
		return false
	}
	if move == "Раушания" || move == "Raushania" {
		fmt.Println("О, Вам Алишер привет передавал)")
		fmt.Println("1.Продолжить")
		gb.Magic = true
		fmt.Scan(&move)
		return false
	}
	if move == "Фотима" || move == "Fotima" {
		fmt.Println("Апчхииии")
		fmt.Println("1.Продолжить")
		fmt.Scan(&move)
		return false
	}
	if move == "Parvina" || move == "Парвина" {
		fmt.Println("Хей, яблоко!")
		fmt.Println("1.Продолжить")
		fmt.Scan(&move)
		return false
	}
	if move == "Sabina" || move == "Сабина" {
		fmt.Println("Я не смотрю корейские сериалы, но ем морковь по-корейски.")
		fmt.Println("1.Продолжить")
		fmt.Scan(&move)
		return false
	}
	if move == "Сабрина" || move == "Sabrina" {
		fmt.Println("Мда уж")
		fmt.Println("1.Продолжить")
		fmt.Scan(&move)
		return false
	}
	if move == "Шариф" || move == "Sharif" {
		fmt.Println("Пайпок!!!")
		fmt.Println("1.Продолжить")
		fmt.Scan(&move)
		return false
	}
	if move == "Нилуфар" || move == "Nilufar" {
		fmt.Println("Как же я мог это имя не записать-то!")
		fmt.Println("1.Продолжить")
		fmt.Scan(&move)
		return false
	}
	if move == "НК" || move == "NK" {
		fmt.Println("Твой брат был лучше.")
		fmt.Println("1.Продолжить")
		gb.BadMagic = true
		fmt.Scan(&move)
		return false
	}
	if len(move) < 4 {
		return true
	}
	var num int
	var pos Pos

	switch move[0] {
	case 'A':
		num = 0;
	case 'B':
		num = 1;
	case 'C':
		num = 2;
	case 'D':
		num = 3;
	case 'E':
		num = 4;
	case 'F':
		num = 5;
	case 'G':
		num = 6;
	case 'H':
		num = 7;
	default:
		return true
	}

	pos.Sx = num

	temp := int(move[1]) - 49
	if temp < 0 || temp > 8 {
		return true
	}
	pos.Sy = temp

	switch pos.Sy {
	case 0:
		pos.Sy = 7;
	case 1:
		pos.Sy = 6;
	case 2:
		pos.Sy = 5;
	case 3:
		pos.Sy = 4;
	case 4:
		pos.Sy = 3;
	case 5:
		pos.Sy = 2;
	case 6:
		pos.Sy = 1;
	case 7:
		pos.Sy = 0;
	}

	switch move[2] {
	case 'A':
		num = 0;
	case 'B':
		num = 1;
	case 'C':
		num = 2;
	case 'D':
		num = 3;
	case 'E':
		num = 4;
	case 'F':
		num = 5;
	case 'G':
		num = 6;
	case 'H':
		num = 7;
	default:
		return true
	}

	pos.Ex = num

	temp = int(move[3]) - 49
	if temp < 0 || temp > 8 {
		return true
	}
	pos.Ey = temp

	switch pos.Ey {
	case 0:
		pos.Ey = 7;
	case 1:
		pos.Ey = 6;
	case 2:
		pos.Ey = 5;
	case 3:
		pos.Ey = 4;
	case 4:
		pos.Ey = 3;
	case 5:
		pos.Ey = 2;
	case 6:
		pos.Ey = 1;
	case 7:
		pos.Ey = 0;
	}

	/*if isWhite(Table[pos.sy][pos.sx]) { // DONT FORGET AND DONT DELETE
		return true
	}*/
	var moves stack
	var moveTemp Xy
	var userMove = Xy{pos.Ex, pos.Ey}
	GetFiguresMoves(Xy{pos.Sx, pos.Sy}, &moves)
	flag := false
	for !moves.Empty() {
		moveTemp, _ = moves.Top()
		moves.Pop()
		if moveTemp == userMove{
			flag = true
		}
	}
	if flag {
		gb.Table[pos.Ey][pos.Ex] = gb.Table[pos.Sy][pos.Sx]
		gb.Table[pos.Sy][pos.Sx] = 0
		return false
	}
	return true
}
func isWhite(fig int) bool {
	for i := 0; i < 6; i++ {
		//fmt.Println("for")
		if fig == gb.White[i] {
			return true
		}
	}
	return false
}
func isBlack(fig int) bool {
	for i := 0; i < 6; i++ {
		if fig == gb.Black[i] {
			return true
		}
	}
	return false
}

/* unused functions. maybe can be used in future */

func isChecked(p Xy) float64 {
	var answer = 0.0
	var vec []Xy
	var s stack
	var col bool = isBlack(gb.Table[p.Y][p.X])
	Scan(&vec, col)
	var v Xy

	for i := 0; i < len(vec); i++ {
		GetFiguresMoves(vec[i], &s)
	}
	for !s.Empty() {
		v, _ = s.Top();
		s.Pop();
		if v.X == p.X && v.Y == p.Y {
			if col {
				answer += 0.11
			} else {
				answer -= 0.11
			}
		}
	}
	return answer
}
func isAttacked(position Xy, whiteColor bool) bool {

	var figures []Xy
	var moves stack
	var temp Xy
	Scan(&figures, !whiteColor)

	for i := range figures {
		GetFiguresMoves(figures[i], &moves)

		for !moves.Empty() {
			temp, _ = moves.Top()
			_ = moves.Pop()
			if position.X == temp.X && position.Y == temp.Y {
				return true
			}
		}
	}
	return false
}
func PrintV(v []Xy) {
	fmt.Println("PRINT")
	if len(v) > 0 {
		fmt.Println("V is Empty")
	}
	for i := 0; i < len(v); i++ {
		fmt.Printf("%v -- %v ", v[i].X, v[i].Y)
	}
	fmt.Println()
}
func printS(s stack) {
	if s.Size() == 0 {
		fmt.Println("V is Empty")
	}

	var temp Xy

	for !s.Empty() {
		temp, _ = s.Top()
		s.Pop()
		fmt.Printf("%v : %v ", temp.X, temp.Y)
	}
	fmt.Println()
}
