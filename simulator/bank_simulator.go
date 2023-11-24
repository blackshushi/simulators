package simulator

import (
	"math/rand"
	"simulators/calculator"
	"time"
)

type BankSimulator struct {
	IBaseSimulator
	BaseSimulator
	calculator calculator.BankCalculator
}

var ProbabiilityMap map[string]float64 = map[string]float64{
	"leaf":       0.6,
	"gold":       0.2,
	"gold_ingot": 0.1,
	"diamond":    0.1,
}

func NewSimulator() *BankSimulator {
	rand.Seed(time.Now().UnixNano())

	return &BankSimulator{
		calculator: calculator.BankCalculator{},
		BaseSimulator: BaseSimulator{
			Rtps: map[string]float64{
				"leaf":       0,
				"gold":       0,
				"gold_ingot": 0,
				"diamond":    0,
			},
		},
	}
}

func (bs *BankSimulator) Simulate() *SimulateResult {
	result := map[string]int{
		"leaf":       0,
		"gold":       0,
		"gold_ingot": 0,
		"diamond":    0,
	}
	// leaf, gold, goldingot, diamond
	// 0, 3, 5, 10

	for i := 0; i < 14; i++ {
		tile := bs.openTile()
		result[tile] += 1

		if result[tile] == 3 {
			return &SimulateResult{
				Data: result,
			}
		}
	}

	return &SimulateResult{
		Data: result,
	}
}

func (bs *BankSimulator) CalculatePlayerRTP(results []*SimulateResult) float64 {
	totalPayout := 0

	for _, result := range results {
		for betOption := range bs.Rtps {
			totalPayout += bs.calculator.CalculatePayout(result.Data, betOption)
		}
	}

	return float64(totalPayout) / float64(len(results)) * 100
}

func (bs *BankSimulator) CalculateRTP(results []*SimulateResult, betOption string) float64 {
	totalPayout := 0

	for _, result := range results {
		totalPayout += bs.calculator.CalculatePayout(result.Data, betOption)
	}

	return (float64(totalPayout) / float64(len(results))) * 100
}

func (bs *BankSimulator) openTile() string {
	r := rand.Float64()

	if r < ProbabiilityMap["leaf"] {
		return "leaf"
	} else if r < ProbabiilityMap["leaf"]+ProbabiilityMap["gold"] {
		return "gold"
	} else if r < ProbabiilityMap["leaf"]+ProbabiilityMap["gold"]+ProbabiilityMap["gold_ingot"] {
		return "gold_ingot"
	} else {
		return "diamond"
	}
}
