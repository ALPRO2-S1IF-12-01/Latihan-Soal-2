// Raihan Adi Arba_103112400071

package main

import "fmt"

// hitung jumlah digit
func countDigits(n int) int {
	if n == 0 {
		return 0
	}
	return 1 + countDigits(n/10)
}

// hitung pangkat 10
func power10(exp int) int {
	if exp == 0 {
		return 1
	}
	return 10 * power10(exp-1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif (>10): ")
	fmt.Scan(&n)

	if n <= 10 {
		fmt.Println("Error: Bilangan harus lebih besar dari 10")
		return
	}

	digit := countDigits(n)

	// Menentukan titik bagi
	split := digit / 2
	if digit%2 != 0 {
		split++
	}

	divisor := power10(digit - split)

	bil1 := n / divisor
	bil2 := n % divisor

	fmt.Println("Bilangan 1:", bil1)
	fmt.Println("Bilangan 2:", bil2)
	fmt.Println("Hasil penjumlahan:", bil1+bil2)
}
