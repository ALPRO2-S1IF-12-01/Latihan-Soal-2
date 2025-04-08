// DWI OKTA SURYANINGRUM / 103112400066

package main

import "fmt"

// Fungsi untuk menghitung total biaya berdasarkan menu, jumlah orang, dan apakah ada sisa
func hitungBiaya(menu, orang int, sisa bool) int {
	var biaya int

	// Hitung biaya dasar tergantung jumlah menu
	switch {
	case menu <= 3:
		biaya = 10000 // Kalau menunya 3 atau kurang, tarif flat
	case menu > 50:
		biaya = 100000 // Kalau lebih dari 50, tarif maksimal
	default:
		// Kalau di antara 4-50 menu, biaya dasar ditambah per menu tambahan
		biaya = 10000 + (menu-3)*2500
	}

	// Kalau ada sisa makanan, biaya dikali jumlah orang (sebagai denda mungkin)
	if sisa {
		biaya *= orang
	}

	return biaya
}

func main() {
	var M int

	fmt.Print("Masukkan jumlah rombongan: ")
	fmt.Scan(&M)

	for i := 1; i <= M; i++ {
		var menu, orang int
		var sisaInput int
		var sisa bool

		fmt.Println("Masukkan jumlah menu, jumlah orang, dan status sisa makanan (0 untuk tidak, 1 untuk iya)")
		fmt.Print(": ")
		fmt.Scan(&menu, &orang, &sisaInput)

		// Ubah input 0/1 menjadi tipe boolean
		sisa = sisaInput != 0

		// Hitung dan tampilkan biaya
		total := hitungBiaya(menu, orang, sisa)
		fmt.Printf("Total biaya untuk rombongan %d: Rp %d\n", i, total)
	}
}
