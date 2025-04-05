/*
Nama: Dimas Ramadhani
NIM: 103112400065
*/

package main

import "fmt"

func main() {
	var bil int
	fmt.Print("Masukkan bilangan bulat positif (>10): ")
	fmt.Scan(&bil)
	membagiDigit(bil)
}

func membagiDigit(a int) {
	var bil1, bil2, pembagi int
	pembagi = 1
	tempBil := a
	for a >= 10 {
		pembagi *= 10
		a = a / pembagi
	}
	if tempBil>=100 {
		pembagi/=10
	}
	bil1 = tempBil / pembagi
	bil2 = tempBil % pembagi
	fmt.Println("Bilangan 1:", bil1)
	fmt.Println("Bilangan 2:", bil2)
	fmt.Println("Hasil penjumlahan:", bil1+bil2)
}
