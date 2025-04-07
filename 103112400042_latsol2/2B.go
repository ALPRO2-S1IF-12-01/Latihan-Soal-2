// M.HANIF AL FAIZ
// 103112400042
package main

import "fmt"

func main() {
	var M int
	fmt.Print("Masukkan jumlah rombongan: ")
	fmt.Scan(&M)
	results := make([]int, M)
	for i := 0; i < M; i++ {
		var menu, orang int
		var sisa bool
		fmt.Printf("\nData rombongan ke-%d:\n", i+1)
		fmt.Print("Jumlah menu: ")
		fmt.Scan(&menu)
		fmt.Print("Jumlah orang: ")
		fmt.Scan(&orang)
		fmt.Print("Ada sisa? (true/false): ")
		fmt.Scan(&sisa)
		results[i] = hitungBiaya(menu, orang, sisa)
	}
	fmt.Println("\nHasil Perhitungan Biaya:")
	for i, total := range results {
		fmt.Printf("Rombongan ke-%d: Rp %d\n", i+1, total)
	}
}
func hitungBiaya(menu, orang int, sisa bool) int {
	const (
		hargaDasar    = 10000
		hargaTambahan = 2500
		hargaBulk     = 100000
		batasDasar    = 3
		batasBulk     = 50
	)
	var total int
	switch {
	case menu > batasBulk:
		total = hargaBulk
	case menu > batasDasar:
		total = hargaDasar + (menu-batasDasar)*hargaTambahan
	default:
		total = hargaDasar
	}
	if sisa {
		total *= orang
	}
	return total
}
