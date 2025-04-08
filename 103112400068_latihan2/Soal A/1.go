package main

import "fmt"

func main() {
	var jam, menit, tarif, totaljam int
	var member bool
	var voucher string
	fmt.Print("Masukkan durasi: ")
	fmt.Scan(&jam)
	fmt.Print("Masukkan durasi(menit: )")
	fmt.Scan(&menit)
	fmt.Print("apakah anda member; ")
	fmt.Scan(&member)
	if member {
		tarif = 3500
	} else {
		tarif = 5000
	}
	if menit >= 60 {
		totaljam = jam + 1
	} else {
		totaljam = jam
	}
	totalbiaya := totaljam * tarif
	if totaljam > 3 {
		fmt.Print("masukkan voucher")
		fmt.Scan(&voucher)
	}
	fmt.Print(totalbiaya)
}
