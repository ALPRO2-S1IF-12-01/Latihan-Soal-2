package main

import (
	"fmt"
)

func main() {
	var a, b int
	var NIM string = "103112400076" 
	fmt.Println("NIM :", NIM)
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	jumlahGanjil := 0
	for i := a; i <= b; i++ {
		if i%2 != 0 {
			jumlahGanjil++
		}
	}
	fmt.Printf("Banyaknya angka ganjil: %d\n", jumlahGanjil)
}
