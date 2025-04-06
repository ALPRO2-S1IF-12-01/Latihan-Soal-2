package main

import "fmt"


func hitungAngkaGanjil(a, b int) int {
	hitung := 0
	for i := a; i <= b; i++ {
		if i%2 != 0 {
			hitung++
		}
	}
	return hitung
}

func main() {
	var a_103112400073, b int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a_103112400073)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	
	angkaGanjil := hitungAngkaGanjil(a_103112400073, b)

	fmt.Print("Banyaknya angka ganjil : " , angkaGanjil)
}
