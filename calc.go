package main

type Calc struct {
	A int
	B int
}

func NewCalc(A int, B int) *Calc {
	return &Calc{
		A,
		B,
	}
}

func (calc *Calc) add() int {
	return calc.A + calc.B
}

func (calc *Calc) sub() int {
	return calc.A - calc.B
}

func (calc *Calc) mul() int {
	return calc.A * calc.B
}

func (calc *Calc) div() int {
	return calc.A / calc.B
}

func (calc *Calc) parseRoman(num int) string {
	result := ""
	for i := len(NumbersParse) - 1; num > 0; i-- {
		for NumbersParse[i] <= num {
			result += RomansParse[i]
			num -= NumbersParse[i]
		}
	}
	return result
}
