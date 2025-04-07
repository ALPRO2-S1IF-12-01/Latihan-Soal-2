package main

import (
	"fmt"
	"math"
)

func hitungBiaya(durasiJam, durasiMenit int, member bool, voucher string) float64 {
	const tarifMember = 3500.0
	const tarifBiasa = 5000.0

	totalDurasi := float64(durasiJam) + float64(durasiMenit)/60

	tarif := tarifBiasa
	if member {
		tarif = tarifMember
	}

	//biaya awal
	total := totalDurasi * tarif

	if totalDurasi > 3 {
		total = total * 0.9
	}

	if len(voucher) == 5 {
		total = total * 0.9
	}

	// pembulatan
	total = math.Round(total*100) / 100
	return total
}

func main() {
	var jam, menit int
	var status, vocs string
	var total float64

	fmt.Print("masukkan durasi jam : ")
	fmt.Scan(&jam)

	fmt.Print("masukkan durasi menit : ")
	fmt.Scan(&menit)

	fmt.Print(" apakah member (true/false): ")
	fmt.Scan(&status)
	memberBukan := (status == "true")

	fmt.Print(" masukkan nomor voucher (kalau ada): ")
	fmt.Scan(&vocs)

	total = hitungBiaya(jam, menit, memberBukan, vocs)

	fmt.Printf("biaya sewa setelah discount (jika memenuhi syarat): Rp %.2f\n", total)
	fmt.Print("Sheila Stephanie A [103112400086]")
}
