package main

import (
	"fmt"
	"math"
)

func main() {
	var jam, menit int
	var anggota_103112400073 bool
	var kupon string
	var totalBiaya float64

	
	fmt.Print("Masukkan durasi (jam): ")
	fmt.Scan(&jam)
	fmt.Print("Masukkan durasi (menit): ")
	fmt.Scan(&menit)
	fmt.Print("Apakah member? (true/false): ")
	fmt.Scan(&anggota_103112400073)
	fmt.Print("Masukkan nomor kupon (jika ada): ")
	fmt.Scan(&kupon)

	
	totalWaktu := float64(jam) + float64(menit)/60

	
	var tarifPerJam float64
	if anggota_103112400073 {
		tarifPerJam = 3500
	} else {
		tarifPerJam = 5000
	}

	
	totalBiaya = totalWaktu * tarifPerJam

	
	if totalWaktu > 3 {
		totalBiaya *= 0.9 
	}

	
	if len(kupon) == 5 {
		totalBiaya *= 0.9 
	}

	totalBiaya = math.Round(totalBiaya*100) / 100
	fmt.Printf("Biaya sewa setelah diskon (jika memenuhi syarat): Rp %.2f\n", totalBiaya)
}
