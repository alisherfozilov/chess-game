package showCode

import (
	"fmt"
	"temp/cmd"
)

func ShowCode() {
	cmd.Clear()
	fmt.Println(
		"package main\n" +
			"\n" +
			"import (\n" +
			"	fmt\n" +
			"	os\n" +
			"	os/exec\n" +
			")\n" +
			"\n" +
			"type NegativeIndexError struct {\n" +
			"	functionName string\n" +
			"	index        int\n" +
			"}\n" +
			"func (s NegativeIndexError) Error() string {\n" +
			"	return fmt.Sprintf(v: negative index [v], s.functionName, s.index)\n" +
			"}\n" +
			"\n" +
			"type stack struct {\n" +
			"	Slice []Xy\n" +
			"}\n" +
			"func (s *stack) push(elem Xy) {\n" +
			"	s.Slice = append(s.Slice, elem)\n" +
			"}\n" +
			"func (s *stack) top() (Xy, error) {\n" +
			"	if len(s.Slice) < 1 {\n" +
			"		return Xy{0, 0}, NegativeIndexError{\n" +
			"			functionName: stack.top(),\n" +
			"			index:        len(s.Slice) - 1,\n" +
			"		}\n" +
			"	}\n" +
			"	return s.Slice[len(s.Slice)-1], nil\n" +
			"}\n" +
			"func (s *stack) pop() error {\n" +
			"	if len(s.Slice) < 1 {\n" +
			"		return NegativeIndexError{\n" +
			"			functionName: stack.pop(),\n" +
			"			index:        len(s.Slice) - 1,\n" +
			"		}\n" +
			"	}\n" +
			"	s.Slice = s.Slice[:len(s.Slice)-1]\n" +
			"	return nil\n" +
			"}\n" +
			"func (s *stack) empty() bool {\n" +
			"	if len(s.Slice) == 0 {\n" +
			"		return true\n" +
			"	}\n" +
			"	return false\n" +
			"}\n" +
			"func (s *stack) size() int {\n" +
			"	return len(s.Slice)\n" +
			"}\n" +
			"\n" +
			"type Pos struct {\n" +
			"	sx int\n" +
			"	sy int\n" +
			"	ex int\n" +
			"	ey int\n" +
			"}\n" +
			"type Xy struct {\n" +
			"	x int\n" +
			"	y int\n" +
			"}\n" +
			"\n" +
			"var letters = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H' }\n" +
			"var numbers = []int{8, 7, 6, 5, 4, 3, 2, 1}\n" +
			"var black = []int{1, 2, 3, 5, 10, 99}\n" +
			"var white = []int{11, 21, 31, 51, 101, 990}\n" +
			"\n" +
			"const SIZE = 8\n" +
			"\n" +
			"var table = [SIZE][SIZE]int {\n" +
			"	{5, 2, 3, 10, 99, 3, 2, 5},\n" +
			"	{1, 1, 1, 1, 1, 1, 1, 1},\n" +
			"	{0, 0, 0, 0, 0, 0, 0, 0},\n" +
			"	{0, 0, 0, 0, 0, 0, 0, 0},\n" +
			"	{0, 0, 0, 0, 0, 0, 0, 0},\n" +
			"	{0, 0, 0, 0, 0, 0, 0, 0},\n" +
			"	{11, 11, 11, 11, 11, 11, 11, 11},\n" +
			"	{51, 21, 31, 101, 990, 31, 21, 51},\n" +
			"}\n" +
			"var tableStandartPosition = [SIZE][SIZE]int{\n" +
			"	{5, 2, 3, 10, 99, 3, 2, 5},\n" +
			"	{1, 1, 1, 1, 1, 1, 1, 1},\n" +
			"	{0, 0, 0, 0, 0, 0, 0, 0},\n" +
			"	{0, 0, 0, 0, 0, 0, 0, 0},\n" +
			"	{0, 0, 0, 0, 0, 0, 0, 0},\n" +
			"	{0, 0, 0, 0, 0, 0, 0, 0},\n" +
			"	{11, 11, 11, 11, 11, 11, 11, 11},\n" +
			"	{51, 21, 31, 101, 990, 31, 21, 51},\n" +
			"}\n" +
			"\n" +
			"var counter = 0\n" +
			"\n" +
			"type ComputerMove struct {\n" +
			"	fromX int\n" +
			"	fromY int\n" +
			"	toX   int\n" +
			"	toY   int\n" +
			"}\n" +
			"\n" +
			"func Label() {\n" +
			"	fmt.Printf(ttttt   ALICHESSn)\n" +
			"	printTable()\n" +
			"}\n" +
			"\n" +
			"type Menu struct {\n" +
			"	choice int\n" +
			"}\n" +
			"func (m *Menu) print () {\n" +
			"	fmt.Printf(ttМенюnn)\n" +
			"	fmt.Println(1.Играть)\n" +
			"	fmt.Println(2.Настройки)\n" +
			"	fmt.Println(3.Правила)\n" +
			"	fmt.Println(4.О программе)\n" +
			"	fmt.Println(5.Выход)\n" +
			"	fmt.Println(6.Шутка про шахматы)\n" +
			"}\n" +
			"func (m *Menu) inputChoice () {\n" +
			"	_, err := fmt.Scan(&m.choice)\n" +
			"	if err != nil {\n" +
			"		fmt.Println(ERROR, err)\n" +
			"		return\n" +
			"	}\n" +
			"}\n" +
			"func (m *Menu) switchChoice() bool {\n" +
			"	clearWindowsConsole()\n" +
			"	switch m.choice {\n" +
			"	case 1:\n" +
			"		play()\n" +
			"	case 2:\n" +
			"		settingsFunc()\n" +
			"	case 3:\n" +
			"		rules()\n" +
			"	case 4:\n" +
			"		about()\n" +
			"	case 5:\n" +
			"		return exitChess(0)\n" +
			"	case 6:\n" +
			"		joke()\n" +
			"	}\n" +
			"	return true\n" +
			"}\n" +
			"\n" +
			"type Settings struct {\n" +
			"	choice int\n" +
			"	thinkLevel int\n" +
			"	sound bool\n" +
			"}\n" +
			"var settings Settings\n" +
			"func (s *Settings) print() {\n" +
			"	fmt.Printf(ttНастройкиnn)\n" +
			"	fmt.Println(1.Сложность)\n" +
			"	fmt.Println(2.Звук)\n" +
			"	fmt.Println(3.Танцевать при взятии фигуры)\n" +
			"	fmt.Println(4.Выйти в меню)\n" +
			"}\n" +
			"func (s *Settings) switchChoice() bool {\n" +
			"	clearWindowsConsole()\n" +
			"	fmt.Println()\n" +
			"	Label()\n" +
			"	fmt.Printf(ttНастройкиnn)\n" +
			"	switch s.choice {\n" +
			"	case 1:\n" +
			"		s.setThinkLevel()\n" +
			"	case 2:\n" +
			"		s.setSound()\n" +
			"	case 3:\n" +
			"		s.dance()\n" +
			"	case 4:\n" +
			"		return false\n" +
			"	}\n" +
			"	return true\n" +
			"}\n" +
			"func (s *Settings) inputChoice () {\n" +
			"	_, err := fmt.Scan(&s.choice)\n" +
			"	if err != nil {\n" +
			"		fmt.Println(ERROR, err)\n" +
			"		return\n" +
			"	}\n" +
			"}\n" +
			"func (s *Settings) setThinkLevel() {\n" +
			"	fmt.Println(Самый лучший вариант - 4. Он установлен по умолчанию,)\n" +
			"	fmt.Println(а уровень 5 немного медлит(ага, немного))\n" +
			"	fmt.Println(3 - просто. 2 - слишком просто.)\n" +
			"	fmt.Println(Если вы попугай, смело ставьте 1n)\n" +
			"	metka:\n" +
			"	fmt.Printf(Введите уровень сложности от 1 до 6 включительно:n)\n" +
			"	s.inputChoice()\n" +
			"	if s.choice > 0 && s.choice < 7 {\n" +
			"		if s.choice == 6 {\n" +
			"			fmt.Println(На этом режиме сложности можете спокойно идти пить кофе,)\n" +
			"			fmt.Println(в ожидании хода компьютера. Если вы запустили ALICHESS)\n" +
			"			fmt.Println(с калькулятора, можете не возвращаться)\n" +
			"		} else {\n" +
			"			s.thinkLevel = s.choice - 1\n" +
			"			return\n" +
			"		}\n" +
			"	} else if s.choice < 1 {\n" +
			"		fmt.Println(Я понимаю Ваше желание выиграть у компьютера,)\n" +
			"		fmt.Println(но не до такой же степени!)\n" +
			"	} else {\n" +
			"		fmt.Println(Вижу Вам нравится хардкор. Мне тоже. Но у Вашего ноута не хватит)\n" +
			"		fmt.Println(вычислительной мощности. Да и вряд ли у кого-то хватит)\n" +
			"		fmt.Println(Но это пока что...)\n" +
			"	}\n" +
			"	goto metka\n" +
			"}\n" +
			"func (s *Settings) setSound() {\n" +
			"	metka:\n" +
			"	fmt.Println(1.Вкл)\n" +
			"	fmt.Println(2.Выкл)\n" +
			"	s.inputChoice()\n" +
			"	if s.choice == 1 {\n" +
			"		s.sound = true\n" +
			"		return\n" +
			"	} else if s.choice == 2 {\n" +
			"		s.sound = false\n" +
			"		return\n" +
			"	} else {\n" +
			"		fmt.Println(Вы сделали неверный выбор. Конец человечеству.)\n" +
			"	}\n" +
			"	goto metka\n" +
			"}\n" +
			"func (s *Settings) dance() {\n" +
			"	fmt.Println(1.Выкл)\n" +
			"	fmt.Println(P.S. компьютер непротив того, чтобы Вы сами потанцевали:D)\n" +
			"	s.inputChoice()\n" +
			"}\n" +
			"\n" +
			"func clearWindowsConsole() {\n" +
			"	cmd := exec.Command(cmd, /c, cls) //Windows example, its tested\n" +
			"	cmd.Stdout = os.Stdout\n" +
			"	_ = cmd.Run()\n" +
			"}\n" +
			"// настройки по умолчанию\n" +
			"func initSettings() {\n" +
			"	settings.sound = true\n" +
			"	settings.thinkLevel = 3\n" +
			"}\n" +
			"\n" +
			"func main() {\n" +
			"\n" +
			"	clearWindowsConsole()\n" +
			"	initSettings()\n" +
			"	flag := true\n" +
			"	var menu Menu\n" +
			"\n" +
			"	for flag {\n" +
			"		clearWindowsConsole()\n" +
			"		fmt.Println()\n" +
			"		Label()\n" +
			"		menu.print()\n" +
			"		menu.inputChoice()\n" +
			"		flag = menu.switchChoice()\n" +
			"	}\n" +
			"}\n" +
			"\n" +
			"func play()  {\n" +
			"\n" +
			"	if table != tableStandartPosition {\n" +
			"		Label()\n" +
			"		fmt.Println(Вы хотите продолжить предыдущую игру?)\n" +
			"		fmt.Printf(1.Даttt2.Нетn)\n" +
			"\n" +
			"		var move int\n" +
			"		_, err := fmt.Scan(&move)\n" +
			"\n" +
			"		if err != nil {\n" +
			"			fmt.Println(ERROR, err)\n" +
			"			return\n" +
			"		}\n" +
			"		if move == 2 {\n" +
			"			table = tableStandartPosition\n" +
			"		}\n" +
			"	}\n" +
			"\n" +
			"	var move string\n" +
			"	clearWindowsConsole()\n" +
			"	fmt.Println(Для выхода в меню введите EXIT )\n" +
			"	for {\n" +
			"		Label()\n" +
			"		_, err := fmt.Scan(&move)\n" +
			"		if err != nil {\n" +
			"			fmt.Println(ERROR, err)\n" +
			"			return\n" +
			"		}\n" +
			"		if move == EXIT {\n" +
			"			return\n" +
			"		}\n" +
			"\n" +
			"		for makeMove(move) {\n" +
			"			clearWindowsConsole()\n" +
			"			fmt.Println()\n" +
			"			Label()\n" +
			"			fmt.Println(Некорректный ход. Повторите попытку:)\n" +
			"			_, err := fmt.Scan(&move)\n" +
			"			if err != nil {\n" +
			"				fmt.Println(ERROR, err)\n" +
			"				return\n" +
			"			}\n" +
			"			if move == EXIT {\n" +
			"				return\n" +
			"			}\n" +
			"		}\n" +
			"\n" +
			"		clearWindowsConsole()\n" +
			"		fmt.Println()\n" +
			"		Label()\n" +
			"\n" +
			"		computerMove := think(settings.thinkLevel) // here's a place for 'thinking' algorithm!\n" +
			"		if settings.sound {\n" +
			"			fmt.Printf(a)\n" +
			"		}\n" +
			"		clearWindowsConsole()\n" +
			"		table[computerMove.toY][computerMove.toX] = table[computerMove.fromY][computerMove.fromX]\n" +
			"		table[computerMove.fromY][computerMove.fromX] = 0\n" +
			"\n" +
			"		fmt.Printf(Computer's move | from  qv  to  qvn, letters[computerMove.fromX], numbers[computerMove.fromY], letters[computerMove.toX], numbers[computerMove.toY])\n" +
			"	}\n" +
			"}\n" +
			"func settingsFunc() {\n" +
			"	flag := true\n" +
			"	for flag {\n" +
			"		clearWindowsConsole()\n" +
			"		fmt.Println()\n" +
			"		Label()\n" +
			"		settings.print()\n" +
			"		settings.inputChoice()\n" +
			"		flag = settings.switchChoice()\n" +
			"	}\n" +
			"}\n" +
			"func rules() {\n" +
			"	clearWindowsConsole()\n" +
			"	fmt.Println()\n" +
			"	Label()\n" +
			"	fmt.Printf(ttПравилаnn)\n" +
			"	fmt.Println(Как играть в шахматы? Загуг... Заяндекси:)n(Скрытая реклама яндекса. Ой! Уже не скрытая))\n" +
			"	fmt.Printf(nttУправлениеnn)\n" +
			"	fmt.Println(Вы играете за черных. Ваши фигуры - те, которые сверху. Строчные буквы.)\n" +
			"	fmt.Println(Чтобы сделать ход, нужно подумать. После этого нужно ввести команду в формате)\n" +
			"	fmt.Println(буква-цифра-буква-цифра, например вот так: )\n" +
			"	fmt.Println(ttE7E5 или A2A3)\n" +
			"	fmt.Println(Буквы английские, обязательно заглавные(как в примере))\n" +
			"	fmt.Println(n)\n" +
			"	fmt.Println(Парочка важных моментов!)\n" +
			"	fmt.Println(К сожалению короля - невозможно сделать ракировку.)\n" +
			"	fmt.Println(Глупый автор вместо обдумывания нормальной архитектуры программы)\n" +
			"	fmt.Println(сломя голову начал писать код, поэтому сейчас представляется немножечко)\n" +
			"	fmt.Println(сложным добавить ракировку. Постараюсь добавить позже.)\n" +
			"	fmt.Println(А ещё пешка ни в кого не превращается по достижению конца доски:D)\n" +
			"	fmt.Println(Лень сейчас это доделывать! Ну и нет взятия пешки на проходе...)\n" +
			"	fmt.Println(Надеюсь это всё, что я не добавил)\n" +
			"	fmt.Println()\n" +
			"	fmt.Println(Предупреждение!)\n" +
			"	fmt.Println(Компьютер крааайне маловероятно выиграет вас матом(хе-хе)n чисто по техническим причинам)\n" +
			"	fmt.Println(НО Вашей задачей является победить этого электронного монстра.n А тут он будет сопротивляться)\n" +
			"	fmt.Println(только так! Повторю, вы играете за шоколадных.n Сегодня у меня отвратительный день,)\n" +
			"	fmt.Println(и мне не хочется,n чтобы у вас день тоже не задался, поэтому я дарю Вам...)\n" +
			"	fmt.Println()\n" +
			"	fmt.Println(...право первого хода!nПокажите этому куску железа на что способен человек!)\n" +
			"	fmt.Println(Пускай знает своё место!)\n" +
			"	fmt.Println(Успехов!)\n" +
			"	fmt.Println()\n" +
			"	fmt.Println(P.S.Если что-то непонятно - разберетёсь сами.n В крайнем случае пишите на почту.)\n" +
			"	fmt.Println(Письма я не читаю, но мне будет приятно:))\n" +
			"	fmt.Printf(nn1.Выйти в менюn)\n" +
			"	var temp string\n" +
			"	_, _ = fmt.Scan(&temp)\n" +
			"}\n" +
			"func about() {\n" +
			"	fmt.Println()\n" +
			"	Label()\n" +
			"	fmt.Printf(ttО программеnn)\n" +
			"	fmt.Println(ALICHESS version 1.0001&24sjf(^89)\n" +
			"	fmt.Println(Author : Alisher F.)\n" +
			"	fmt.Println(email : sambusa001@mail.ru)\n" +
			"	fmt.Println(nn1.Показать исходный кодn)\n" +
			"	fmt.Println(2.Выйти в меню)\n" +
			"	var choice int\n" +
			"	_, err := fmt.Scan(&choice)\n" +
			"	if err != nil {\n" +
			"		fmt.Println(ERROR , err)\n" +
			"		return\n" +
			"	}\n" +
			"\n" +
			"	if choice == 2 {\n" +
			"		return\n" +
			"	} else {\n" +
			"		showCode()\n" +
			"	}\n" +
			"}\n" +
			"func joke() {\n" +
			"	fmt.Println()\n" +
			"	Label()\n" +
			"	fmt.Printf(ttШутка про шахматыnn)\n" +
			"	fmt.Println(Мой компьютер постоянно обыгрывает меня в шахматы.)\n" +
			"	fmt.Println(-Зато я всегда побеждаю его)\n" +
			"	fmt.Println(в боксёрском поединке!)\n" +
			"	fmt.Println(n1.Выйти в меню)\n" +
			"	var temp string\n" +
			"	_, _ = fmt.Scan(&temp)\n" +
			"}\n" +
			"\n" +
			"var stringsExit = []string{\n" +
			"	Вы точно хотите выйти?,\n" +
			"	Вы абсолютно уверены, что хотите выйти?,\n" +
			"	Мне кажется вы сомневаетесь. Подумайте хорошенько.,\n" +
			"	Я понял, что Вам нужна помощь и решил всё за Вас. Не благодарите.,\n" +
			"}\n" +
			"func exitChess(i int) bool {\n" +
			"	clearWindowsConsole()\n" +
			"	fmt.Printf(nnnnnnnn)\n" +
			"	fmt.Println(stringsExit[i])\n" +
			"	if i != 3 {\n" +
			"		fmt.Printf(1.Даttt2.Нетn)\n" +
			"	} else {\n" +
			"		fmt.Printf(    ttt2.Нетn)\n" +
			"	}\n" +
			"	var choice int\n" +
			"	_, err := fmt.Scan(&choice)\n" +
			"	if err != nil {\n" +
			"		fmt.Println(ERROR , err)\n" +
			"		return false\n" +
			"	}\n" +
			"\n" +
			"	if choice == 2 {\n" +
			"		return true\n" +
			"	} else if choice == 1 {\n" +
			"		if i == 3 {\n" +
			"			fmt.Println(nЧёрт, меня раскрыли. Ладно. Ты заслужил выход.)\n" +
			"			fmt.Printf(Красный крестик сверху справа.nnttНаслаждайся.nn)\n" +
			"			for {\n" +
			"				_, _ = fmt.Scan()\n" +
			"			}\n" +
			"		}\n" +
			"		i++\n" +
			"		exitChess(i)\n" +
			"	} else {\n" +
			"		fmt.Println(Неверное число. Повторите ещё раз)\n" +
			"		exitChess(i)\n" +
			"	}\n" +
			"	return true\n" +
			"}\n" +
			"\n" +
			"func think(recur int) ComputerMove { // such like this!\n" +
			"	var bestMove ComputerMove\n" +
			"	var whiteSide bool = true\n" +
			"	var isFirst bool = true\n" +
			"\n" +
			"	helpThink(recur, whiteSide, &bestMove, 0.0, isFirst)\n" +
			"	//fmt.Println(Price: , price)\n" +
			"	return bestMove\n" +
			"}\n" +
			"func helpThink(recStep int, whiteSide bool, bestMove *ComputerMove, tempScore float64, isFirst bool) float64 {\n" +
			"\n" +
			"	var finalScore float64 = 0.0\n" +
			"	var moves stack\n" +
			"	var figures []Xy\n" +
			"	var score, bestScore float64\n" +
			"	var isLast = true\n" +
			"\n" +
			"	if whiteSide {\n" +
			"		bestScore = -1000000.0\n" +
			"	} else {\n" +
			"		bestScore = 1000000.0\n" +
			"	}\n" +
			"\n" +
			"	scan(&figures, whiteSide) // get figures for this side\n" +
			"\n" +
			"	for i := range figures { // for every figure\n" +
			"		getFiguresMoves(figures[i], &moves)\n" +
			"\n" +
			"		for !moves.empty() { // for every move\n" +
			"			move, _ := moves.top() // get move\n" +
			"			_ = moves.pop()// delete figure\n" +
			"\n" +
			"			// make move\n" +
			"			whoYouKilled := table[move.y][move.x]\n" +
			"			table[move.y][move.x] = table[figures[i].y][figures[i].x]\n" +
			"			table[figures[i].y][figures[i].x] = 0\n" +
			"\n" +
			"			score = rank(whoYouKilled) + positionScore(move, whiteSide)\n" +
			"			score += tempScore\n" +
			"\n" +
			"			if recStep > 0 {\n" +
			"				isLast = false\n" +
			"				finalScore = helpThink(recStep - 1, !whiteSide, bestMove, score,false)\n" +
			"			}\n" +
			"\n" +
			"			if !isLast {\n" +
			"				score = finalScore\n" +
			"			}\n" +
			"\n" +
			"			table[figures[i].y][figures[i].x] = table[move.y][move.x]\n" +
			"			table[move.y][move.x] = whoYouKilled\n" +
			"			/*{\n" +
			"				printTable()\n" +
			"				fmt.Println()\n" +
			"				fmt.Println(recurStep : , recStep)\n" +
			"				fmt.Println(score : , score)\n" +
			"				fmt.Println(bestScore : , bestScore)\n" +
			"				fmt.Println(figures[i].y, figures[i].x :  move.y, move.x)\n" +
			"				fmt.Println(figures[i].y, figures[i].x,  : , move.y, move.x)\n" +
			"				fmt.Println()\n" +
			"			}*/\n" +
			"			if compare(bestScore, score, whiteSide) {\n" +
			"				bestScore = score\n" +
			"				//fmt.Println(recStep: , recStep,BestScore: , bestScore)\n" +
			"				if isFirst {\n" +
			"					bestMove.fromY = figures[i].y\n" +
			"					bestMove.fromX = figures[i].x\n" +
			"					bestMove.toY = move.y\n" +
			"					bestMove.toX = move.x\n" +
			"				}\n" +
			"			}\n" +
			"		}\n" +
			"\n" +
			"	}\n" +
			"	return bestScore\n" +
			"}\n" +
			"func compare(bestScore, score float64, whiteSide bool) bool {\n" +
			"	if whiteSide {\n" +
			"		return bestScore < score\n" +
			"	} else {\n" +
			"		return bestScore > score\n" +
			"	}\n" +
			"}\n" +
			"func isBlackWin() {\n" +
			"	//TODO\n" +
			"}\n" +
			"\n" +
			"func scan(vec *[]Xy, colour bool) { // fill slice with chess figures\n" +
			"	var a Xy\n" +
			"\n" +
			"	if colour {\n" +
			"		//fmt.Println(if)\n" +
			"		for i := 0; i < SIZE; i++ {\n" +
			"			for j := 0; j < SIZE; j++ {\n" +
			"				if isWhite(table[i][j]) {\n" +
			"					//fmt.Println(a : , a)\n" +
			"					a.x = j\n" +
			"					a.y = i\n" +
			"					*vec = append(*vec, a) // WARNING\n" +
			"				}\n" +
			"			}\n" +
			"		}\n" +
			"	} else {\n" +
			"		for i := 0; i < SIZE; i++ {\n" +
			"			for j := 0; j < SIZE; j++ {\n" +
			"				if isBlack(table[i][j]) {\n" +
			"					//fmt.Println(a : , a)\n" +
			"					a.x = j\n" +
			"					a.y = i\n" +
			"					*vec = append(*vec, a) // WARNING\n" +
			"				}\n" +
			"			}\n" +
			"		}\n" +
			"	}\n" +
			"}\n" +
			"func getFiguresMoves(move Xy, steak *stack) { //fill your stack with all possible moves for the figure\n" +
			"	var pos Xy\n" +
			"	var flag bool = true\n" +
			"\n" +
			"	pos.x = move.x;\n" +
			"	pos.y = move.y;\n" +
			"	var fig int = table[pos.y][pos.x]\n" +
			"	var temp int\n" +
			"	var posT Xy\n" +
			"\n" +
			"	switch fig {\n" +
			"	case 1:\n" +
			"		posT = pos\n" +
			"\n" +
			"		if posT.y == 1 && table[posT.y+2][posT.x] == 0 && table[posT.y+1][posT.x] == 0 {\n" +
			"			posT.y += 2\n" +
			"			steak.push(posT)\n" +
			"		}\n" +
			"\n" +
			"		posT = pos;\n" +
			"		posT.y += 1;\n" +
			"		if posT.y < 8 && table[posT.y][pos.x] == 0 {\n" +
			"			steak.push(posT)\n" +
			"		}\n" +
			"\n" +
			"		if (pos.y+1) < 8 && (pos.x+1) < 8 && table[pos.y+1][pos.x+1] != 0 {\n" +
			"			temp = table[pos.y+1][pos.x+1]\n" +
			"			var flag = true\n" +
			"			for i := 0; i < 6; i++ {\n" +
			"				if temp == black[i] {\n" +
			"					flag = false\n" +
			"					break\n" +
			"				}\n" +
			"			}\n" +
			"\n" +
			"			if flag {\n" +
			"				posT.x = pos.x + 1\n" +
			"				posT.y = pos.y + 1\n" +
			"				steak.push(posT)\n" +
			"			}\n" +
			"\n" +
			"		}\n" +
			"\n" +
			"		if (pos.y+1) < 8 && (pos.x-1) >= 0 && table[pos.y+1][pos.x-1] != 0 {\n" +
			"			temp = table[pos.y+1][pos.x-1]\n" +
			"			flag = true\n" +
			"			for i := 0; i < 6; i++ {\n" +
			"				if temp == black[i] {\n" +
			"					flag = false\n" +
			"					break\n" +
			"				}\n" +
			"			}\n" +
			"\n" +
			"			if flag {\n" +
			"				posT.x = pos.x - 1\n" +
			"				posT.y = pos.y + 1\n" +
			"				steak.push(posT)\n" +
			"			}\n" +
			"\n" +
			"		}\n" +
			"	case 11:\n" +
			"		posT = pos\n" +
			"		if (posT.y == 6) && (table[posT.y-2][posT.x] == 0) && (table[posT.y-1][posT.x] == 0) {\n" +
			"			posT.y -= 2\n" +
			"			steak.push(posT)\n" +
			"		}\n" +
			"		posT = pos\n" +
			"		posT.y -= 1\n" +
			"\n" +
			"		if (posT.y >= 0) && (table[posT.y][pos.x] == 0) {\n" +
			"			steak.push(posT);\n" +
			"		}\n" +
			"\n" +
			"		if (pos.y-1) >= 0 && (pos.x-1) >= 0 && table[pos.y-1][pos.x-1] != 0 {\n" +
			"			temp = table[pos.y-1][pos.x-1]\n" +
			"			flag = true\n" +
			"			for i := 0; i < 6; i++ {\n" +
			"				if temp == white[i] {\n" +
			"					flag = false\n" +
			"					break\n" +
			"				}\n" +
			"			}\n" +
			"\n" +
			"			if flag {\n" +
			"				posT.x = pos.x - 1\n" +
			"				posT.y = pos.y - 1\n" +
			"				steak.push(posT)\n" +
			"			}\n" +
			"		}\n" +
			"\n" +
			"		if (pos.y-1) >= 0 && (pos.x+1) < 8 && table[pos.y-1][pos.x+1] != 0 {\n" +
			"			temp = table[pos.y-1][pos.x+1]\n" +
			"			flag = true\n" +
			"			for i := 0; i < 6; i++ {\n" +
			"				if temp == white[i] {\n" +
			"					flag = false\n" +
			"					break\n" +
			"				}\n" +
			"			}\n" +
			"\n" +
			"			if flag {\n" +
			"				posT.x = pos.x + 1\n" +
			"				posT.y = pos.y - 1\n" +
			"				steak.push(posT)\n" +
			"			}\n" +
			"		}\n" +
			"	case 2:\n" +
			"		posT = pos\n" +
			"		posT.x++\n" +
			"		posT.y -= 2\n" +
			"\n" +
			"		horseB(posT, temp, flag, steak)\n" +
			"\n" +
			"		posT = pos\n" +
			"		posT.x += 2\n" +
			"		posT.y--\n" +
			"\n" +
			"		horseB(posT, temp, flag, steak);\n" +
			"		posT = pos;\n" +
			"		posT.x += 2;\n" +
			"		posT.y++;\n" +
			"		horseB(posT, temp, flag, steak);\n" +
			"		posT = pos;\n" +
			"		posT.x++;\n" +
			"		posT.y += 2;\n" +
			"		horseB(posT, temp, flag, steak);\n" +
			"		posT = pos;\n" +
			"		posT.x--;\n" +
			"		posT.y += 2;\n" +
			"		horseB(posT, temp, flag, steak);\n" +
			"		posT = pos;\n" +
			"		posT.x -= 2;\n" +
			"		posT.y++;\n" +
			"		horseB(posT, temp, flag, steak);\n" +
			"		posT = pos;\n" +
			"		posT.x -= 2;\n" +
			"		posT.y--;\n" +
			"		horseB(posT, temp, flag, steak);\n" +
			"		posT = pos;\n" +
			"		posT.x--;\n" +
			"		posT.y -= 2;\n" +
			"		horseB(posT, temp, flag, steak);\n" +
			"	case 21:\n" +
			"		posT = pos;\n" +
			"		posT.x++;\n" +
			"		posT.y -= 2;\n" +
			"		horse(posT, temp, flag, steak);\n" +
			"		posT = pos;\n" +
			"		posT.x += 2;\n" +
			"		posT.y--;\n" +
			"		horse(posT, temp, flag, steak);\n" +
			"		posT = pos;\n" +
			"		posT.x += 2;\n" +
			"		posT.y++;\n" +
			"		horse(posT, temp, flag, steak);\n" +
			"		posT = pos;\n" +
			"		posT.x++;\n" +
			"		posT.y += 2;\n" +
			"		horse(posT, temp, flag, steak);\n" +
			"		posT = pos;\n" +
			"		posT.x--;\n" +
			"		posT.y += 2;\n" +
			"		horse(posT, temp, flag, steak);\n" +
			"		posT = pos;\n" +
			"		posT.x -= 2;\n" +
			"		posT.y++;\n" +
			"		horse(posT, temp, flag, steak);\n" +
			"		posT = pos;\n" +
			"		posT.x -= 2;\n" +
			"		posT.y--;\n" +
			"		horse(posT, temp, flag, steak);\n" +
			"		posT = pos;\n" +
			"		posT.x--;\n" +
			"		posT.y -= 2;\n" +
			"		horse(posT, temp, flag, steak)\n" +
			"	case 3: // WARNING! '{' WAS AFTER CASE 3\n" +
			"		br := true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			posT.x++;\n" +
			"			posT.y++;\n" +
			"			br = elefB(posT, temp, flag, steak);\n" +
			"		}\n" +
			"\n" +
			"		br = true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			posT.x++;\n" +
			"			posT.y--;\n" +
			"			br = elefB(posT, temp, flag, steak);\n" +
			"		}\n" +
			"\n" +
			"		br = true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			posT.x--;\n" +
			"			posT.y++;\n" +
			"			br = elefB(posT, temp, flag, steak);\n" +
			"		}\n" +
			"\n" +
			"		br = true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			posT.x--;\n" +
			"			posT.y--;\n" +
			"			br = elefB(posT, temp, flag, steak);\n" +
			"		}\n" +
			"	case 31:\n" +
			"		br := true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			posT.x++;\n" +
			"			posT.y++;\n" +
			"			br = elef(posT, temp, flag, steak);\n" +
			"		}\n" +
			"\n" +
			"		br = true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			posT.x++;\n" +
			"			posT.y--;\n" +
			"			br = elef(posT, temp, flag, steak);\n" +
			"		}\n" +
			"\n" +
			"		br = true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			posT.x--;\n" +
			"			posT.y++;\n" +
			"			br = elef(posT, temp, flag, steak);\n" +
			"		}\n" +
			"\n" +
			"		br = true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			posT.x--;\n" +
			"			posT.y--;\n" +
			"			br = elef(posT, temp, flag, steak);\n" +
			"		}\n" +
			"	case 5:\n" +
			"		br := true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			////cout << ER1n;\n" +
			"			posT.x++;\n" +
			"			br = elefB(posT, temp, flag, steak);\n" +
			"		}\n" +
			"\n" +
			"		br = true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			////cout << posT.y <<  YYYn;\n" +
			"			posT.y++;\n" +
			"			br = elefB(posT, temp, flag, steak);\n" +
			"		}\n" +
			"\n" +
			"		br = true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			////cout << ER3n;\n" +
			"			posT.x--;\n" +
			"			br = elefB(posT, temp, flag, steak);\n" +
			"		}\n" +
			"\n" +
			"		br = true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			////cout << ER4n;\n" +
			"			posT.y--;\n" +
			"			br = elefB(posT, temp, flag, steak);\n" +
			"		}\n" +
			"	case 51:\n" +
			"		br := true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			////cout << ER1n;\n" +
			"			posT.x++;\n" +
			"			br = elef(posT, temp, flag, steak);\n" +
			"		}\n" +
			"\n" +
			"		br = true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			////cout << posT.y <<  YYYn;\n" +
			"			posT.y++;\n" +
			"			br = elef(posT, temp, flag, steak);\n" +
			"		}\n" +
			"\n" +
			"		br = true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			////cout << ER3n;\n" +
			"			posT.x--;\n" +
			"			br = elef(posT, temp, flag, steak);\n" +
			"		}\n" +
			"\n" +
			"		br = true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			////cout << ER4n;\n" +
			"			posT.y--;\n" +
			"			br = elef(posT, temp, flag, steak);\n" +
			"		}\n" +
			"	case 10:\n" +
			"		br := true\n" +
			"		posT = pos;\n" +
			"		////cout << !n;\n" +
			"		for br {\n" +
			"			////cout << ER1n;\n" +
			"			posT.x++;\n" +
			"			br = elefB(posT, temp, flag, steak);\n" +
			"		}\n" +
			"\n" +
			"		br = true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			////cout << posT.y <<  YYYn;\n" +
			"			posT.y++;\n" +
			"			br = elefB(posT, temp, flag, steak);\n" +
			"		}\n" +
			"\n" +
			"		br = true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			////cout << ER3n;\n" +
			"			posT.x--;\n" +
			"			br = elefB(posT, temp, flag, steak);\n" +
			"		}\n" +
			"\n" +
			"		br = true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			////cout << ER4n;\n" +
			"			posT.y--;\n" +
			"			br = elefB(posT, temp, flag, steak);\n" +
			"		}\n" +
			"\n" +
			"		br = true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			posT.x++;\n" +
			"			posT.y++;\n" +
			"			br = elefB(posT, temp, flag, steak);\n" +
			"		}\n" +
			"\n" +
			"		br = true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			posT.x++;\n" +
			"			posT.y--;\n" +
			"			br = elefB(posT, temp, flag, steak);\n" +
			"		}\n" +
			"\n" +
			"		br = true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			posT.x--;\n" +
			"			posT.y++;\n" +
			"			br = elefB(posT, temp, flag, steak);\n" +
			"		}\n" +
			"\n" +
			"		br = true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			posT.x--;\n" +
			"			posT.y--;\n" +
			"			br = elefB(posT, temp, flag, steak);\n" +
			"		}\n" +
			"	case 101:\n" +
			"		br := true\n" +
			"		posT = pos;\n" +
			"		////cout << !n;\n" +
			"		for br {\n" +
			"			////cout << ER1n;\n" +
			"			posT.x++;\n" +
			"			br = elef(posT, temp, flag, steak);\n" +
			"		}\n" +
			"\n" +
			"		br = true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			////cout << posT.y <<  YYYn;\n" +
			"			posT.y++;\n" +
			"			br = elef(posT, temp, flag, steak);\n" +
			"		}\n" +
			"\n" +
			"		br = true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			////cout << ER3n;\n" +
			"			posT.x--;\n" +
			"			br = elef(posT, temp, flag, steak);\n" +
			"		}\n" +
			"\n" +
			"		br = true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			////cout << ER4n;\n" +
			"			posT.y--;\n" +
			"			br = elef(posT, temp, flag, steak);\n" +
			"		}\n" +
			"\n" +
			"		br = true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			posT.x++;\n" +
			"			posT.y++;\n" +
			"			br = elef(posT, temp, flag, steak);\n" +
			"		}\n" +
			"\n" +
			"		br = true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			posT.x++;\n" +
			"			posT.y--;\n" +
			"			br = elef(posT, temp, flag, steak);\n" +
			"		}\n" +
			"\n" +
			"		br = true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			posT.x--;\n" +
			"			posT.y++;\n" +
			"			br = elef(posT, temp, flag, steak);\n" +
			"		}\n" +
			"\n" +
			"		br = true\n" +
			"		posT = pos;\n" +
			"		for br {\n" +
			"			posT.x--;\n" +
			"			posT.y--;\n" +
			"			br = elef(posT, temp, flag, steak);\n" +
			"		}\n" +
			"	case 99:\n" +
			"		posT = pos;\n" +
			"		posT.x++;\n" +
			"		horseB(posT, temp, flag, steak);\n" +
			"		posT = pos;\n" +
			"		posT.x--;\n" +
			"		horseB(posT, temp, flag, steak);\n" +
			"		posT = pos;\n" +
			"		posT.y++;\n" +
			"		horseB(posT, temp, flag, steak);\n" +
			"		posT = pos;\n" +
			"		posT.y--;\n" +
			"		horseB(posT, temp, flag, steak);\n" +
			"		posT = pos;\n" +
			"		posT.x++;\n" +
			"		posT.y++;\n" +
			"		horseB(posT, temp, flag, steak);\n" +
			"		posT = pos;\n" +
			"		posT.x++;\n" +
			"		posT.y--;\n" +
			"		horseB(posT, temp, flag, steak);\n" +
			"		posT = pos;\n" +
			"		posT.x--;\n" +
			"		posT.y++;\n" +
			"		horseB(posT, temp, flag, steak);\n" +
			"		posT = pos;\n" +
			"		posT.x--;\n" +
			"		posT.y--;\n" +
			"		horseB(posT, temp, flag, steak);\n" +
			"	case 990:\n" +
			"		//cout << KING: ; //cout << pos.x << ' ' << pos.y << endl;\n" +
			"		posT = pos;\n" +
			"		posT.x++;\n" +
			"		horse(posT, temp, flag, steak);\n" +
			"		posT = pos;\n" +
			"		posT.x--;\n" +
			"		horse(posT, temp, flag, steak);\n" +
			"		posT = pos;\n" +
			"		posT.y++;\n" +
			"		horse(posT, temp, flag, steak);\n" +
			"		posT = pos;\n" +
			"		posT.y--;\n" +
			"		horse(posT, temp, flag, steak);\n" +
			"		posT = pos;\n" +
			"		posT.x++;\n" +
			"		posT.y++;\n" +
			"		horse(posT, temp, flag, steak);\n" +
			"		posT = pos;\n" +
			"		posT.x++;\n" +
			"		posT.y--;\n" +
			"		horse(posT, temp, flag, steak);\n" +
			"		posT = pos;\n" +
			"		posT.x--;\n" +
			"		posT.y++;\n" +
			"		horse(posT, temp, flag, steak);\n" +
			"		posT = pos;\n" +
			"		posT.x--;\n" +
			"		posT.y--;\n" +
			"		horse(posT, temp, flag, steak);\n" +
			"	default:\n" +
			"		fmt.Println(FUCKING PROGER!)\n" +
			"		fmt.Printf(v v x yn, move.x, move.y)\n" +
			"		fmt.Printf(v FIGn, table[move.y][move.x])\n" +
			"	} // switch end\n" +
			"}\n" +
			"//TODO ракировка!\n" +
			"/*func rakirovkaWhite(moves *stack) {\n" +
			"	if didWhiteKingMove {\n" +
			"		return\n" +
			"	}\n" +
			"	if table[0][7] == 0 && table[1][7] == 0 && table[2][7] == 0 && table[3][7] == 0 {\n" +
			"\n" +
			"	}\n" +
			"}\n" +
			"func rakirovkaBlack(moves *stack) {\n" +
			"	if didBlackKingMove {\n" +
			"		return\n" +
			"	}\n" +
			"}\n" +
			"var didWhiteKingMove bool\n" +
			"var didBlackKingMove bool*/\n" +
			"func elef(posT Xy, temp int, flag bool, steak *stack) bool {\n" +
			"	if posT.x >= 0 && posT.x < 8 && posT.y >= 0 && posT.y < 8 {\n" +
			"		temp = table[posT.y][posT.x]\n" +
			"		flag = true\n" +
			"		for i := 0; i < 6; i++ {\n" +
			"			if temp == white[i] {\n" +
			"				return false\n" +
			"			}\n" +
			"			if temp == black[i] {\n" +
			"				steak.push(posT)\n" +
			"				return false\n" +
			"			}\n" +
			"		}\n" +
			"		if flag {\n" +
			"			steak.push(posT)\n" +
			"		}\n" +
			"		return true\n" +
			"	}\n" +
			"	return false\n" +
			"}\n" +
			"func elefB(posT Xy, temp int, flag bool, steak *stack) bool {\n" +
			"	if posT.x >= 0 && posT.x < 8 && posT.y >= 0 && posT.y < 8 {\n" +
			"		temp = table[posT.y][posT.x]\n" +
			"		flag = true\n" +
			"		for i := 0; i < 6; i++ {\n" +
			"			if temp == black[i] {\n" +
			"				return false\n" +
			"			}\n" +
			"			if temp == white[i] {\n" +
			"				steak.push(posT)\n" +
			"				return false\n" +
			"			}\n" +
			"		}\n" +
			"		if flag {\n" +
			"			steak.push(posT)\n" +
			"		}\n" +
			"		return true\n" +
			"	}\n" +
			"	return false\n" +
			"}\n" +
			"func horse(posT Xy, temp int, flag bool, steak *stack) {\n" +
			"	if posT.x >= 0 && posT.x < 8 && posT.y >= 0 && posT.y < 8 {\n" +
			"		temp = table[posT.y][posT.x]\n" +
			"		flag = true\n" +
			"		for i := 0; i < 6; i++ {\n" +
			"			if temp == white[i] {\n" +
			"				flag = false\n" +
			"				break\n" +
			"			}\n" +
			"		}\n" +
			"		if flag {\n" +
			"			steak.push(posT)\n" +
			"		}\n" +
			"	}\n" +
			"}\n" +
			"func horseB(posT Xy, temp int, flag bool, steak *stack) {\n" +
			"	if posT.x >= 0 && posT.x < 8 && posT.y >= 0 && posT.y < 8 {\n" +
			"		temp = table[posT.y][posT.x]\n" +
			"		flag = true\n" +
			"		for i := 0; i < 6; i++ {\n" +
			"			if temp == black[i] {\n" +
			"				flag = false\n" +
			"				break\n" +
			"			}\n" +
			"		}\n" +
			"		if flag {\n" +
			"			steak.push(posT)\n" +
			"		}\n" +
			"	}\n" +
			"}\n" +
			"\n" +
			"func printTable() {\n" +
			"	var figure byte\n" +
			"	fmt.Println()\n" +
			"	for i := 0; i < SIZE; i++ {\n" +
			"		fmt.Printf(			v| , numbers[i])\n" +
			"		for j := 0; j < SIZE; j++ {\n" +
			"			switch table[i][j] {\n" +
			"			case 1:\n" +
			"				figure = 'p';\n" +
			"			case 2:\n" +
			"				figure = 'k';\n" +
			"			case 3:\n" +
			"				figure = 's';\n" +
			"			case 5:\n" +
			"				figure = 'l';\n" +
			"			case 10:\n" +
			"				figure = 'f';\n" +
			"			case 99:\n" +
			"				figure = 'i';\n" +
			"			case 11:\n" +
			"				figure = 'P';\n" +
			"			case 21:\n" +
			"				figure = 'K';\n" +
			"			case 31:\n" +
			"				figure = 'S';\n" +
			"			case 51:\n" +
			"				figure = 'L';\n" +
			"			case 101:\n" +
			"				figure = 'F';\n" +
			"			case 990:\n" +
			"				figure = 'I';\n" +
			"			case 7:\n" +
			"				figure = '*';\n" +
			"			default:\n" +
			"				figure = '.';\n" +
			"			}\n" +
			"			fmt.Printf(  v, string(figure))\n" +
			"		}\n" +
			"		fmt.Println()\n" +
			"	}\n" +
			"	fmt.Printf(			     ______________________n)\n" +
			"	fmt.Printf(			     A  B  C  D  E  F  G  Hnn)\n" +
			"}\n" +
			"func makeMove(move string) bool {\n" +
			"	if move[0] == 'Z' {\n" +
			"		return false\n" +
			"	}\n" +
			"	if len(move) < 4 {\n" +
			"		return true\n" +
			"	}\n" +
			"	var num int\n" +
			"	var pos Pos\n" +
			"\n" +
			"	switch move[0] {\n" +
			"	case 'A':\n" +
			"		num = 0;\n" +
			"	case 'B':\n" +
			"		num = 1;\n" +
			"	case 'C':\n" +
			"		num = 2;\n" +
			"	case 'D':\n" +
			"		num = 3;\n" +
			"	case 'E':\n" +
			"		num = 4;\n" +
			"	case 'F':\n" +
			"		num = 5;\n" +
			"	case 'G':\n" +
			"		num = 6;\n" +
			"	case 'H':\n" +
			"		num = 7;\n" +
			"	default:\n" +
			"		return true\n" +
			"	}\n" +
			"\n" +
			"	pos.sx = num\n" +
			"\n" +
			"	temp := int(move[1]) - 49\n" +
			"	if temp < 1 || temp > 8 {\n" +
			"		return true\n" +
			"	}\n" +
			"	pos.sy = temp\n" +
			"\n" +
			"	switch pos.sy {\n" +
			"	case 0:\n" +
			"		pos.sy = 7;\n" +
			"	case 1:\n" +
			"		pos.sy = 6;\n" +
			"	case 2:\n" +
			"		pos.sy = 5;\n" +
			"	case 3:\n" +
			"		pos.sy = 4;\n" +
			"	case 4:\n" +
			"		pos.sy = 3;\n" +
			"	case 5:\n" +
			"		pos.sy = 2;\n" +
			"	case 6:\n" +
			"		pos.sy = 1;\n" +
			"	case 7:\n" +
			"		pos.sy = 0;\n" +
			"	}\n" +
			"\n" +
			"	switch move[2] {\n" +
			"	case 'A':\n" +
			"		num = 0;\n" +
			"	case 'B':\n" +
			"		num = 1;\n" +
			"	case 'C':\n" +
			"		num = 2;\n" +
			"	case 'D':\n" +
			"		num = 3;\n" +
			"	case 'E':\n" +
			"		num = 4;\n" +
			"	case 'F':\n" +
			"		num = 5;\n" +
			"	case 'G':\n" +
			"		num = 6;\n" +
			"	case 'H':\n" +
			"		num = 7;\n" +
			"	default:\n" +
			"		return true\n" +
			"	}\n" +
			"\n" +
			"	pos.ex = num\n" +
			"\n" +
			"	temp = int(move[3]) - 49\n" +
			"	if temp < 1 || temp > 8 {\n" +
			"		return true\n" +
			"	}\n" +
			"	pos.ey = temp\n" +
			"\n" +
			"	switch pos.ey {\n" +
			"	case 0:\n" +
			"		pos.ey = 7;\n" +
			"	case 1:\n" +
			"		pos.ey = 6;\n" +
			"	case 2:\n" +
			"		pos.ey = 5;\n" +
			"	case 3:\n" +
			"		pos.ey = 4;\n" +
			"	case 4:\n" +
			"		pos.ey = 3;\n" +
			"	case 5:\n" +
			"		pos.ey = 2;\n" +
			"	case 6:\n" +
			"		pos.ey = 1;\n" +
			"	case 7:\n" +
			"		pos.ey = 0;\n" +
			"\n" +
			"	}\n" +
			"\n" +
			"	var moves stack\n" +
			"	var moveTemp Xy\n" +
			"	var userMove = Xy{pos.ex, pos.ey}\n" +
			"	getFiguresMoves(Xy{pos.sx, pos.sy}, &moves)\n" +
			"	flag := false\n" +
			"	for !moves.empty() {\n" +
			"		moveTemp, _ = moves.top()\n" +
			"		moves.pop()\n" +
			"		if moveTemp == userMove{\n" +
			"			flag = true\n" +
			"		}\n" +
			"	}\n" +
			"	if flag {\n" +
			"		table[pos.ey][pos.ex] = table[pos.sy][pos.sx]\n" +
			"		table[pos.sy][pos.sx] = 0\n" +
			"		return false\n" +
			"	}\n" +
			"	return true\n" +
			"}\n" +
			"func isWhite(fig int) bool {\n" +
			"	for i := 0; i < 6; i++ {\n" +
			"		//fmt.Println(for)\n" +
			"		if fig == white[i] {\n" +
			"			return true\n" +
			"		}\n" +
			"	}\n" +
			"	return false\n" +
			"}\n" +
			"func isBlack(fig int) bool {\n" +
			"	for i := 0; i < 6; i++ {\n" +
			"		if fig == black[i] {\n" +
			"			return true\n" +
			"		}\n" +
			"	}\n" +
			"	return false\n" +
			"}\n" +
			"\n" +
			"func rank(fig int) float64 { // thanks to my stupidness this func returns rank only for white side :D\n" +
			"	switch fig {\n" +
			"	case 0:\n" +
			"		return 0;\n" +
			"\n" +
			"	case 11:\n" +
			"		return -1;\n" +
			"	case 1:\n" +
			"		return 1;\n" +
			"\n" +
			"	case 21:\n" +
			"		return -2;\n" +
			"	case 2:\n" +
			"		return 2;\n" +
			"\n" +
			"	case 31:\n" +
			"		return -3;\n" +
			"	case 3:\n" +
			"		return 3;\n" +
			"\n" +
			"	case 51:\n" +
			"		return -5;\n" +
			"	case 5:\n" +
			"		return 5;\n" +
			"\n" +
			"	case 101:\n" +
			"		return -10;\n" +
			"	case 10:\n" +
			"		return 10;\n" +
			"\n" +
			"	case 990:\n" +
			"		return -9990;\n" +
			"	case 99:\n" +
			"		return 1000;\n" +
			"\n" +
			"	default:\n" +
			"		fmt.Println(RANKer)\n" +
			"	}\n" +
			"	return -9876543.21\n" +
			"}\n" +
			"/*var cofficientWhite = [][]float64 {\n" +
			"	{0.0000071, 0.0000071, 0.0000071, 0.0000071, 0.0000071, 0.0000071, 0.0000071, 0.0000071},\n" +
			"	{0.0000061, 0.0000062, 0.0000062, 0.0000062, 0.0000062, 0.0000062, 0.0000062, 0.0000061},\n" +
			"	{0.0000051, 0.0000052, 0.0000053, 0.0000053, 0.0000053, 0.0000053, 0.0000052, 0.0000051},\n" +
			"	{0.0000041, 0.0000042, 0.0000043, 0.0000044, 0.0000044, 0.0000043, 0.0000042, 0.0000041},\n" +
			"	{0.0000031, 0.0000032, 0.0000033, 0.0000034, 0.0000034, 0.0000033, 0.0000032, 0.0000031},\n" +
			"	{0.0000021, 0.0000022, 0.0000023, 0.0000023, 0.0000023, 0.0000023, 0.0000022, 0.0000021},\n" +
			"	{0.0000011, 0.0000012, 0.0000012, 0.0000012, 0.0000012, 0.0000012, 0.0000012, 0.0000011},\n" +
			"	{0.0000001, 0.0000001, 0.0000001, 0.0000001, 0.0000001, 0.0000001, 0.0000001, 0.0000001},\n" +
			"}\n" +
			"var cofficientBlack = [][]float64 {\n" +
			"	{0.0000001, 0.0000001, 0.0000001, 0.0000001, 0.0000001, 0.0000001, 0.0000001, 0.0000001},\n" +
			"	{0.0000011, 0.0000012, 0.0000012, 0.0000012, 0.0000012, 0.0000012, 0.0000012, 0.0000011},\n" +
			"	{0.0000021, 0.0000022, 0.0000023, 0.0000023, 0.0000023, 0.0000023, 0.0000022, 0.0000021},\n" +
			"	{0.0000031, 0.0000032, 0.0000033, 0.0000034, 0.0000034, 0.0000033, 0.0000032, 0.0000031},\n" +
			"	{0.0000041, 0.0000042, 0.0000043, 0.0000044, 0.0000044, 0.0000043, 0.0000042, 0.0000041},\n" +
			"	{0.0000051, 0.0000052, 0.0000053, 0.0000053, 0.0000053, 0.0000053, 0.0000052, 0.0000051},\n" +
			"	{0.0000061, 0.0000062, 0.0000062, 0.0000062, 0.0000062, 0.0000062, 0.0000062, 0.0000061},\n" +
			"	{0.0000071, 0.0000071, 0.0000071, 0.0000071, 0.0000071, 0.0000071, 0.0000071, 0.0000071},\n" +
			"}*/\n" +
			"var cofficientWhite = [][]float64 {\n" +
			"	{0.071, 0.071, 0.071, 0.071, 0.071, 0.071, 0.071, 0.071},\n" +
			"	{0.061, 0.062, 0.062, 0.062, 0.062, 0.062, 0.062, 0.061},\n" +
			"	{0.051, 0.052, 0.053, 0.053, 0.053, 0.053, 0.052, 0.051},\n" +
			"	{0.041, 0.042, 0.043, 0.044, 0.044, 0.043, 0.042, 0.041},\n" +
			"	{0.031, 0.032, 0.033, 0.033, 0.033, 0.033, 0.032, 0.031},\n" +
			"	{0.021, 0.022, 0.023, 0.023, 0.023, 0.023, 0.022, 0.021},\n" +
			"	{0.011, 0.012, 0.012, 0.012, 0.012, 0.012, 0.012, 0.012},\n" +
			"	{0.001, 0.001, 0.001, 0.001, 0.001, 0.001, 0.001, 0.001},\n" +
			"}\n" +
			"var cofficientBlack = [][]float64 {\n" +
			"	{0.001, 0.001, 0.001, 0.001, 0.001, 0.001, 0.001, 0.001},\n" +
			"	{0.011, 0.012, 0.012, 0.012, 0.012, 0.012, 0.012, 0.012},\n" +
			"	{0.021, 0.022, 0.023, 0.023, 0.023, 0.023, 0.022, 0.021},\n" +
			"	{0.031, 0.032, 0.033, 0.033, 0.033, 0.033, 0.032, 0.031},\n" +
			"	{0.041, 0.042, 0.043, 0.044, 0.044, 0.043, 0.042, 0.041},\n" +
			"	{0.051, 0.052, 0.053, 0.053, 0.053, 0.053, 0.052, 0.051},\n" +
			"	{0.061, 0.062, 0.062, 0.062, 0.062, 0.062, 0.062, 0.061},\n" +
			"	{0.071, 0.071, 0.071, 0.071, 0.071, 0.071, 0.071, 0.071},\n" +
			"}\n" +
			"func positionScore(p Xy, whiteColor bool) float64 {\n" +
			"	var st stack\n" +
			"	getFiguresMoves(p, &st)\n" +
			"	var temp Xy\n" +
			"	var answer float64 = 0.0\n" +
			"\n" +
			"	for !st.empty() {\n" +
			"		temp, _ = st.top()\n" +
			"		_ = st.pop()\n" +
			"		switch table[temp.y][temp.x] {\n" +
			"		case 0:\n" +
			"			answer += 0;\n" +
			"\n" +
			"		case 11:\n" +
			"			answer += -0.001;\n" +
			"		case 1:\n" +
			"			answer += 0.001;\n" +
			"\n" +
			"		case 21:\n" +
			"			answer += -0.002;\n" +
			"		case 2:\n" +
			"			answer += 0.002;\n" +
			"\n" +
			"		case 31:\n" +
			"			answer += -0.003;\n" +
			"		case 3:\n" +
			"			answer += 0.003;\n" +
			"\n" +
			"		case 51:\n" +
			"			answer += -0.005;\n" +
			"		case 5:\n" +
			"			answer += 0.005;\n" +
			"\n" +
			"		case 101:\n" +
			"			answer += -0.01;\n" +
			"		case 10:\n" +
			"			answer += 0.01;\n" +
			"\n" +
			"		case 990:\n" +
			"			answer += -0.1;\n" +
			"		case 99:\n" +
			"			answer += 0.1;\n" +
			"\n" +
			"		default:\n" +
			"			fmt.Println(rankER_POSIT)\n" +
			"		}\n" +
			"	}\n" +
			"	//answer += isChecked(p)\n" +
			"	if whiteColor {\n" +
			"		answer += cofficientWhite[p.y][p.x]\n" +
			"	} else {\n" +
			"		answer -= cofficientBlack[p.y][p.x]\n" +
			"	}\n" +
			"\n" +
			"	return answer\n" +
			"}\n" +
			"\n" +
			"/* unused functions. maybe can be used in future */\n" +
			"\n" +
			"func isChecked(p Xy) float64 {\n" +
			"	var answer = 0.0\n" +
			"	var vec []Xy\n" +
			"	var s stack\n" +
			"	var col bool = isBlack(table[p.y][p.x])\n" +
			"	scan(&vec, col)\n" +
			"	var v Xy\n" +
			"\n" +
			"	for i := 0; i < len(vec); i++ {\n" +
			"		getFiguresMoves(vec[i], &s)\n" +
			"	}\n" +
			"	for !s.empty() {\n" +
			"		v, _ = s.top();\n" +
			"		s.pop();\n" +
			"		if v.x == p.x && v.y == p.y {\n" +
			"			if col {\n" +
			"				answer += 0.11\n" +
			"			} else {\n" +
			"				answer -= 0.11\n" +
			"			}\n" +
			"		}\n" +
			"	}\n" +
			"	return answer\n" +
			"}\n" +
			"func isAttacked(position Xy, whiteColor bool) bool {\n" +
			"\n" +
			"	var figures []Xy\n" +
			"	var moves stack\n" +
			"	var temp Xy\n" +
			"	scan(&figures, !whiteColor)\n" +
			"\n" +
			"	for i := range figures {\n" +
			"		getFiguresMoves(figures[i], &moves)\n" +
			"\n" +
			"		for !moves.empty() {\n" +
			"			temp, _ = moves.top()\n" +
			"			_ = moves.pop()\n" +
			"			if position.x == temp.x && position.y == temp.y {\n" +
			"				return true\n" +
			"			}\n" +
			"		}\n" +
			"	}\n" +
			"	return false\n" +
			"}\n" +
			"func printV(v []Xy) {\n" +
			"	fmt.Println(PRINT)\n" +
			"	if len(v) > 0 {\n" +
			"		fmt.Println(V is Empty)\n" +
			"	}\n" +
			"	for i := 0; i < len(v); i++ {\n" +
			"		fmt.Printf(v -- v , v[i].x, v[i].y)\n" +
			"	}\n" +
			"	fmt.Println()\n" +
			"}\n" +
			"func printS(s stack) {\n" +
			"	if s.size() == 0 {\n" +
			"		fmt.Println(V is Empty)\n" +
			"	}\n" +
			"\n" +
			"	var temp Xy\n" +
			"\n" +
			"	for !s.empty() {\n" +
			"		temp, _ = s.top()\n" +
			"		s.pop()\n" +
			"		fmt.Printf(v : v , temp.x, temp.y)\n" +
			"	}\n" +
			"	fmt.Println()\n" +
			"}\n" +
			"\n" +
			"func showCode() {\n" +
			"\n" +
			"}\n\n")

	fmt.Println("1.Выйти в меню")
	var choice string
	fmt.Scan(&choice)
}
