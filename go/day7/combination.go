package main

import "strconv"

type Combination struct {
	A        int
	B        int
	Operator Operator
}

func (c Combination) Calculate() int {
	if c.Operator == OperatorConcat {
		num, err := strconv.Atoi(strconv.Itoa(c.A) + strconv.Itoa(c.B))
		if err != nil {
			panic("error parsing number")
		}
		return num
	}
	if c.Operator == OperatorAdd {
		return c.A + c.B
	}
	return c.A * c.B
}

func GetCombinations(a, b int) []Combination {
	return []Combination{
		{A: a, B: b, Operator: OperatorAdd},
		{A: a, B: b, Operator: OperatorMult},
		{A: a, B: b, Operator: OperatorConcat},
	}
}
