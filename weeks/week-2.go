package weeks

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type converter func(v float64, isReverse bool) float64

var temperatures = map[string]converter{
	"C": func(v float64, isReverse bool) float64 { return v },
	"K": func(v float64, isReverse bool) float64 {
		if isReverse {
			return v - 273.15
		}

		return v + 273.15
	},
	"F": func(v float64, isReverse bool) float64 {
		if isReverse {
			return (v - 32) * 5 / 9
		}

		return (v * 9 / 5) + 32
	},
}

var distances = map[string]converter{
	"km": func(v float64, isReverse bool) float64 { return v },
	"m": func(v float64, isReverse bool) float64 {
		if isReverse {
			return v / 1000
		}
		return v * 1000
	},
	"mi": func(v float64, isReverse bool) float64 {
		if isReverse {
			return v / 0.621371
		}
		return v * 0.621371
	},
}

var dictionaries = [...]map[string]converter{
	temperatures,
	distances,
}

func ConvertUnit() {
	arguments := os.Args

	if len(arguments) != 4 {
		fmt.Println("Укажите аргументы вызова в следующем формате: <Значение> <Исходная ед. измерения> <Конечная ед. измерения>")
		os.Exit(2)
	}

	value, err := strconv.ParseFloat(arguments[1], 64)

	if err != nil {
		fmt.Printf("Первый аргумент должен быть числом, дробное число должно быть указано через точку. Вы указали %s\n", arguments[1])
		os.Exit(2)
	}

	fromUnit, toUnit := arguments[2], arguments[3]

	dictionary, err := canConvert(fromUnit, toUnit)

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	toValue := convertTo(value, fromUnit, toUnit, dictionary)

	fmt.Printf("%.2f %s = %.2f %s", value, fromUnit, toValue, toUnit)
}

func canConvert(from, to string) (map[string]converter, error) {
	findFrom, findTo := false, false
	for _, dictionary := range dictionaries {
		_, fromOk := dictionary[from]
		_, toOk := dictionary[to]

		if fromOk && toOk {
			return dictionary, nil
		}

		findFrom = findFrom || fromOk
		findTo = findTo || toOk
	}

	errText := "не удалось определить единицы измерения"

	switch true {
	case findFrom && !findTo:
		errText = "вы указали не существующую единицу перевода"
	case !findFrom && findTo:
		errText = "вы указали не существующую исходную единицу"
	case !findFrom && !findTo:
		errText = "указанные единицы измерения не существуют"
	case findFrom && findTo:
		errText = "указанные единицы измерения не совместимы"
	}

	return nil, errors.New(errText)
}

func convertTo(value float64, from, to string, dictionaries map[string]converter) float64 {
	value = dictionaries[from](value, true)

	return dictionaries[to](value, false)
}
