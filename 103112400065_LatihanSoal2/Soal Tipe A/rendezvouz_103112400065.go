/*
Nama: Dimas Ramadhani
NIM: 103112400065
*/

package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai Y: ")
	fmt.Scan(&y)
	fmt.Print("Jumlah pertemuan dalam setahun: ", rendezvous(x, y))
}

func rendezvous(x, y int) int {
	var jumlah int
	for i := 1; i <= 365; i++ {
		if i%x == 0 && i%y != 0 {
			jumlah += 1
		}
	}
	return jumlah
}
