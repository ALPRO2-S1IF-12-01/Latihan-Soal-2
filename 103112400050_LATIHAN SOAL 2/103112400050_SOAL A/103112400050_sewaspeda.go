package main

//103112400050
import "fmt"

func main() {
	var jam, menit int
	var member_103112400050 string
	var kupon string
	fmt.Print("masukkan Durasi (jam): ")
	fmt.Scan(&jam)
	fmt.Print("masukkan Durasi (menit): ")
	fmt.Scan(&menit)
	fmt.Print("apakah member? (true/false): ")
	fmt.Scan(&member_103112400050)
	fmt.Print("masukkan nomor kupon (jika ada): ")
	fmt.Scan(&kupon)
	member := member_103112400050 == "true"
	durasi := hitungDurasi(jam, menit)
	tarif := 5000.0
	if member {
		tarif = 3500.0
	}
	biaya := durasi * tarif
	if durasi > 3 && (len(kupon) == 5 || len(kupon) == 6) {
		biaya *= 0.9
	}
	fmt.Printf("biaya sewa setelah diskon (jika memenuhi syarat): Rp %.2f\n", biaya)
}
func hitungDurasi(jam, menit int) float64 {
	totalMenit := jam*60 + menit
	if totalMenit < 60 {
		return 1.0
	}
	totalJam := totalMenit / 60
	sisaMenit := totalMenit % 60
	if sisaMenit >= 10 {
		return float64(totalJam) + float64(sisaMenit)/60.0
	}
	return float64(totalJam)
}
