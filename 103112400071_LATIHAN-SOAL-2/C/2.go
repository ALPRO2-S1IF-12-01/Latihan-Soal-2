// Raihan Adi Arba_103112400071
package main

import "fmt"

//  cek apakah semua digit sama
func semuaDigitSama(n, prev int) bool {
	if n == 0 {
		return true
	}
	curr := n % 10
	if curr != prev {
		return false
	}
	return semuaDigitSama(n/10, prev)
}

// berfungsi embandingkan dua angka dan cek apakah ada digit yang sama
func adaDigitSama(a, b int) bool {
	if a == 0 {
		return false
	}
	digitA := a % 10
	if cekDigit(b, digitA) {
		return true
	}
	return adaDigitSama(a/10, b)
}

//cek kalau ada digit tertentu ada di angka
func cekDigit(n, target int) bool {
	if n == 0 {
		return false
	}
	if n%10 == target {
		return true
	}
	return cekDigit(n/10, target)
}

//  cek apakah semua digit berbeda
func semuaDigitBerbeda(n int) bool {
	if n < 10 {
		return true
	}
	lastDigit := n % 10
	sisa := n / 10
	if cekDigit(sisa, lastDigit) {
		return false
	}
	return semuaDigitBerbeda(sisa)
}

func main() {
	var N int
	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&N)

	countA, countB, countC := 0, 0, 0

	for i := 1; i <= N; i++ {
		var nomor int
		fmt.Printf("Masukkan nomor kartu peserta ke-%d: ", i)
		fmt.Scan(&nomor)

		last := nomor % 10
		if semuaDigitSama(nomor, last) {
			fmt.Println("Hadiah A")
			countA++
		} else if semuaDigitBerbeda(nomor) {
			fmt.Println("Hadiah B")
			countB++
		} else {
			fmt.Println("Hadiah C")
			countC++
		}
	}

	fmt.Printf("\nJumlah yang memperoleh Hadiah A: %d\n", countA)
	fmt.Printf("Jumlah yang memperoleh Hadiah B: %d\n", countB)
	fmt.Printf("Jumlah yang memperoleh Hadiah C: %d\n", countC)
}
