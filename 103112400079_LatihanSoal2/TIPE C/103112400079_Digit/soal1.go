package main

import "fmt"

func main() {
	var n_103112400079 int
	fmt.Print("Masukkan bilangan bulat positif (>10): ")
	fmt.Scan(&n_103112400079)

	panjang := 0
	temp := n_103112400079
	for temp > 0 {
		temp /= 10
		panjang++
	}

	tengah := panjang / 2
	pembagi := 1
	for i := 0; i < panjang-(tengah+(panjang%2)); i++ {
		pembagi *= 10
	}

	n1 := n_103112400079 / pembagi
	n2 := n_103112400079 % pembagi

	fmt.Println("Bilangan 1:", n1)
	fmt.Println("Bilangan 2:", n2)
	fmt.Println("Hasil penjumlahan:", n1+n2)
}
