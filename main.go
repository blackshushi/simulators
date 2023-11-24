package main

import (
	"fmt"
	"os"
	"simulators/simulator"
	"strconv"
)

func main() {
	game := os.Args[1]
	loopTimes, err := strconv.Atoi(os.Args[2])

	if err != nil {
		fmt.Println(err)
		return
	}

	if game == "bank" {
		bs := simulator.NewSimulator()
		results := []*simulator.SimulateResult{}

		for i := 0; i < loopTimes; i++ {
			results = append(results, bs.Simulate())
		}

		fmt.Println("Gold Rtp: ", bs.CalculateRTP(results, "gold"), "%")
		fmt.Println("Gold Ingot Rtp: ", bs.CalculateRTP(results, "gold_ingot"), "%")
		fmt.Println("Diamond Rtp: ", bs.CalculateRTP(results, "diamond"), "%")
		fmt.Println("Player Rtp: ", bs.CalculatePlayerRTP(results), "%")

		return
	}

	fmt.Println("Game not found")
}
