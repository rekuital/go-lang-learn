package main

import (
	"fmt"
	"go-lang-learn/weeks"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		printHelp()
	}

	switch true {
	case args[0] == "help":
		printHelp()

	case args[0] == "guess-number":
		weeks.GuessNumber()

	case args[0] == "convert" && len(args) > 1 && args[1] == "list":
		for _, item := range weeks.PrintConvertUnits() {
			fmt.Println(item)
		}

	case args[0] == "convert":
		var convertData []string

		if len(args) > 1 {
			convertData = args[1:]
		}

		applyConverter(&convertData)

	default:
		fmt.Println("Указана неверная команда: ", args)
		os.Exit(2)
	}
}

func applyConverter(args *[]string) {
	value, fromSymbol, toSymbol, err := weeks.ParseArgs(args)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	convertValue, err := weeks.ConvertUnit(value, fromSymbol, toSymbol)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	fmt.Printf("%.2f %s = %.2f %s", value, fromSymbol, convertValue, toSymbol)
}

func printHelp() {
	fmt.Println(
		"Привет, я умею делать следующие штуки:\n",
		" - guess-number - Сыграть в игру \"Угадай число\"\n",
		" - convert list - Посмотреть список доступных един для конвертации\n",
		" - convert <значение> <из какой> <в какую> - Сконвертировать значение, например, convert 100 m km\n",
		" - help - Посмотреть доступные команды",
	)
	os.Exit(0)
}
