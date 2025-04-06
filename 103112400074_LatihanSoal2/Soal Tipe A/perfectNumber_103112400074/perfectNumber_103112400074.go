//Ahmad Ruba'i
//103112400074
package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	fmt.Printf("Perfect numbers antara %d dan %d:", a, b)
	for i := a; i <= b; i++ {
		if perfectNumber(i) {
			fmt.Printf(" %d", i)
		}
	}
	fmt.Println()
}

func perfectNumber(n int) bool {
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum == n
}
