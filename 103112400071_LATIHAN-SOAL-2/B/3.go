package main

import "fmt"

func main() {
	fmt.Println("Masukkan bilangan (negatif untuk berhenti):")
	fmt.Println("Jumlah bilangan kelipatan 4:", kelipatanEmpat(0))
}

func kelipatanEmpat(jumlah int) int {
	var x int
	fmt.Scan(&x)

	if x < 0 {
		return jumlah
	}

	if x > 0 && x%4 == 0 {
		return kelipatanEmpat(jumlah + x)
	}

	return kelipatanEmpat(jumlah)
}
