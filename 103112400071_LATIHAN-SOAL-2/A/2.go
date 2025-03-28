package main

import "fmt"

// Fungsi untuk memeriksa apakah sebuah bilangan adalah perfect number
func isPerfectNumber(n int) bool {
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum == n
}

func main() {
	var a, b int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)

	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	// Validasi input
	if a <= 0 || b <= 0 || a > b {
		fmt.Println("Error: a dan b harus bilangan bulat positif dan a <= b")
		return
	}

	// Cetak perfect number dalam rentang a hingga b
	fmt.Printf("Perfect numbers antara %d dan %d: ", a, b)
	found := false
	for i := a; i <= b; i++ {
		if isPerfectNumber(i) {
			fmt.Printf("%d ", i)
			found = true
		}
	}
	if !found {
		fmt.Print("Tidak ada")
	}
	fmt.Println()
}
