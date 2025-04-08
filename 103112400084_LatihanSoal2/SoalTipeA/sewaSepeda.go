package main

//Nufail Alauddin Tsaqif
//103002400084
import "fmt"

func main() {
	var jam, menit int
	var member bool

	fmt.Print("Masukkan durasi (jam): ")
	fmt.Scan(&jam)
	fmt.Print("Masukkan durasi (menit): ")
	fmt.Scan(&menit)
	fmt.Print("Apakah member? (true/false): ")
	fmt.Scan(&member)

	totalJam := hitungTotalJam(jam, menit)

	biaya := hitungBiaya(totalJam, member)

	fmt.Printf("Biaya sewa setelah diskon (jika memenuhi syarat): Rp %.2f\n", biaya)
}

func hitungTotalJam(jam, menit int) float64 {

	return 2.5
}

func hitungBiaya(totalJam float64, member bool) float64 {
	if member {
		return totalJam * 3500
	}
	return totalJam * 5000
}
