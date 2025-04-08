package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	// NIM: 103112400061
	fmt.Println("Banyaknya angka ganjil:", ganjil(a,b))
}

func ganjil(a, b int) int {
	var ganjil_103112400061 int
	for i := a; i <= b; i++ {
		if i%2 != 0 { // AUTHOR: KEISHIN NAUFA ALFARDZHI
			ganjil_103112400061++
		}
	}
	return ganjil_103112400061
}