package main

import "fmt"

func main() {
	var a, b, jumlah int

	fmt.Print("masukkan nilai a : ")
	fmt.Scan(&a)
	fmt.Print("masukkan nilai b : ")
	fmt.Scan(&b)

	for i := a; i <= b; i++ {
		if i%2 != 0 {
			jumlah++
		}
	}
	fmt.Printf("banyaknya angka ganjil : %d\n", jumlah)
	fmt.Print("Sheila Stephanie A [103112400086]")
}
