// Savila Nur Fadilla
// 103112400031

package main

import "fmt"

func main() {
	var jumlahPeserta int
	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&jumlahPeserta)

	jumlahHadiahA := 0
	jumlahHadiahB := 0
	jumlahHadiahC := 0

	for i := 1; i <= jumlahPeserta; i++ {
		var nomorKartu int
		fmt.Printf("Masukkan nomor kartu peserta ke-%v: ", i)
		fmt.Scan(&nomorKartu)

		if semuaDigitSama(nomorKartu) {
			fmt.Println("Hadiah A")
			jumlahHadiahA++
		} else if semuaDigitBerbeda(nomorKartu) {
			fmt.Println("Hadiah B")
			jumlahHadiahB++
		} else {
			fmt.Println("Hadiah C")
			jumlahHadiahC++
		}
	}

	fmt.Println()
	fmt.Println("Jumlah yang memperoleh Hadiah A:", jumlahHadiahA)
	fmt.Println("Jumlah yang memperoleh Hadiah B:", jumlahHadiahB)
	fmt.Println("Jumlah yang memperoleh Hadiah C:", jumlahHadiahC)
}

func semuaDigitSama(nomor int) bool {
	digitTerakhir := nomor % 10
	for nomor > 0 {
		if nomor % 10 != digitTerakhir {
			return false
		}
		nomor /= 10
	}
	return true
}

func semuaDigitBerbeda(nomor int) bool {
	for nomor > 0 {
		digit := nomor % 10  
		sisa := nomor / 10     
		temp := sisa            

		for temp > 0 {
			if temp % 10 == digit {
				return false
			}
			temp /= 10
		}
		nomor = sisa
	}
	return true
}
