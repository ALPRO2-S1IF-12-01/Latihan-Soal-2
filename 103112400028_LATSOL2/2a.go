// MUHAMMAD GAMEL AL GHIFARI
// 103112400028
package main

import "fmt"

func countOddNumbers(lowerBound, upperBound int) int {
	oddCount := 0
	for number := lowerBound; number <= upperBound; number++ {
		if isOdd(number) {
			oddCount++
		}
	}
	return oddCount
}

func isOdd(n int) bool {
	return n%2 != 0
}

func main() {
	upperLimit := 1000
	oddNumbersCount := countOddNumbers(1, upperLimit)
	fmt.Printf("Total bilangan ganjil antara 1 dan %d: %d\n", upperLimit, oddNumbersCount)
}
