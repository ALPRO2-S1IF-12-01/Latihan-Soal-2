package main

import (
	"fmt"
)

func main() {
	var jumlahPeserta int

	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&jumlahPeserta)

	jumlahHadiahA := 0
	jumlahHadiahB := 0
	jumlahHadiahC := 0

	for i := 1; i <= jumlahPeserta; i++ {
		var nomorKartu int

		fmt.Printf("Masukkan nomor kartu peserta ke-%d: ", i)
		fmt.Scan(&nomorKartu)

		jenisHadiah := tentukanHadiah(nomorKartu)

		fmt.Printf("Hadiah %s\n", jenisHadiah)

		if jenisHadiah == "A" {
			jumlahHadiahA++
		} else if jenisHadiah == "B" {
			jumlahHadiahB++
		} else {
			jumlahHadiahC++
		}
	}

	fmt.Printf("Jumlah yang memperoleh Hadiah A: %d\n", jumlahHadiahA)
	fmt.Printf("Jumlah yang memperoleh Hadiah B: %d\n", jumlahHadiahB)
	fmt.Printf("Jumlah yang memperoleh Hadiah C: %d\n", jumlahHadiahC)
}

func tentukanHadiah(nomorKartu int) string {

	if semuaDigitSama(nomorKartu) {
		return "A"
	}

	if semuaDigitBerbeda(nomorKartu) {
		return "B"
	}

	return "C"
}

func semuaDigitSama(n int) bool {
	if n < 10 {
		return true
	}

	digitPertama := n % 10
	for n > 0 {
		digit := n % 10
		if digit != digitPertama {
			return false
		}
		n /= 10
	}

	return true
}

func semuaDigitBerbeda(n int) bool {

	digitSudahAda := [10]bool{}

	for n > 0 {
		digit := n % 10

		if digitSudahAda[digit] {
			return false
		}

		digitSudahAda[digit] = true
		n /= 10
	}

	return true
}
