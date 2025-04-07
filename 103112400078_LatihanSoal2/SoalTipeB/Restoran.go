package main
//103112400078
//Mohammad Reyhan Aretha Fatin
import "fmt"

func main() {
	var rombongan int
	fmt.Print("Masukkan jumlah rombongan: ")
	fmt.Scan(&rombongan)
	totalBayar(rombongan)

}

func totalBayar(a int) {
	var jumlahMenu, banyakOrang, sisa, harga int
	for i := 1; i <= a; i++ {
		fmt.Print("Masukkan jumlah menu, jumlah orang, dan status sisa makanan (0 untuk tidak, 1 untuk iya): ")
		fmt.Scan(&jumlahMenu, &banyakOrang, &sisa)
		if jumlahMenu > 50 {
			harga = 100000
		} else if jumlahMenu > 3 {
			harga = 10000 + 2500*jumlahMenu
		} else if jumlahMenu <= 3 {
			harga = 10000
		}

		if sisa == 1 {
			harga *= banyakOrang
		}
		fmt.Printf("Total biaya untuk rombongan %d: Rp %d \n", i, harga)
	}
}