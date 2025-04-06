package main

import (
	"fmt"
)

func main() {
	var x, y int
	var NIM string = "103112400076" 
	fmt.Println("NIM :", NIM) 	
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)
	jumlahPertemuan := 0
	for hari := 1; hari <= 365; hari++ {
		if hari%x == 0 && hari%y != 0 {
			jumlahPertemuan++
		}
	}
	fmt.Printf("Jumlah pertemuan dalam setahun: %d\n", jumlahPertemuan)
}
