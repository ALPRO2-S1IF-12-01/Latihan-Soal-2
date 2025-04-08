package main

import "fmt"

func hitungKelipatan4Rekursif(jumlah_103112400079 int) int {
	var bilangan int
	fmt.Scan(&bilangan)

	if bilangan < 0 {
		return jumlah_103112400079
	}

	if bilangan > 0 && bilangan%4 == 0 {
		jumlah_103112400079 += bilangan
	}
	return hitungKelipatan4Rekursif(jumlah_103112400079)
}

func main() {
	fmt.Println("Masukkan bilangan (negatif untuk berhenti):")
	jumlahKelipatan := hitungKelipatan4Rekursif(0)
	fmt.Printf("Jumlah bilangan kelipatan 4: %d\n", jumlahKelipatan)
}
