//Anastasia Adinda N.I
package main

import (
	"fmt"
)

func main() {
	var x, y int
	var jumlahPertemuan_103112400085 int

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)

	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	for hari := 1; hari <= 365; hari++ {
		if hari%x == 0 && hari%y != 0 {
			jumlahPertemuan_103112400085++
		}
	}

	fmt.Printf("Jumlah hari pertemuan rahasia dalam setahun: %d\n", jumlahPertemuan_103112400085)
}
