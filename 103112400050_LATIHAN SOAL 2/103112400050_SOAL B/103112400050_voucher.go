package main

import "fmt"

//103112400050
func hitungAngkaGanjil(a, b int) (jumlah int) {
	for i := a; i <= b; i++ {
		if i%2 != 0 {
			jumlah++
		}
	}
	return jumlah
}

func main() {
	var a, b int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	fmt.Printf("Jumlah angka ganjil: %d\n", hitungAngkaGanjil(a, b))
}
