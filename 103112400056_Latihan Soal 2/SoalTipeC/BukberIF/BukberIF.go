package main

import (
	"fmt"
	"strconv"
)

// Fungsi untuk memeriksa apakah semua digit sama
func semuaDigitSama(n int) bool {
	nStr := strconv.Itoa(n)
	digit := nStr[0]
	for i := 1; i < len(nStr); i++ {
		if nStr[i] != digit {
			return false
		}
	}
	return true
}

// Fungsi untuk memeriksa apakah semua digit berbeda
func semuaDigitBerbeda(n int) bool {
	digitMap := make(map[rune]bool)
	nStr := strconv.Itoa(n)
	for _, digit := range nStr {
		if digitMap[digit] {
			return false
		}
		digitMap[digit] = true
	}
	return true
}

func main() {
	var n_103112400056 int
	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&n_103112400056)

	hadiahs := make([]string, n_103112400056)
	hadiahA, hadiahB, hadiahC := 0, 0, 0

	for i := 0; i < n_103112400056; i++ {
		var kartu int
		fmt.Printf("Masukkan nomor kartu peserta ke-%d: ", i+1)
		fmt.Scan(&kartu)

		if semuaDigitSama(kartu) {
			hadiahs[i] = "Hadiah A"
			hadiahA++
		} else if semuaDigitBerbeda(kartu) {
			hadiahs[i] = "Hadiah B"
			hadiahB++
		} else {
			hadiahs[i] = "Hadiah C"
			hadiahC++
		}
	}

	for _, hadiah := range hadiahs {
		fmt.Println(hadiah)
	}

	fmt.Printf("Jumlah yang memperoleh Hadiah A: %d\n", hadiahA)
	fmt.Printf("Jumlah yang memperoleh Hadiah B: %d\n", hadiahB)
	fmt.Printf("Jumlah yang memperoleh Hadiah C: %d\n", hadiahC)
}