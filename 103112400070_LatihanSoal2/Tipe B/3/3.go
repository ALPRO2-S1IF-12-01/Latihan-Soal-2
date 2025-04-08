// Achmad Zulvan Nur Hakim 103112400070
package main

import "fmt"

func kelipatan4(bil int, total int) int {
	if bil < 0 {
		return total
	}
	var n int
	fmt.Scan(&n)
	if bil > 0 && bil%4 == 0 {
		return kelipatan4(n, total+bil)
	}
	return kelipatan4(n, total)
}

func main() {
	fmt.Println("Masukkan bil (negatif stop):")
	var x int
	fmt.Scan(&x)
	jumlah := kelipatan4(x, 0)
	fmt.Printf("Jumlah kelipatan 4: %d\n", jumlah)
}