package main

const (
	GreetingMessage     = "Добро пожаловать в калькулятор римских и арабских цифр."
	MenuMessage         = "Пожалуйста, напишите 2 числа от 1 до 10 или от I до X включительно и между ними укажите одно из операторов - (+, -, *, /)"
	AvailableOperations = "+-*/"
)

var roman = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

var NumbersParse = []int{1, 4, 5, 9, 10, 40, 50, 90, 100}
var RomansParse = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}
