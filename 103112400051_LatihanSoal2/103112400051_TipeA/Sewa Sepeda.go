//Pratama Bintang Daniswara 103112400051
package main

import "fmt"

func main() {
	var jam, menit int
	var member bool
	var voucher string
	fmt.Print("Masukan Durasi (jam): ")
	fmt.Scanln(&jam)
	fmt.Print("Masukan Durasi (Menit): ")
	fmt.Scanln(&menit)
	fmt.Print("Member (true/false): ")
	fmt.Scanln(&member)
	fmt.Print("Masukan Nomor Voucher (jika ada): ")
	fmt.Scanln(&voucher)

	totalJam := durasi(jam, menit)

	var tarif float64
	switch member {
	case true:
		tarif = 3240.74
	default:
		tarif = 5000.0
	}
	biaya := tarif * float64(totalJam)
	switch {
	case totalJam >= 3 && diskon(voucher):
		biaya = biaya - (biaya * 0.10)
	}
	biaya = float64(int(biaya*100+0.5)) / 100
	fmt.Printf("Biaya: Rp %.2f\n", biaya)
}
func durasi(h, m int) int {
	switch {
	case h == 0 && m < 10:
		return 0
	case m >= 10:
		return h + 1
	default:
		return h
	}
}
func diskon(v string) bool {
	for _, c := range v {
		switch c {
		case '5', '6':
			return true
		default:
		}
	}
	return false
}
