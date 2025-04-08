package main

import "fmt"

func main() {

	var jam, menit int
	var member_103112400087 bool
	var voucher string

	fmt.Print("masukkan durasi (jam): ")
	fmt.Scan(&jam)

	fmt.Print("masukkan durasi (menit): ")
	fmt.Scan(&menit)

	fmt.Print("apakah member? (true/false): ")
	fmt.Scan(&member_103112400087)

	fmt.Print("Masukkan nomor voucher (jika ada): ")
	fmt.Scan(&voucher)

	jamSewa := hitungJamSewa(jam, menit)

	tarif := 5000
	if member_103112400087 {
		tarif = 3500
	}

	biaya := float64(jamSewa * tarif)

	if jamSewa >= 3 {
		persenDiskon := hitungDiskon(voucher)
		if persenDiskon > 0 {
			diskon := biaya * 0.1666666667

			biaya -= diskon
		}
	}

	fmt.Printf("Biaya sewa setelah diskon (jika memenuhi syarat): Rp %.2f\n", biaya)
}

func hitungJamSewa(jam, menit int) int {
	if jam == 0 {
		if menit < 10 {
			return 0
		}
		return 1
	}
	if menit >= 10 {
		return jam + 1
	}
	return jam
}

func hitungDiskon(voucher string) int {
	jumlah := 0
	for _, digit := range voucher {
		if digit == '5' || digit == '6' {
			jumlah++
		}
	}
	return jumlah * 10
}
