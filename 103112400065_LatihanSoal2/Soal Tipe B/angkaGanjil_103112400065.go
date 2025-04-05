/*
Nama: Dimas Ramadhani
NIM: 10311240065
*/
package main

import "fmt"

func main() {
	var bil1, bil2, hasil int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&bil1)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&bil2)
	voucher(bil1, bil2, &hasil)
	fmt.Print("Banyaknya angka ganjil: ", hasil)
}

func voucher(a, b int, c *int) {
	for i := a; i <= b; i++ {
		if i%2 != 0 {
			*c += 1
		}
	}
}
