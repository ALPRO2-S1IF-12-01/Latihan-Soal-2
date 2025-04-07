// JESIKA METANIA RAHMA ARIFIN
// 103112400080

package main

import "fmt"

func PerfectNumber(num int) bool {
	sum := 0
	for i := 1; i <= num/2; i++ {
		if num%i == 0 {
			sum += i
		}
	}
	return sum == num
}

func main() {
	var a, b int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	fmt.Printf("Perfect numbers antara %d dan %d: ", a, b)
	found := false
	for i := a; i <= b; i++ {
		if PerfectNumber(i) {
			fmt.Printf("%d ", i)
			found = true
		}
	}
	if !found {
		fmt.Print("Tidak ada bilangan perfect dalam rentang ini.")
	}
	fmt.Println()
}
