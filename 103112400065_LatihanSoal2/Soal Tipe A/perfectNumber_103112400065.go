/*
Nama: Dimas Ramadhani
NIM: 103112400065
*/

package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	fmt.Printf("Perfect numbers antara %d dan %d:", a, b)
	perfectNumber(a, b)
}

func perfectNumber(a, b int) {
	var jumlah int
	if a > b {
		return
	}
	for i := 1; i <= a; i++ {
		if a%i == 0 {
			if i!=a {
				jumlah+=i
			}
		}
	}
	if jumlah==a {
		fmt.Print(" ",jumlah)
	}
	perfectNumber(a+1, b)
}
