// MUHAMMAD GAMEL AL GHIFARI
// 103112400028
package main

import "fmt"

const (
	basePrice       = 10000
	additionalPrice = 2500
	bulkPrice       = 100000
	baseThreshold   = 3
	bulkThreshold   = 50
)

func main() {
	var groupCount int
	fmt.Print("Masukkan jumlah rombongan: ")
	fmt.Scan(&groupCount)

	groupCosts := calculateAllGroupCosts(groupCount)
	printResults(groupCosts)
}

func calculateAllGroupCosts(groupCount int) []int {
	costs := make([]int, groupCount)
	for i := 0; i < groupCount; i++ {
		fmt.Printf("\nData rombongan ke-%d:\n", i+1)

		menuCount, peopleCount, hasLeftover := getGroupInput()
		costs[i] = calculateGroupCost(menuCount, peopleCount, hasLeftover)
	}
	return costs
}

func getGroupInput() (int, int, bool) {
	var menuCount, peopleCount int
	var hasLeftover bool

	fmt.Print("Jumlah menu: ")
	fmt.Scan(&menuCount)
	fmt.Print("Jumlah orang: ")
	fmt.Scan(&peopleCount)
	fmt.Print("Ada sisa? (true/false): ")
	fmt.Scan(&hasLeftover)

	return menuCount, peopleCount, hasLeftover
}

func calculateGroupCost(menuCount, peopleCount int, hasLeftover bool) int {
	var totalCost int

	switch {
	case menuCount > bulkThreshold:
		totalCost = bulkPrice
	case menuCount > baseThreshold:
		totalCost = basePrice + (menuCount-baseThreshold)*additionalPrice
	default:
		totalCost = basePrice
	}

	if hasLeftover {
		totalCost *= peopleCount
	}

	return totalCost
}

func printResults(costs []int) {
	fmt.Println("\nHasil Perhitungan Biaya:")
	for i, cost := range costs {
		fmt.Printf("Rombongan ke-%d: Rp %d\n", i+1, cost)
	}
}
