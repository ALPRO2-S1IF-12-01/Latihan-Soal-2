// MUHAMMAD GAMEL AL GHIFARI
// 103112400028
package main

import "fmt"

func main() {
	bilangan := mintaInputBilangan()
	bagian1, bagian2 := potongBilangan(bilangan)
	hasil := hitungJumlah(bagian1, bagian2)
	tampilkanHasil(bagian1, bagian2, hasil)
}

func mintaInputBilangan() int {
	var n int
	for {
		fmt.Print("Masukkan bilangan bulat positif (>10): ")
		fmt.Scan(&n)
		if n > 10 {
			return n
		}
		fmt.Println("Error: Bilangan harus >10")
	}
}

func potongBilangan(n int) (int, int) {
	// Hitung jumlah digit
	digit := 0
	temp := n
	for temp > 0 {
		temp /= 10
		digit++
	}

	// Hitung titik potong
	titik := digit / 2
	if digit%2 != 0 {
		titik++
	}

	// Pisahkan bilangan
	pembagi := 1
	for i := 0; i < digit-titik; i++ {
		pembagi *= 10
	}

	return n / pembagi, n % pembagi
}

func hitungJumlah(b1, b2 int) int {
	return b1 + b2
}

func tampilkanHasil(b1, b2, hasil int) {
	fmt.Printf("Bilangan 1: %d\n", b1)
	fmt.Printf("Bilangan 2: %d\n", b2)
	fmt.Printf("Hasil penjumlahan: %d\n", hasil)
}
