/*
Nama: Dimas Ramadhani
NIM: 103112400065
*/

package main

import "fmt"

func main() {
	var N, kode, jumA, jumB, jumC int
	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&N)
	for i := 1; i <= N; i++ {
		fmt.Printf("Masukkan nomor kartu peserta ke-%d: ", i)
		fmt.Scan(&kode)

		if bilanganSama(kode) {
			fmt.Println("Hadiah A")
			jumA++
		} else if bilanganBeda(kode) {
			fmt.Println("Hadiah B")
			jumB++
		} else {
			fmt.Println("Hadiah C")
			jumC++
		}
	}
	fmt.Println()
	fmt.Println("Jumlah yang memperoleh Hadiah A:", jumA)
	fmt.Println("Jumlah yang memperoleh Hadiah B:", jumB)
	fmt.Println("Jumlah yang memperoleh Hadiah C:", jumC)
}

func bilanganSama(a int) bool {
	var b1, b2, b3 int
	b1 = a / 100
	b2 = (a / 10) % 10
	b3 = a % 10
	if b1 == b2 && b2 == b3 && b3 == b1 {
		return true
	}
	return false
}

func bilanganBeda(a int) bool {
	var b1, b2, b3 int
	b1 = a / 100
	b2 = (a / 10) % 10
	b3 = a % 10
	if b1 != b2 && b2 != b3 && b3 != b1 {
		return true
	}
	return false
}
