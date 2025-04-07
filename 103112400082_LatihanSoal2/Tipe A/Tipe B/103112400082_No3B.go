// Rizkina Azizah_103112400082
package main

import "fmt"

func jumlahKelipatan4() int {
	var n int
	fmt.Scan(&n)

	if n < 0 {
		return 0
	}

	if n%4 == 0 {
		return n + jumlahKelipatan4() 
	}

	return jumlahKelipatan4()
}

func main() {
	fmt.Print("Masukkan bilangan (negatif untuk berhenti):")
	total := jumlahKelipatan4() 
	fmt.Printf("Jumlah bilangan kelipatan 4: %d\n", total)
}