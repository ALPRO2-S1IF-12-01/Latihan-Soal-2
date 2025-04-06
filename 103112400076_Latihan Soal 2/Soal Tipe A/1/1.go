package main

import (
	"fmt"
	"math"
)
func main() {
	var jam, menit int
	var isMember bool
	var voucher string
	var NIM string = "103112400076" 
	fmt.Println("NIM :", NIM) 
	fmt.Print("Durasi (jam) : ")
	fmt.Scanln(&jam)
	fmt.Print("Durasi (menit) : ")
	fmt.Scanln(&menit)
	fmt.Print("Apakah Member? (true/false) : ")
	fmt.Scanln(&isMember)
	fmt.Print("Masukkan no voucher (jika ada) : ")
	fmt.Scanln(&voucher)
	totalJam := hitungDurasi(jam, menit)
	tarif := 5000
	if isMember {
		tarif = 3500
	}
	totalBiaya := float64(tarif * totalJam)
	if totalJam >= 3 && adaDiskon(voucher) {
		diskon := totalBiaya * 0.1666666667 // diskon 16.666...%
		totalBiaya -= diskon
	}
	totalBiaya = math.Round(totalBiaya*100) / 100
	if totalJam >= 3 && adaDiskon(voucher) {
		fmt.Printf("Biaya sewa setelah diskon (jika memenuhi syarat) : Rp %.2f\n", totalBiaya)
	} else {
		fmt.Printf("Biaya sewa : Rp %.2f\n", totalBiaya)
	}
}
func hitungDurasi(jam, menit int) int {
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
func adaDiskon(voucher string) bool {
	for _, char := range voucher {
		if char == '5' || char == '6' {
			return true
		}
	}
	return false
}
