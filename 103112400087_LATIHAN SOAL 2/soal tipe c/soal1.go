package main

import (
	"fmt"
	"math"
)

func main() {
	var num_103112400087 int
	fmt.Print("Masukkan bilangan bulat positif (>10): ")
	fmt.Scan(&num_103112400087)

	temp := num_103112400087
	digitCount := 0
	for temp > 0 {
		digitCount++
		temp /= 10
	}

	var divider int
	if digitCount%2 == 0 {
		divider = int(math.Pow(10, float64(digitCount/2)))
	} else {
		divider = int(math.Pow(10, float64((digitCount / 2))))
	}

	bil1 := num_103112400087 / divider
	bil2 := num_103112400087 % divider

	fmt.Println("Bilangan 1:", bil1)
	fmt.Println("Bilangan 2:", bil2)
	fmt.Println("Hasil penjumlahan:", bil1+bil2)
}
