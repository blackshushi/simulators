package simulator

type IBaseSimulator interface {
	NewSimulator() *IBaseSimulator
	Simulate() *SimulateResult
	CalculateRTP(results []*SimulateResult, betOption string) float64
	CalculatePlayerRTP(results []*SimulateResult) float64
}

type SimulateResult struct {
	Data map[string]int
}

type BaseSimulator struct {
	Rtps map[string]float64
}
