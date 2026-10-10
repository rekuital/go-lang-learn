package weeks

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func GuessNumber() {
	var attempt int

	goal := rand.Intn(100) + 1
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Я загадал число, попробуй угадать его >:D")
	fmt.Println("Введите число: ")

	for scanner.Scan() {
		userNumber, inputErr := strconv.Atoi(scanner.Text())

		if inputErr != nil {
			fmt.Printf("Ошибка %q не является целым числом\n", scanner.Text())
			continue
		}

		attempt++

		if userNumber == goal {
			fmt.Printf("Поздравляю вы угадали. Загадано было число - %d. Количество попыток: %d", goal, attempt)
			break
		} else if userNumber < goal {
			fmt.Println("Загаданное число больше вашего, попробуй ещё раз:")
		} else {
			fmt.Println("Загаданное число меньше вашего, попробуй ещё раз:")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения:", err)
		return
	}

}
