//Ahmad Ruba'i
//103112400074
package main

import "fmt"

func main() {
	var bilangan int
	fmt.Print("Masukkan bilangan bulat positif (>10): ")
	fmt.Scanln(&bilangan)

	n := bilangan
	jumlahDigit := 0
	for n > 0 {
		jumlahDigit++
		n /= 10
	}
	digitAwal := jumlahDigit / 2
	if jumlahDigit%2 != 0 {
		digitAwal++ 
	}
	pangkat := 1
	for i := 0; i < jumlahDigit-digitAwal; i++ {
		pangkat *= 10
	}
	bilangan1 := bilangan / pangkat
	bilangan2 := bilangan % pangkat
	hasil := bilangan1 + bilangan2

	fmt.Println("Bilangan 1:", bilangan1)
	fmt.Println("Bilangan 2:", bilangan2)
	fmt.Println("Hasil penjumlahan:", hasil)
}
