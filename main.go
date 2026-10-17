package main

import (
	"errors"
	"fmt"
	"go-lang-learn/weeks"
	"os"
)

func main() {
	var err error

	args := os.Args[1:]

	if len(args) < 1 {
		printHelp()

		return
	}

	switch args[0] {
	case "help":
		printHelp()

	case "guess-number":
		weeks.GuessNumber()

	case "convert":
		if len(args) > 1 && args[1] == "list" {
			err = printUnits()
		} else {
			err = applyConverter(args)
		}
	default:
		err = errors.New(fmt.Sprintf("Указана неверная команда: %v", args))
	}

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
}

func applyConverter(args []string) error {
	if len(args) > 1 {
		args = args[1:]
	}

	value, fromSymbol, toSymbol, err := weeks.ParseArgs(args)

	if err != nil {
		return err
	}

	convertValue, err := weeks.ConvertUnit(value, fromSymbol, toSymbol)

	if err != nil {
		return err
	}

	fmt.Printf("%.2f %s = %.2f %s\n", value, fromSymbol, convertValue, toSymbol)

	return nil
}

func printUnits() error {
	lines, err := weeks.DictionaryInfo()

	if err != nil {
		return err
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	return nil
}

func printHelp() {
	fmt.Println(
		"Привет, я умею делать следующие штуки:\n",
		" - guess-number - Сыграть в игру \"Угадай число\"\n",
		" - convert list - Посмотреть список доступных един для конвертации\n",
		" - convert <значение> <из какой> <в какую> - Сконвертировать значение, например, convert 100 m km\n",
		" - help - Посмотреть доступные команды",
	)
}
