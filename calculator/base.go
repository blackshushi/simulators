package calculator

type BaseCalculator interface {
	CalculatePayout(result map[string]int, betOption string) int
}
