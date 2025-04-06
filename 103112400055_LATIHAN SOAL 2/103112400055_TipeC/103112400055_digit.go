package main

import "fmt"

func main() {
	var number_103112400055 int
	fmt.Print("Masukkan bilangan bulat positif (>10): ")
	fmt.Scan(&number_103112400055)

	if number_103112400055 <= 10 {
		fmt.Println("Silakan masukkan bilangan lebih besar dari 10.")
		return
	}

	digits := []int{}
	temp := number_103112400055
	for temp > 0 {
		digits = append([]int{temp % 10}, digits...)
		temp /= 10
	}

	length := len(digits)
	var bilangan1, bilangan2 int

	if length%2 == 0 {
		mid := length / 2
		for i := 0; i < mid; i++ {
			bilangan1 = bilangan1*10 + digits[i]
		}
		for i := mid; i < length; i++ {
			bilangan2 = bilangan2*10 + digits[i]
		}
	} else {
		mid := length / 2
		for i := 0; i <= mid; i++ {
			bilangan1 = bilangan1*10 + digits[i]
		}
		for i := mid + 1; i < length; i++ {
			bilangan2 = bilangan2*10 + digits[i]
		}
	}

	hasilPenjumlahan := bilangan1 + bilangan2

	fmt.Printf("Bilangan 1: %d\n", bilangan1)
	fmt.Printf("Bilangan 2: %d\n", bilangan2)
	fmt.Printf("Hasil penjumlahan: %d\n", hasilPenjumlahan)
}
