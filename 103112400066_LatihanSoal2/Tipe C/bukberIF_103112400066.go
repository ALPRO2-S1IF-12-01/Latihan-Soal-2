// DWI OKTA SURYANINGRUM / 103112400066

package main

import "fmt"

// Fungsi buat cek apakah semua digit sama atau semua digit berbeda
func cekDigit(nomor int) (bool, bool) {
	if nomor < 10 {
		// Kalau hanya satu digit, dianggap semua sama & semua beda (karena cuma satu angka)
		return true, true
	}

	digitTerakhir := nomor % 10
	nomor /= 10
	semuaSama := true
	semuaBeda := true
	digitMap := make(map[int]bool)
	digitMap[digitTerakhir] = true

	for nomor > 0 {
		digit := nomor % 10
		if digit != digitTerakhir {
			semuaSama = false
		}
		if digitMap[digit] {
			semuaBeda = false
		}
		digitMap[digit] = true
		nomor /= 10
	}

	return semuaSama, semuaBeda
}

func main() {
	var N int
	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&N)

	countA, countB, countC := 0, 0, 0

	for i := 0; i < N; i++ {
		var nomor int
		fmt.Printf("Masukkan nomor kartu peserta ke-%d: ", i+1)
		fmt.Scan(&nomor)

		semuaSama, semuaBeda := cekDigit(nomor)

		// Klasifikasikan berdasarkan kondisi digit
		if semuaSama {
			fmt.Println("Hadiah A")
			countA++
		} else if semuaBeda {
			fmt.Println("Hadiah B")
			countB++
		} else {
			fmt.Println("Hadiah C")
			countC++
		}
	}

	// Tampilkan ringkasan hadiah
	fmt.Printf("\nJumlah yang memperoleh Hadiah A: %d\n", countA)
	fmt.Printf("Jumlah yang memperoleh Hadiah B: %d\n", countB)
	fmt.Printf("Jumlah yang memperoleh Hadiah C: %d\n", countC)
}
