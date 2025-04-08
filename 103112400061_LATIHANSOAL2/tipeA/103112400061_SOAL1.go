package main

import "fmt"

func main() {
	var (
		jam, menit float32
		member_103112400061 bool
		voucher int
	)
	fmt.Print("Masukkan durasi (jam): ")
	fmt.Scan(&jam)
	fmt.Print("Masukkan durasi (menit): ")
	fmt.Scan(&menit)
	fmt.Print("Apakah member? (true/false): ")
	fmt.Scan(&member_103112400061)
	fmt.Print("Masukkan nomor voucher (jika ada): ")
	fmt.Scan(&voucher) // NIM: 103112400061
	biaya := sewaSepeda(jam, menit, member_103112400061, voucher)
	fmt.Printf("Biaya sewa setelah diskon (jika memenuhi syarat): Rp %.2f\n", biaya)
}

func sewaSepeda(jam, menit float32, member_103112400061 bool, voucher int) float32 {
	var biaya float32
	if menit > 10 {
		jam = jam + 1
	} // AUTHOR: KEISHIN NAUFA ALFARIDZHI
	if member_103112400061 {
		biaya = jam * 3500
	} else {
		biaya = jam * 5000
	}
	if voucher > 9999 && voucher < 1000000 {
		biaya = biaya - (biaya * 0.1)
	}

	return biaya
}