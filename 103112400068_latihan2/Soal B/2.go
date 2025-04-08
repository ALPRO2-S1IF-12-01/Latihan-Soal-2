package main

import (
	"fmt"
)

func main() {
	var j, o, r int
	var status int

	fmt.Print("Masukkan jumlah rombongan:")
	fmt.Scan(&j)
	for i := 0; i < j; i++ {
		fmt.Print("Masukkan jumlah menu, jumlah orang, dan status sisa makanan: ")
		fmt.Scan(&o, &r, &status)
		hasil := 10000
		if o > 3 {
			hasil += 2500 * o
		}
		if 0 >= 50 {
			hasil = 100000
		}
		if status == 1 {
			hasil += 5000
		}
		fmt.Printf("Total biaya untuk rombongan %d: RP %d\n, ", i+1, hasil)
	}
}
