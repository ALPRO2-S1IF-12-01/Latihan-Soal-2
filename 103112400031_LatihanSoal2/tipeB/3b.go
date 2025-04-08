// Savila Nur Fadilla
// 103112400031

package main

import "fmt"

func kelipatanEmpat(total int) int {
	var n int
	fmt.Print("Masukkan bilangan (negatif untuk berhenti): ")
	fmt.Scan(&n)

	if n < 0 {
		return total
	}

	if n % 4 == 0 {
		total += n
	}

	return kelipatanEmpat(total)
}

func main() {
	total := kelipatanEmpat(0)
	fmt.Println("Jumlah bilangan kelipatan 4:", total)
}
