package main

import "fmt"

func hitungPertemuan(x, y int) int {
	count := 0
	for i := 1; i <= 365; i++ {
		if i%x == 0 && i%y != 0 { 
			count++
		}
	}
	return count
}

func main() {
	fmt.Println("Nama: Raja Muhammad Lufhti\nNim : 103112400027")
	var x, y int

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	jumlahPertemuan := hitungPertemuan(x, y)

	fmt.Printf("Jumlah pertemuan dalam setahun: %d\n", jumlahPertemuan)
}
