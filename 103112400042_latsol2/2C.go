// M.HANIF AL FAIZ
// 103112400042
package main

import "fmt"

func main() {
	fmt.Println("Program menjumlahkan bilangan kelipatan 4")
	fmt.Println("Masukkan bilangan (akhiri dengan bilangan negatif):")
	total := startRecursion()
	fmt.Printf("Total jumlah bilangan kelipatan 4: %d\n", total)
}
func startRecursion() int {
	return recursiveSum(0)
}
func recursiveSum(accumulator int) int {
	var num int
	fmt.Scan(&num)
	if num < 0 {
		return accumulator
	}
	if num > 0 && num%4 == 0 {
		return recursiveSum(accumulator + num)
	}
	return recursiveSum(accumulator)
}
