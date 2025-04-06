// Felix Pedrosa V

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input_103112400056 string
	fmt.Print("Masukkan bilangan bulat positif (>10): ")
	fmt.Scan(&input_103112400056)

	num, err := strconv.Atoi(input_103112400056)
	if err != nil || num <= 10 {
		fmt.Println("Input tidak valid, harus lebih besar dari 10.")
		return
	}

	// Mendapatkan panjang digit
	length := len(input_103112400056)

	var bil1, bil2 string

	if length%2 == 0 {
		mid := length / 2
		bil1 = input_103112400056[:mid]
		bil2 = input_103112400056[mid:]
	} else {
		mid := length / 2
		bil1 = input_103112400056[:mid+1]
		bil2 = input_103112400056[mid+1:]
	}

	bil1Int, _ := strconv.Atoi(bil1)
	bil2Int, _ := strconv.Atoi(bil2)
	sum := bil1Int + bil2Int

	// Menampilkan hasil
	fmt.Println("Bilangan 1:", bil1)
	fmt.Println("Bilangan 2:", bil2)
	fmt.Println("Hasil penjumlahan:", sum)
}