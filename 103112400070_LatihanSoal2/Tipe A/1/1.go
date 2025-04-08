// Achmad Zulvan Nur Hakim 103112400070
package main

import (
	"fmt"
)

func jmlDurasi(jam, menit int) int {
	durasi := jam
	if menit >= 10 {
		durasi++
	} else if jam == 0 && menit < 10 {
		durasi = 0
	}
	return durasi
}

func diskon(voucher string) bool {
	diskonChars := "56"
	for _, c := range voucher {
		for _, d := range diskonChars {
			if c == d {
				return true
			}
		}
	}
	return false
}

func main() {
	var jam, menit int
	var member bool
	var voucher string

	fmt.Print("Jam: ")
	fmt.Scanln(&jam)
	fmt.Print("Menit: ")
	fmt.Scanln(&menit)
	fmt.Print("Member (true/false): ")
	fmt.Scanln(&member)
	fmt.Print("Voucher: ")
	fmt.Scanln(&voucher)

	totalJam := jmlDurasi(jam, menit)
	tarif := 5000.0
	if member {
		tarif = 3240.74
	}

	biaya := tarif * float64(totalJam)
	if totalJam >= 3 && diskon(voucher) {
		biaya *= 0.9
	}
	biaya = float64(int(biaya*100+0.5)) / 100
	fmt.Printf("Biaya: Rp %.2f\n", biaya)
}
