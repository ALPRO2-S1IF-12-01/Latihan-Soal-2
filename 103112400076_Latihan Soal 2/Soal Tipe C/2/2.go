package main

import (
	"fmt"
)
func semuaSama(n int) bool {
	digit := n % 10
	n /= 10
	for n > 0 {
		if n%10 != digit {
			return false
		}
		n /= 10
	}
	return true
}
func semuaBerbeda(n int) bool {
	angka := 0
	for n > 0 {
		d := n % 10
		if angka&(1<<d) != 0 {
			return false
		}
		angka |= 1 << d
		n /= 10
	}
	return true
}

func main() {
	NIM := "103112400076"
	fmt.Println("NIM:", NIM)
	var n int
	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&n)
	hadiahA := 0
	hadiahB := 0
	hadiahC := 0
	for i := 1; i <= n; i++ {
		var nomor int
		fmt.Printf("Masukkan nomor kartu peserta ke-%d: ", i)
		fmt.Scan(&nomor)
		if semuaSama(nomor) {
			fmt.Println("Hadiah A")
			hadiahA++
		} else if semuaBerbeda(nomor) {
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
