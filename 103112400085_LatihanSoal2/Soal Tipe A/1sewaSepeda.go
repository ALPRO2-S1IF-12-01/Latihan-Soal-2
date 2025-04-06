// Anastasia Adinda N.I
package main

import (
	"fmt"
	"math"
)

func main() {
	var durasi_jam, durasi_menit int
	var anggota_103112400085 bool
	var kupon string

	fmt.Print("Masukkan durasi (jam): ")
	fmt.Scan(&durasi_jam)
	fmt.Print("Masukkan durasi (menit): ")
	fmt.Scan(&durasi_menit)
	fmt.Print("Apakah member? (true/false): ")
	fmt.Scan(&anggota_103112400085)
	fmt.Print("Masukkan nomor kupon (jika ada): ")
	fmt.Scan(&kupon)

	waktuAkhir := hitungWaktu(durasi_jam, durasi_menit)
	tarif_jam := tarifMember_103112400085(anggota_103112400085)
	biaya := hitungBiaya(waktuAkhir, tarif_jam, kupon)

	fmt.Printf("Biaya sewa setelah diskon (jika memenuhi syarat): Rp %.2f\n", biaya)
}

func hitungWaktu(jam int, menit int) float64 {
	return float64(jam) + float64(menit)/60
}

func tarifMember_103112400085(member bool) float64 {
	if member {
		return 3500
	}
	return 5000
}

func hitungBiaya(durasi float64, tarif float64, kupon string) float64 {
	biaya := durasi * tarif

	if durasi > 3 {
		biaya *= 0.9
	}
	if len(kupon) == 5 {
		biaya *= 0.9
	}

	return math.Round(biaya*100) / 100
}
