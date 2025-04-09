package main

import "fmt"

//103112400050
func main() {
	total := jumlahKelipatan4(0)
	fmt.Printf("Jumlah bilangan kelipatan 4: %d\n", total)
}

func jumlahKelipatan4(total int) int {
	var bilangan int
	fmt.Print("Masukkan bilangan (negatif untuk berhenti): ")
	fmt.Scan(&bilangan)

	if bilangan < 0 {
		return total
	}
	if bilangan%4 == 0 {
		total += bilangan
	}

	return jumlahKelipatan4(total)
}
