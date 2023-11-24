package calculator

type BankCalculator struct {
	BaseCalculator
}

var PayoutMap map[string]int = map[string]int{
	"leaf":       0,
	"gold":       3,
	"gold_ingot": 5,
	"diamond":    10,
}

func (bc *BankCalculator) CalculatePayout(result map[string]int, betOption string) int {
	c := result[betOption]

	if c == 3 {
		return PayoutMap[betOption]
	}

	return 0
}
