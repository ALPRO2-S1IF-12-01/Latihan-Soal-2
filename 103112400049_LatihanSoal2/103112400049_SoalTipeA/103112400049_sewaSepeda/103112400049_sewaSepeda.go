package main

//HISYAM NURDIATMOKO - 103112400049
import "fmt"

func main() {
	var jam, menit int
	var member_103112400049, voucher string
	fmt.Print("masukkan durasi (jam): ")
	fmt.Scan(&jam)
	fmt.Print("masukkan durasi (menit): ")
	fmt.Scan(&menit)
	fmt.Print("apakah member? (true/false): ")
	fmt.Scan(&member_103112400049)
	fmt.Print("masukkan nomor voucher (jika ada): ")
	fmt.Scan(&voucher)
	member := member_103112400049 == "true"
	durasi := hitungDurasi(jam, menit)
	tarif := 5000.0
	if member {
		tarif = 3500.0
	}
	biaya := durasi * tarif
	if durasi > 3 && (len(voucher) == 5 || len(voucher) == 6) {
		biaya *= 0.9
	}
	fmt.Printf("biaya sewa setelah diskon (jika memenuhi syarat): Rp %.2f\n", biaya)
}

func hitungDurasi(jam, menit int) float64 {
	totalMenit := jam*60 + menit
	if totalMenit < 60 {
		return 1.0
	}
	jamTotal := totalMenit / 60
	sisaMenit := totalMenit % 60
	if sisaMenit >= 10 {
		return float64(jamTotal) + float64(sisaMenit)/60.0
	}
	return float64(jamTotal)
}
