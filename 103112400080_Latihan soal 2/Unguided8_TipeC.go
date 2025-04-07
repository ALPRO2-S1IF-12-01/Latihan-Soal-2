// JESIKA METANIA RAHMA ARIFIN
// 103112400080

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
		var nomorKartu string
		fmt.Printf("Masukkan nomor kartu peserta ke-%d: ", i)
		fmt.Scan(&nomorKartu)

		if semuaAngkaSama(nomorKartu) {
			fmt.Println("Hadiah A")
			jumlahHadiahA++
		} else if semuaAngkaBerbeda(nomorKartu) {
			fmt.Println("Hadiah B")
			jumlahHadiahB++
		} else {
			fmt.Println("Hadiah C")
			jumlahHadiahC++
		}
	}

	fmt.Printf("\nHadiah A: %d\nHadiah B: %d\nHadiah C: %d\n", jumlahHadiahA, jumlahHadiahB, jumlahHadiahC)
}

func semuaAngkaSama(kartu string) bool {
	for i := 1; i < len(kartu); i++ {
		if kartu[i] != kartu[0] {
			return false
		}
	}
	return true
}

func semuaAngkaBerbeda(kartu string) bool {
	angkaSudahAda := make(map[rune]bool)
	for _, angka := range kartu {
		if angkaSudahAda[angka] {
			return false
		}
		angkaSudahAda[angka] = true
	}
	return true
}
