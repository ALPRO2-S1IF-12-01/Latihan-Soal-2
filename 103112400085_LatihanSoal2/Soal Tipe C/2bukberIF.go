//Anastasia Adinda N.I
package main

import (
	"fmt"
)

func semuaDigitSama(n int) bool {
	digit := n % 10
	n /= 10
	for n > 0 {
		if n%10 != digit {
			return false
		}
		n /= 10
	}
	return true
}

func semuaDigitBerbeda(n int) bool {
	cek := make(map[int]bool)
	for n > 0 {
		d := n % 10
		if cek[d] {
			return false
		}
		cek[d] = true
		n /= 10
	}
	return true
}

func main() {
	var jumlahPeserta int
	var nomor_kartu_103112400085 int
	var hadiahA, hadiahB, hadiahC int

	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&jumlahPeserta)

	for i := 1; i <= jumlahPeserta; i++ {
		fmt.Printf("Masukkan nomor kartu peserta ke-%d: ", i)
		fmt.Scan(&nomor_kartu_103112400085)

		if semuaDigitSama(nomor_kartu_103112400085) {
			fmt.Println("Hadiah A")
			hadiahA++
		} else if semuaDigitBerbeda(nomor_kartu_103112400085) {
			fmt.Println("Hadiah B")
			hadiahB++
		} else {
			fmt.Println("Hadiah C")
			hadiahC++
		}
	}

	fmt.Printf("Jumlah yang memperoleh Hadiah A: %d\n", hadiahA)
	fmt.Printf("Jumlah yang memperoleh Hadiah B: %d\n", hadiahB)
	fmt.Printf("Jumlah yang memperoleh Hadiah C: %d\n", hadiahC)
}
