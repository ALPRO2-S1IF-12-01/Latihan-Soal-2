// M.HANIF AL FAIZ
// 103112400042
package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)
	jumlahPertemuan := hitungPertemuan(x, y)
	fmt.Printf("Jumlah pertemuan rahasia dalam setahun: %d\n", jumlahPertemuan)
}
func hitungPertemuan(x, y int) int {
	count := 0
	for hari := 1; hari <= 365; hari++ {
		if hari%x == 0 && hari%y != 0 {
			count++
		}
	}
	return count
}
