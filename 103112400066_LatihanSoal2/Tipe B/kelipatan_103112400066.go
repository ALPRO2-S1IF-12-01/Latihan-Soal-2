// DWI OKTA SURYANINGRUM / 103112400066

package main

import "fmt"

// Fungsi rekursif buat menjumlahkan bilangan kelipatan 4
func jumlahKelipatan4(total int) int {
	var bilangan int
	fmt.Scan(&bilangan)

	// Kalau user input angka negatif, berhenti dan balikin total
	if bilangan < 0 {
		return total
	}

	// Kalau bilangan positif & kelipatan 4, tambahkan ke total
	if bilangan > 0 && bilangan%4 == 0 {
		total += bilangan
	}

	// Panggil fungsi lagi (rekursi)
	return jumlahKelipatan4(total)
}

func main() {
	fmt.Println("Masukkan bilangan (negatif untuk berhenti):")
	total := jumlahKelipatan4(0)
	fmt.Printf("Jumlah bilangan kelipatan 4: %d\n", total)
}
