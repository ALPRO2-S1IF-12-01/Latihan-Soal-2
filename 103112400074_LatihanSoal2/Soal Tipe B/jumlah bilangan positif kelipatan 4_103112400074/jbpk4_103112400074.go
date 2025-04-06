//Ahmad Ruba'i
//103112400074
package main

import "fmt"

func main() {
	fmt.Println("Masukkan bilangan (negatif untuk berhenti):")
	hasil := Kelipatan(0)
	fmt.Println("Jumlah bilangan kelipatan 4:", hasil)
}

func Kelipatan(hasil int) int {
	var n int
	fmt.Scan(&n)

	if n < 0 {
		return hasil
	}
	if n > 0 && n%4 == 0 {
		hasil += n
	}
	return Kelipatan(hasil)
}
