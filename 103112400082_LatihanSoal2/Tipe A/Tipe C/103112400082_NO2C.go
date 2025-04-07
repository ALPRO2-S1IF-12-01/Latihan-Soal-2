// Rizkina Azizah_103112400082
package main

import "fmt"

func main() {
	var N int
	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&N)

	hadiahA := 0
	hadiahB := 0
	hadiahC := 0

	for i := 1; i <= N; i++ {
		var nomorKartu int
		fmt.Printf("Masukkan nomor kartu peserta ke-%d: ", i)
		fmt.Scan(&nomorKartu)

		if semuaDigitSama(nomorKartu) {
			fmt.Println("Hadiah A")
			hadiahA++
		} else if semuaDigitBerbeda(nomorKartu) {
			fmt.Println("Hadiah B")
			hadiahB++
		} else {
			fmt.Println("Hadiah C")
			hadiahC++
		}
	}

	fmt.Printf("\nJumlah yang memperoleh Hadiah A: %d\n", hadiahA)
	fmt.Printf("Jumlah yang memperoleh Hadiah B: %d\n", hadiahB)
	fmt.Printf("Jumlah yang memperoleh Hadiah C: %d\n", hadiahC)
}

func semuaDigitSama(M int) bool {
	var a, b, c int
	a = M / 100
	b = (M /10) % 10
	c = M % 10
	return a == b && b == c

	}
	
func semuaDigitBerbeda(M int) bool {
	var a, b, c int
	a = M / 100
	b = (M /10) % 10
	c = M % 10
	
	return a != b && b != c && a != c
	}
	
