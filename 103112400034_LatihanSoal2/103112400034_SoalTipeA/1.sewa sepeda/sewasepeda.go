package main

import "fmt"

func main() {
	var jam, menit int
	var adalahMember bool
	var voucher string

	fmt.Println("NAMA: MULIA AKBAR NANDA PRATAMA\nNIM: 103112400034")
	fmt.Print("JAM: ")
	fmt.Scan(&jam)
	fmt.Print("MENIT: ")
	fmt.Scan(&menit)
	fmt.Print("MEMBER (TRUE/FALSE): ")
	fmt.Scan(&adalahMember)
	fmt.Print("KODE VOUCHER: ")
	fmt.Scan(&voucher)

	durasi := float64(jam) + float64(menit)/60
	tarif := 5000.0
	if adalahMember {
		tarif = 3500.0
	}
	totalBiaya := durasi * tarif
	if durasi >= 3 && (terhitung(voucher, '5') || terhitung(voucher, '6')) {
		diskon := totalBiaya * 0.10
		totalBiaya -= diskon
	}
	fmt.Printf("Total biaya sewa: Rp %.2f\n", totalBiaya)
}
func terhitung(teks string, karakter byte) bool {
	for i := 0; i < len(teks); i++ {
		if teks[i] == karakter {
			return true
		}
	}
	return false
}
