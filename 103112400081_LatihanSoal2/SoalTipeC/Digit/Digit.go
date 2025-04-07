// RYAN AKEYLA NOVIANTO WIDODO
// 103112400081

package main

import (
	"fmt"
	"strconv"
)

func bagistring(input string) (string, string) {
	length := len(input)
	mid := length / 2
	if length%2 == 0 {
		return input[:mid], input[mid:]
	} else {
		return input[:mid+1], input[mid+1:]
	}
}

func main() {
	var input string
	fmt.Print("Masukkan bilangan bulat positif (>10): ")
	fmt.Scan(&input)

	num, err := strconv.Atoi(input)
	if err != nil || num <= 10 {
		fmt.Println("Input tidak valid, harus lebih dari 10")
		return
	}

	bil1, bil2 := bagistring(input)
	bil1Int, _ := strconv.Atoi(bil1)
	bil2Int, _ := strconv.Atoi(bil2)
	sum := bil1Int + bil2Int

	fmt.Println("Bilangan 1:", bil1)
	fmt.Println("Bilangan 2:", bil2)
	fmt.Println("Hasil penjumlahan", sum)
}
