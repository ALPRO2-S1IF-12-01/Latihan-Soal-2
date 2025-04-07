package main

import "fmt"

func main() {
	var n_103112400077 int
	fmt.Print("Masukkan bilangan bulat positif (>10): ")
	fmt.Scan(&n_103112400077)

	temp := n_103112400077
	pjg := 0
	for temp > 0 {
		temp /= 10
		pjg++
	}

	var pisah int = 1
	for i := 0; i < pjg/2; i++ {
		pisah *= 10
	}

	if pjg%2 != 0 {
		pisah *= 10
	}

	kiri := n_103112400077 / pisah
	kanan := n_103112400077 % pisah

	fmt.Println("Bilangan 1:", kiri)
	fmt.Println("Bilangan 2:", kanan)
	fmt.Println("Hasil penjumlahan:", kiri+kanan)
}
