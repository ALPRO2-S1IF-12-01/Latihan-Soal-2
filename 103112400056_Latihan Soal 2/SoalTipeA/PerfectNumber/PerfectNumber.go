// Felix Pedrosa V

package main

import "fmt"

// Fungsi untuk memeriksa apakah suatu bilangan adalah perfect number atau tidak
func isPerfectNumber(n int) bool {
	sum := 0
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum == n
}

// Fungsi untuk mencetak perfect numbers dalam rentang a hingga b
func printPerfectNumbers(a, b int) {
	fmt.Printf("Perfect numbers antara %d dan %d: ", a, b)
	for i := a; i <= b; i++ {
		if isPerfectNumber(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println() // Untuk baris baru setelah output
}

func main() {
	var a, b_103112400056 int

	// Input
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b_103112400056)

	// Validasi input agar a <= b
	if a <= b_103112400056 {
		printPerfectNumbers(a, b_103112400056)
	} else {
		fmt.Println("Nilai a harus kurang dari atau sama dengan b.")
	}
}