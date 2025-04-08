package main

import "fmt"

func semuaDigitSama(n int) bool {
	digitTerakhir := n % 10 
	for n > 0 {
		if n%10 != digitTerakhir {
			return false
		}
		n /= 10 
	}
	return true
}

func semuaDigitBerbeda(n int) bool {
	digitAda := [10]bool{}

	for n > 0 {
		digit := n % 10
		if digitAda[digit] {
			return false
		}
		digitAda[digit] = true
		n /= 10
	}
	return true
}

func main() {
	fmt.Println("Nama: Raja Muhammad Lufhti\nNim : 103112400027")
	var jumlahPeserta int
	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&jumlahPeserta)

	jumlahA, jumlahB, jumlahC := 0, 0, 0

	for i := 1; i <= jumlahPeserta; i++ {
		var nomorKartu int
		fmt.Printf("Masukkan nomor kartu peserta ke-%d: ", i)
		fmt.Scan(&nomorKartu)

		if semuaDigitSama(nomorKartu) {
			fmt.Println("Hadiah A")
			jumlahA++
		} else if semuaDigitBerbeda(nomorKartu) {
			fmt.Println("Hadiah B")
			jumlahB++
		} else {
			fmt.Println("Hadiah C")
			jumlahC++
		}
	}

	fmt.Println("\nJumlah yang memperoleh Hadiah A:", jumlahA)
	fmt.Println("Jumlah yang memperoleh Hadiah B:", jumlahB)
	fmt.Println("Jumlah yang memperoleh Hadiah C:", jumlahC)
}
