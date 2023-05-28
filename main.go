package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(GreetingMessage)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(MenuMessage)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		values := strings.Split(text, " ")

		if len(values) != 3 {
			panic("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		}

		digitFirst := values[0]
		operation := values[1]
		digitSecond := values[2]

		var isFirstDigitRoman bool
		if v, err := strconv.Atoi(digitFirst); err == nil {
			if v > 0 && v <= 10 {
				isFirstDigitRoman = false
			} else {
				panic("Вывод ошибки, так как калькулятор должен принимать на вход числа от 1 до 10 включительно, не более.")
			}
		} else if _, ok := roman[digitFirst]; ok {
			isFirstDigitRoman = true
		} else {
			panic("Вывод ошибки, так как калькулятор должен принимать на вход от I до X включительно, не более.")
		}

		var isSecondDigitRoman bool
		if v, err := strconv.Atoi(digitSecond); err == nil {
			if v > 0 && v <= 10 {
				isSecondDigitRoman = false
			} else {
				panic("Вывод ошибки, так как калькулятор должен принимать на вход числа от 1 до 10 включительно, не более.")
			}
		} else if _, ok := roman[digitSecond]; ok {
			isSecondDigitRoman = true
		} else {
			panic("Вывод ошибки, так как калькулятор должен принимать на вход от I до X включительно, не более.")
		}

		if ok := strings.Contains(AvailableOperations, operation); ok {
			if (isFirstDigitRoman && isSecondDigitRoman) || (!isFirstDigitRoman && !isSecondDigitRoman) {
				var a, b int
				if isFirstDigitRoman {
					a, b = roman[digitFirst], roman[digitSecond]
				} else {
					a, _ = strconv.Atoi(digitFirst)
					b, _ = strconv.Atoi(digitSecond)
				}

				calc := NewCalc(a, b)

				switch operation {
				case "+":
					{
						result := calc.add()
						if isFirstDigitRoman {
							if result < 1 {
								panic("Вывод ошибки, так как в римской системе нет отрицательных чисел или нуля.")
							}
							fmt.Println(calc.parseRoman(result))
						} else {
							fmt.Println(result)
						}
						break
					}
				case "-":
					{
						result := calc.sub()
						if isFirstDigitRoman {
							if result < 1 {
								panic("Вывод ошибки, так как в римской системе нет отрицательных чисел или нуля.")
							}
							fmt.Println(calc.parseRoman(result))
						} else {
							fmt.Println(result)
						}
						break
					}
				case "*":
					{
						result := calc.mul()
						if isFirstDigitRoman {
							if result < 1 {
								panic("Вывод ошибки, так как в римской системе нет отрицательных чисел или нуля.")
							}
							fmt.Println(calc.parseRoman(result))
						} else {
							fmt.Println(result)
						}
						break
					}
				case "/":
					{
						result := calc.div()
						if isFirstDigitRoman {
							if result < 1 {
								panic("Вывод ошибки, так как в римской системе нет отрицательных чисел или нуля.")
							}
							fmt.Println(calc.parseRoman(result))
						} else {
							fmt.Println(result)
						}
						break
					}
				}
			} else {
				panic("Вывод ошибки, так как используются одновременно разные системы счисления.")
			}
		} else {
			panic("Вывод ошибки, так как строка не является математической операцией.")
		}
	}
}
