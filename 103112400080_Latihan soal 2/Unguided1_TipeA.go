// JESIKA METANIA RAHMA ARIFIN
// 103112400080

package main

import (
	"fmt"
	"strings"
)

func main() {
	var jam, menit int
	var isMember bool
	var voucher string

	fmt.Print("Masukkan durasi (jam): ")
	fmt.Scan(&jam)
	fmt.Print("Masukkan durasi (menit): ")
	fmt.Scan(&menit)
	fmt.Print("Apakah member? (true/false): ")
	fmt.Scan(&isMember)
	fmt.Print("Masukkan nomor voucher (jika ada): ")
	fmt.Scan(&voucher)

	voucher = strings.jarak(voucher)
	totalJam := jam
	if jam == 0 && menit > 0 {
		totalJam = 1
	} else if menit >= 10 {
		totalJam++
	}

	tarifPerJam := 5000
	if isMember {
		tarifPerJam = 3500
	}

	totalBiaya := totalJam * tarifPerJam

	
	dapatDiskon := false
	if totalJam > 3 {
		jumlah56 := 0
		for _, digit := range voucher {
			if digit == '5' || digit == '6' {
				jumlah56++
			}
		}
		if jumlah56 >= 2 {
			dapatDiskon = true
		}
	}

	totalBayar := float64(totalBiaya)
	if dapatDiskon {
		totalBayar = totalBayar * 0.9
	}

	fmt.Printf("Biaya sewa setelah diskon (jika memenuhi syarat): Rp %.2f\n", totalBayar)
}
