package menu

import (
	"fmt"
	"temp/chess"
	_ "temp/chess"
	"temp/chess/showCode"
	"temp/cmd"
	"temp/game"
)

var play = game.Play
var settingsFunc = game.SettingsFunc
type Menu struct {
	choice int
}

func (m *Menu) Print () {
	fmt.Printf("\t\tМеню\n\n")
	fmt.Println("1.Играть")
	fmt.Println("2.Настройки")
	fmt.Println("3.Правила")
	fmt.Println("4.О программе")
	fmt.Println("5.Выход")
	fmt.Println("6.Шутка про шахматы")
}
func (m *Menu) InputChoice () {
	_, err := fmt.Scan(&m.choice)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
}
func (m *Menu) SwitchChoice() bool {
	cmd.Clear()
	switch m.choice {
	case 1:
		play()
	case 2:
		settingsFunc()
	case 3:
		rules()
	case 4:
		about()
	case 5:
		return exitChess(0)
	case 6:
		joke()
	}
	return true
}

func rules() {
	cmd.Clear()
	fmt.Println()
	chess.Label()
	fmt.Printf("\t\tПравила\n\n")
	fmt.Println("Как играть в шахматы? Загуг... Заяндекси:)\n(Скрытая реклама яндекса. Ой! Уже не скрытая).")
	fmt.Printf("\n\n\t\tУправление\n\n")
	fmt.Println("Вы играете за черных. Ваши фигуры - те, которые сверху. Строчные буквы.")
	fmt.Println("Чтобы сделать ход, нужно подумать. После этого нужно ввести команду в формате")
	fmt.Println("буква-цифра-буква-цифра, например вот так: ")
	fmt.Println("\t\tE7E5 или A2A3")
	fmt.Println("Буквы английские, обязательно заглавные(как в примере).")
	fmt.Println("")
	fmt.Println("\n\t\tПарочка важных моментов!\n")
	fmt.Println("Настройки сбрасываются при выходе из программы.")
	fmt.Println("К сожалению короля - невозможно сделать ракировку.")
	fmt.Println("Глупый автор вместо обдумывания нормальной архитектуры программы")
	fmt.Println("сломя голову начал писать код, поэтому сейчас представляется немножечко")
	fmt.Println("сложным добавить ракировку. Постараюсь добавить позже.")
	fmt.Println("Выигрыш или проигрыш состоится только после съедения короля.")
	fmt.Println("Нет ничьи после повтора трёх ходов. Нет ничьи вовсе.")
	fmt.Println("А ещё пешка ни в кого не превращается по достижению конца доски:D")
	fmt.Println("Лень сейчас это доделывать! Ну и нет взятия пешки на проходе...")
	fmt.Println("Надеюсь это всё, что я не добавил.")
	fmt.Println()
	fmt.Println("\n\t\tПредупреждение!\n")
	fmt.Println("Компьютер крааайне маловероятно выиграет вас матом(хе-хе)\nчисто по техническим причинам,")
	fmt.Println("НО Вашей задачей является победить этого электронного монстра.\nА тут он будет сопротивляться")
	fmt.Println("только так! Повторю, вы играете за шоколадных.\nСегодня у меня отвратительный день,")
	fmt.Println("и мне не хочется,\nчтобы у вас день тоже не задался, поэтому я дарю Вам...")
	fmt.Println("...право первого хода!\nПокажите этому куску железа на что способен человек!")
	fmt.Println("Пускай знает своё место!")
	fmt.Println("Успехов!")
	fmt.Println()
	fmt.Println("P.S.Если что-то непонятно - разберетёсь сами.\nВ крайнем случае пишите на почту.")
	fmt.Println("Письма я не читаю, но мне будет приятно:)")
	fmt.Printf("\n\n1.Выйти в меню\n")
	var temp string
	_, _ = fmt.Scan(&temp)
}
func about() {
	fmt.Println()
	chess.Label()
	fmt.Printf("\t\tО программе\n\n")
	fmt.Println("Хотел назвать программу ALICHESS, но это ассоциируется")
	fmt.Println("с одной китайской компанией, поэтому назвал в честь")
	fmt.Println("Винни-пуха :)")
	fmt.Println()
	fmt.Println("В игре есть несколько пасхалок. И да, попробуйте ввести ваше имя - ")
	fmt.Println("вдруг вам повезёт:)")
	fmt.Println()
	fmt.Println("'WIN-ME-pooh' program version 1.0001&24sjf(^89")
	fmt.Println("Author : Alisher F.")
	fmt.Println("email : sambusa001@mail.ru")
	fmt.Println("\n\n1.Показать исходный код\n")
	fmt.Println("2.Выйти в меню")
	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("ERROR ", err)
		return
	}

	if choice == 2 {
		return
	} else {
		showCode.ShowCode()
	}
}
func joke() {
	fmt.Println()
	chess.Label()
	fmt.Printf("\t\tШутка про шахматы\n\n")
	fmt.Println("Мой компьютер постоянно обыгрывает меня в шахматы.")
	fmt.Println("-Зато я всегда побеждаю его")
	fmt.Println("в боксёрском поединке!")
	fmt.Println("\n1.Выйти в меню")
	var temp string
	_, _ = fmt.Scan(&temp)
}
var stringsExit = []string{
	"Вы точно хотите выйти?",
	"Вы абсолютно уверены, что хотите выйти?",
	"Мне кажется вы сомневаетесь. Подумайте хорошенько.",
	"Я понял, что Вам нужна помощь и решил всё за Вас. Не благодарите.",
}
func exitChess(i int) bool {
	cmd.Clear()
	fmt.Printf("\n\n\n\n\n\n\n\n")
	fmt.Println(stringsExit[i])
	if i != 3 {
		fmt.Printf("1.Да\t\t\t2.Нет\n")
	} else {
		fmt.Printf("    \t\t\t2.Нет\n")
	}
	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("ERROR ", err)
		return false
	}

	if choice == 2 {
		return true
	} else if choice == 1 {
		if i == 3 {
			fmt.Println("\nЧёрт, меня раскрыли. Ладно. Ты заслужил выход.")
			fmt.Printf("Красный крестик сверху справа.\n\n\t\tНаслаждайся.\n\n")
			for {
				_, _ = fmt.Scan()
			}
		}
		i++
		exitChess(i)
	} else {
		fmt.Println("Неверное число. Повторите ещё раз")
		exitChess(i)
	}
	return true
}