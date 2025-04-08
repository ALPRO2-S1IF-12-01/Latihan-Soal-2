package main

import "fmt"

func main() {
	var bilangan_103112400068 int
	fmt.Print("Masukkan bilangan bulat positif (>10): ")
	fmt.Scan(&bilangan_103112400068)

	angka := 0
	temp := bilangan_103112400068
	for temp != 0 {
		temp /= 10
		angka++
	}

	splitPos := angka / 2
	if angka%2 != 0 {
		splitPos++
	}

	pembagi := 1
	for i := 0; i < splitPos; i++ {
		pembagi *= 10
	}

	bagianKiri := bilangan_103112400068 / pembagi
	bagianKanan := bilangan_103112400068 % pembagi

	jumlah := bagianKiri + bagianKanan

	fmt.Printf("Bilangan 1: %d\n", bagianKiri)
	fmt.Printf("Bilangan 2: %d\n", bagianKanan)
	fmt.Printf("Hasil penjumlahan: %d\n", jumlah)
}
