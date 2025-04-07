// RYAN AKEYLA NOVIANTO WIDODO
// 103112400081

package main

import (
	"fmt"
	"math"
)

func isPerfectNumber(n int) bool {
	sum := 1
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(n)))); i++ {
		if n%i == 0 {
			sum += i
			if i*i != n {
				sum += n / i
			}
		}
	}
	return sum == n
}

func printPerfectNumbers(a, b int) {
	fmt.Printf("Perfect Numbers antara %d dan %d: ", a, b)
	for i := a; i <= b; i++ {
		if isPerfectNumber(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

func main() {
	var a, b int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	if a <= b {
		printPerfectNumbers(a, b)
	} else {
		fmt.Println("Nilai a harus kurang dari atau sama dengan b")
	}
}
