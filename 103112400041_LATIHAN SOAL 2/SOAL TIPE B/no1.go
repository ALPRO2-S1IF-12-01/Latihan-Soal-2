package main

import "fmt"

func ganjil(a, b int) int {
	hitung:=0
	for i := a; i <= b; i++ {
		if i%2!=0 {
			hitung=hitung+1
		}
	}
	return hitung
}

func main() {
	var a_103112400041, b int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scanln(&a_103112400041)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scanln(&b)
	fmt.Print("Banyaknya angka ganjil: ")
	fmt.Print(ganjil(a_103112400041,b))
}