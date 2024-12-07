package main

type Operator string

const (
	MaxOperators            = 2
	OperatorAdd    Operator = "+"
	OperatorMult   Operator = "*"
	OperatorConcat Operator = "||"
)

type Tuple struct {
	A int
	B int
}

type Calibration struct {
	TestValue int
	Numbers   []int
	Operators []Operator
}

func (c Calibration) possibleCombinations() [][]Combination {
	var operatorCombinations = make([][]Combination, 0)
	for i := 1; i < len(c.Numbers); i++ {
		pairing := Tuple{A: c.Numbers[i-1], B: c.Numbers[i]}
		operatorCombinations = append(operatorCombinations, GetCombinations(pairing.A, pairing.B))
	}

	combinations := [][]Combination{{}}
	for _, group := range operatorCombinations {
		var tmp = make([][]Combination, 0)
		for _, combination := range combinations {
			for _, item := range group {
				newCombination := append([]Combination{}, combination...)
				newCombination = append(newCombination, item)
				tmp = append(tmp, newCombination)
			}
		}
		combinations = tmp
	}

	return combinations
}

func (c Calibration) Possible() bool {
	combinations := c.possibleCombinations()
	for _, combination := range combinations {
		sum := 0
		for i, c := range combination {
			if i != 0 {
				c.A = sum
				sum = 0
			}
			sum += c.Calculate()
		}
		if sum == c.TestValue {
			return true
		}
	}
	return false
}

func NewCalibration(testValue int, numbers []int) Calibration {
	return Calibration{
		TestValue: testValue,
		Numbers:   numbers,
		Operators: make([]Operator, 0),
	}
}
