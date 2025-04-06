//Ahmad Ruba'i
//103112400074
package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)
	fmt.Print("jumlah pertemuan dalam setahun: ",rendezvous(x, y))
}

func rendezvous(x, y int) int {
	jumlah := 0
	for hari := 1; hari <= 365; hari++ {
		if hari%x == 0 && hari%y != 0 {
			jumlah++
		}
	}
	return jumlah
}
