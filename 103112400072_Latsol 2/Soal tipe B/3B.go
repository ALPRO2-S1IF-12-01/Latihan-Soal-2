package main

import "fmt"

func sumMultiplesOf4Iterative() int {
	sum := 0
	for {
		var num int
		fmt.Print("Masukkan bilangan (negatif untuk berhenti): ")
		fmt.Scan(&num)

		if num < 0 {
			break
		}

		if num > 0 && num%4 == 0 {
			sum += num
		}
	}
	return sum
}

func sumMultiplesOf4Recursive() int {
	return sumHelper(0)
}

func sumHelper(currentSum int) int {
	var num int
	fmt.Scan(&num)

	if num < 0 {
		return currentSum
	}

	if num > 0 && num%4 == 0 {
		currentSum += num
	}

	return sumHelper(currentSum)
}

func main() {

	sum := sumMultiplesOf4Recursive()
	fmt.Printf("Jumlah bilangan kelipatan 4: %d\n", sum)
}
