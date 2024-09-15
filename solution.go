package main

import (
	"fmt"
)

const (
	TotalHalvings         = 32
	HalvingPeriod   int64 = 210000
	TotalSupplyEver int64 = 2100000000000000
	InitialReward   int64 = 5000000000
	SatsPerBitcoin  int64 = 100000000
)

func main() {
	calculateAll()
}

func calculateAll() {
	totalSupply := int64(0)
	for i := 0; i <= TotalHalvings; i++ {
		blockReward := InitialReward / PowInt(2, i)
		supplyThisPeriod := HalvingPeriod * blockReward
		totalSupply += supplyThisPeriod

		fmt.Printf("Halving %d:\n", i)
		fmt.Printf(" Total supply: %.8f BTC (%d SATS)\n", float64(totalSupply)/float64(SatsPerBitcoin), totalSupply)
		fmt.Printf(" Block reward: %.8f BTC (%d SATS)\n", float64(blockReward)/float64(SatsPerBitcoin), blockReward)
		fmt.Printf(" Percentage mined: %.7f%%\n\n", float64(totalSupply)/float64(TotalSupplyEver)*100)
	}

	fmt.Printf("Final total supply: %.2f BTC (%d SATS)\n", float64(totalSupply)/float64(SatsPerBitcoin), totalSupply)
}

func PowInt(base, exp int) int64 {
	result := int64(1)
	for exp > 0 {
		result *= int64(base)
		exp--
	}
	return result
}