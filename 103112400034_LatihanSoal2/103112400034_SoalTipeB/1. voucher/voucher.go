package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("NAMA: MULIA AKBAR NANDA PRATAMA\nNIM: 103112400034")
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
	fmt.Println("Banyaknya angka ganjil:", jumlahGanjil)
}
