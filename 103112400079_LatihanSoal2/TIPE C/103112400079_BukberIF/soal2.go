package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&n)

	var hadiahA, hadiahB, hadiahC_103112400079 int
	for i := 0; i < n; i++ {
		var nomorKartu int
		fmt.Printf("Masukkan nomor kartu peserta ke-%d: ", i+1)
		fmt.Scan(&nomorKartu)

		// Memeriksa jika semua digit sama (Hadiah A)
		if nomorKartu/100 == nomorKartu%10 && (nomorKartu/10)%10 == nomorKartu%10 {
			fmt.Println("Hadiah A")
			hadiahA++
		} else if isDifferentDigits(nomorKartu) {
			// Memeriksa jika semua digit berbeda (Hadiah B)
			fmt.Println("Hadiah B")
			hadiahB++
		} else {
			// Hadiah C untuk kombinasi angka
			fmt.Println("Hadiah C")
			hadiahC_103112400079++
		}
	}

	// Output jumlah hadiah yang diberikan
	fmt.Printf("\nJumlah yang memperoleh Hadiah A: %d\n", hadiahA)
	fmt.Printf("Jumlah yang memperoleh Hadiah B: %d\n", hadiahB)
	fmt.Printf("Jumlah yang memperoleh Hadiah C: %d\n", hadiahC_103112400079)
}

// Fungsi untuk memeriksa apakah semua digit berbeda
func isDifferentDigits(nomorKartu int) bool {
	digit1 := nomorKartu / 100
	digit2 := (nomorKartu / 10) % 10
	digit3 := nomorKartu % 10
	return digit1 != digit2 && digit1 != digit3 && digit2 != digit3
}
