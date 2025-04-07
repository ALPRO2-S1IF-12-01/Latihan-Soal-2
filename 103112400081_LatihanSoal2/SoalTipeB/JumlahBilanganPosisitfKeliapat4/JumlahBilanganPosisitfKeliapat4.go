// RYAN AKEYLA NOVIANTO WIDODO
// 103112400081

package main

import "fmt"

func jumlahKelipatanEmpatIteratif(angka []int) int {
	jumlah := 0
	for _, nilai := range angka {
		if nilai%4 == 0 && nilai > 0 {
			jumlah += nilai
		}
	}
	return jumlah
}

func main() {
	var input int
	angka := []int{}

	fmt.Println("Masukkan bilangan (negatif untuk berhenti): ")

	for {
		fmt.Scan(&input)
		if input < 0 {
			break
		}
		angka = append(angka, input)
	}

	hasil := jumlahKelipatanEmpatIteratif(angka)
	fmt.Printf("Jumlah bilangan Kelipatan 4: %d\n", hasil)
}
