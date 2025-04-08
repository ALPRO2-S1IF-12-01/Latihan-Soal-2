package main
import "fmt"
func hitungSewa(jam, menit, voucher int, sewa float64, member bool) float64{
	var tarif_103112400041 float64
	var totalJam float64
	var digit int
	digit = 0
	
	if member == true {
		tarif_103112400041 = 3500
	} else {
		tarif_103112400041 = 5000
	}
	for voucher != 0 {
		voucher = voucher/10
		digit = digit+1
	}
	totalJam = float64(jam) + (float64(menit) / 60)
	sewa = tarif_103112400041 * totalJam

	if jam > 3 {
		sewa = sewa - (sewa/10)
	}
	if digit == 5 {
		sewa = sewa - (sewa/10)
	}
	
	return sewa
}

func main() {
	var jam, menit, voucher int
	var sewa float64
	var member bool
	fmt.Print("Masukkan durasi (jam): ")
	fmt.Scanln(&jam)
	fmt.Print("Masukkan durasi (menit): ")
	fmt.Scanln(&menit)
	fmt.Print("Apakah member? (true/false): ")
	fmt.Scanln(&member)
	fmt.Print("Masukkan nomor voucher (jika ada): ")
	fmt.Scanln(&voucher)
	fmt.Print("Biaya sewa setelah diskon (jika memenuhi syarat): ")
	fmt.Printf("%.2f",hitungSewa(jam, menit, voucher, sewa, member))
}