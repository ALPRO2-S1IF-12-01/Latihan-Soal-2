// M. DAVI ILYAS RENALDO/103112400062 
package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)

	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	pertemuan := 0
	for hari := 1; hari <= 365; hari++ {
		if hari%x == 0 && hari%y != 0 {
			pertemuan++
		}
	}

	fmt.Printf("Jumlah pertemuan dalam setahun: %d\n", pertemuan)
}