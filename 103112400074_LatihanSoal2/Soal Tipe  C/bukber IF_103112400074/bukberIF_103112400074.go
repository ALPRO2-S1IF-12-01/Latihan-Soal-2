//Ahmad Ruba'i
//103112400074
package main

import "fmt"

func main() {
	var jumlahPeserta int
	fmt.Print("Masukkan jumlah peserta:")
	fmt.Scan(&jumlahPeserta)

	hadiahA := 0
	hadiahB := 0
	hadiahC := 0

	for i := 1; i <= jumlahPeserta; i++ {
		var nomor int
		fmt.Print("Masukkan nomor kartu peserta ke-", i, ":")
		fmt.Scan(&nomor)

		hasil := cekHadiah(nomor)
		fmt.Println(hasil)

		if hasil == "Hadiah A" {
			hadiahA++
		} else if hasil == "Hadiah B" {
			hadiahB++
		} else {
			hadiahC++
		}
	}
	fmt.Println()
	fmt.Println("Jumlah yang memperoleh Hadiah A:", hadiahA)
	fmt.Println("Jumlah yang memperoleh Hadiah B:", hadiahB)
	fmt.Println("Jumlah yang memperoleh Hadiah C:", hadiahC)
}
func cekHadiah(n int) string {
	digitSama := true
	sisaDigit := n
	digitAwal := sisaDigit % 10
	sisaDigit /= 10

	for sisaDigit > 0 {
		if sisaDigit%10 != digitAwal {
			digitSama = false
		}
		sisaDigit /= 10
	}
	if digitSama {
		return "Hadiah A"
	}
	if digitBerbeda(n) {
		return "Hadiah B"
	}
	return "Hadiah C"
}
func digitBerbeda(n int) bool {
	for i := 0; i <= 9; i++ {
		if hitungDigit(n, i) > 1 {
			return false
		}
	}
	return true
}
func hitungDigit(n, target int) int {
	jumlah := 0
	for n > 0 {
		if n%10 == target {
			jumlah++
		}
		n /= 10
	}
	return jumlah
}
