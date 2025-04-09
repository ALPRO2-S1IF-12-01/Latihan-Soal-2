package main

import "fmt"

//103112400050
func main() {
	var bilangan int
	fmt.Print("Masukkan bilangan bulat positif (>10): ")
	fmt.Scanln(&bilangan)
	digit := 0
	for b := bilangan; b > 0; b /= 10 {
		digit++
	}
	pangkat := 1
	for i := 0; i < digit/2; i++ {
		pangkat *= 10
	}
	angka1 := bilangan / pangkat
	angka2 := bilangan % pangkat
	hasil := angka1 + angka2

	fmt.Println("Bilangan 1:", angka1)
	fmt.Println("Bilangan 2:", angka2)
	fmt.Println("Hasil penjumlahan:", hasil)
}
