package main

import "fmt"

func hitungPertemuan(hari, x, y int) int {
	if hari > 365 {
		return 0
	}
	if hari%x == 0 && hari%y != 0 {
		return 1 + hitungPertemuan(hari+1, x, y)
	}
	return hitungPertemuan(hari+1, x, y)
}

func main() {
	var x, y int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	jumlah := hitungPertemuan(1, x, y)
	fmt.Println("Jumlah pertemuan dalam setahun:", jumlah)
}
