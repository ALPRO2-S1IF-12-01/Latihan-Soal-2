// M. DAVI ILYAS RENALDO/103112400062
package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scanln(&a)

	fmt.Print("Masukkan nilai b: ")
	fmt.Scanln(&b)

	jumlahGanjil := 0
	for i := a; i <= b; i++ {
		if i%2 != 0 {
			jumlahGanjil++
		}
	}

	fmt.Println("Banyaknya angka ganjil:", jumlahGanjil)
}
