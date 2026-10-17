package weeks

import (
	"errors"
	"fmt"
	"strconv"
)

type converter func(v float64) float64

type unit struct {
	symbol   string
	toBase   converter
	fromBase converter
	base     bool
}

type family struct {
	name           string
	hasCoefficient bool
	units          []unit
}

var temperatures = []unit{
	{
		symbol:   "C",
		toBase:   func(v float64) float64 { return v },
		fromBase: func(v float64) float64 { return v },
		base:     true,
	},
	{
		symbol:   "K",
		toBase:   func(v float64) float64 { return v - 273.15 },
		fromBase: func(v float64) float64 { return v + 273.15 },
	},
	{
		symbol:   "F",
		toBase:   func(v float64) float64 { return (v - 32) * 5 / 9 },
		fromBase: func(v float64) float64 { return (v * 9 / 5) + 32 },
	},
}

var distances = []unit{
	{
		symbol:   "km",
		toBase:   func(v float64) float64 { return v },
		fromBase: func(v float64) float64 { return v },
		base:     true,
	},
	{
		symbol:   "m",
		toBase:   func(v float64) float64 { return v / 1000 },
		fromBase: func(v float64) float64 { return v * 1000 },
	},
	{
		symbol:   "mi",
		toBase:   func(v float64) float64 { return v / 0.621371 },
		fromBase: func(v float64) float64 { return v * 0.621371 },
	},
}

var dictionaries = []family{
	{
		name:           "Температура",
		hasCoefficient: false,
		units:          temperatures,
	},
	{
		name:           "Расстояние",
		hasCoefficient: true,
		units:          distances,
	},
}

func DictionaryInfo() ([]string, error) {
	var result = new([]string)

	for _, d := range dictionaries {
		baseUnit, err := findBase(&d.units)

		if err != nil {
			return *result, err
		}

		*result = append(*result, fmt.Sprintf("%s (база: %s):", d.name, baseUnit.symbol))
		for _, item := range d.units {
			addUnitInfo(result, item, baseUnit, d.hasCoefficient)
		}
	}

	return *result, nil
}

func addUnitInfo(result *[]string, item unit, baseUnit unit, hasCoefficient bool) {
	if item.base {
		*result = append(*result, fmt.Sprintf("%s\t- базовая", item.symbol))
	} else {
		var coeffInfo string

		if hasCoefficient {
			coeffInfo = fmt.Sprintf("(коэффициент %f)", item.toBase(1))
		} else {
			coeffInfo = "(нет коэффициента: сдвиг)"
		}

		*result = append(*result, fmt.Sprintf(
			"%s\t- 1 %s = %f %s\t\t%s",
			item.symbol,
			item.symbol,
			item.toBase(1),
			baseUnit.symbol,
			coeffInfo,
		))
	}
}

func ParseArgs(args []string) (float64, string, string, error) {
	if len(args) != 3 {
		return 0, "", "", errors.New("укажите аргументы вызова в следующем формате: <Значение> <Исходная ед. измерения> <Конечная ед. измерения>")
	}

	value, err := strconv.ParseFloat(args[0], 64)

	if err != nil {
		return 0, "", "", errors.New("первый аргумент должен быть числом, дробное число должно быть указано через точку. Вы указали " + args[0])
	}

	return value, args[1], args[2], nil
}

func ConvertUnit(value float64, fromSymbol, toSymbol string) (float64, error) {
	from, to, err := getUnits(fromSymbol, toSymbol)

	if err != nil {
		return 0.0, err
	}

	value = from.toBase(value)

	return to.fromBase(value), nil
}

func getUnits(from, to string) (unit, unit, error) {
	findFrom, findTo := false, false
	for _, dictionary := range dictionaries {
		mapUnits := map[string]unit{}

		for _, item := range dictionary.units {
			mapUnits[item.symbol] = item
		}

		_, fromOk := mapUnits[from]
		_, toOk := mapUnits[to]

		if fromOk && toOk {
			return mapUnits[from], mapUnits[to], nil
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

	return unit{}, unit{}, errors.New(errText)
}

func findBase(units *[]unit) (unit, error) {
	for _, item := range *units {
		if item.base {
			return item, nil
		}
	}

	return unit{}, errors.New("не удалось найти базовую единицу")
}
