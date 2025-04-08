package main

import "fmt"

func main() {
	var x_103112400061 int
	fmt.Print("Masukkan bilangan bulat positif (>10): ")
	fmt.Scan(&x_103112400061)
	splitter(x_103112400061)
}

func countDigits(x int) int {
	var digits int = 0
	
	if x <= 0 {
		return 1
	}

	for x > 0 {
		digits++
		x /= 10
	}
	return digits
}

func splitter(x int) {
	if x <= 10 {
		fmt.Println("Bilangan harus >10 !")
	} else {
		var digits int = countDigits(x)
		var parts int = digits / 2
		var divisor int = 1

		for i := 1; i <= parts; i++ {
			divisor *= 10
		}

		bil1 := x / divisor
		bil2 := x % divisor
		jumlah := bil1 + bil2

		fmt.Println("Bilangan 1:", bil1)
		fmt.Println("Bilangan 2:", bil2)
		fmt.Println("Hasil penjumlahan:", jumlah)
	}
}