package main

import "fmt"

func jumlahKelipatan4(jumlah int) int {
	var bilangan int
	fmt.Scan(&bilangan)

	if bilangan < 0 {
		return jumlah
	}

	if bilangan > 0 && bilangan%4 == 0 {
		jumlah += bilangan
	}

	return jumlahKelipatan4(jumlah)
}

func main() {
	fmt.Println("Masukkan bilangan (negatif untuk berhenti):")
	hasil := jumlahKelipatan4(0)
	fmt.Println("Jumlah bilangan kelipatan 4:", hasil)
}
