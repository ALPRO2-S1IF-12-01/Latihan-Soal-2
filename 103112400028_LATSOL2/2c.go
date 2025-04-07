// MUHAMMAD GAMEL AL GHIFARI
// 103112400028
package main

import "fmt"

func main() {
	printProgramHeader()
	total := calculateMultiplesOfFour()
	printResult(total)
}

func printProgramHeader() {
	fmt.Println("Program menjumlahkan bilangan kelipatan 4")
	fmt.Println("Masukkan bilangan (akhiri dengan bilangan negatif):")
}

func calculateMultiplesOfFour() int {
	return sumMultiplesRecursively(0)
}

func sumMultiplesRecursively(accumulator int) int {
	number := readInputNumber()

	if isTerminationCondition(number) {
		return accumulator
	}

	if isMultipleOfFour(number) {
		return sumMultiplesRecursively(accumulator + number)
	}

	return sumMultiplesRecursively(accumulator)
}

func readInputNumber() int {
	var number int
	fmt.Scan(&number)
	return number
}

func isTerminationCondition(number int) bool {
	return number < 0
}

func isMultipleOfFour(number int) bool {
	return number > 0 && number%4 == 0
}

func printResult(total int) {
	fmt.Printf("Total jumlah bilangan kelipatan 4: %d\n", total)
}
