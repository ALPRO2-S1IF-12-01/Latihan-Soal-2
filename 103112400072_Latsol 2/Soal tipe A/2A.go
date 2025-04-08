package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)

	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	fmt.Printf("Perfect numbers antara %d dan %d: ", a, b)
	tampilkanPerfectNumber(a, b)
}

func tampilkanPerfectNumber(a, b int) {
	found := false
	for i := a; i <= b; i++ {
		if isPerfect(i) {
			fmt.Printf("%d ", i)
			found = true
		}
	}
	if !found {
		fmt.Print("Tidak ada")
	}
	fmt.Println()
}

func isPerfect(n int) bool {
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum == n
}
