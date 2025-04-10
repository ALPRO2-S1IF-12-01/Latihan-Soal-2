// M. DAVI ILYAS RENALDO/103112400062 
package main

import (
	"fmt"
	"strconv"
)

func semuaSama(nomor string) bool {
	for i := 1; i < len(nomor); i++ {
		if nomor[i] != nomor[0] {
			return false
		}
	}
	return true
}

func semuaBerbeda(nomor string) bool {
	digitMap := make(map[rune]bool)
	for _, d := range nomor {
		if digitMap[d] {
			return false
		}
		digitMap[d] = true
	}
	return true
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scanln(&n)

	hadiahA := 0
	hadiahB := 0
	hadiahC := 0

	for i := 1; i <= n; i++ {
		var nomor int
		fmt.Printf("Masukkan nomor kartu peserta ke-%d: ", i)
		fmt.Scanln(&nomor)

		nomorStr := strconv.Itoa(nomor)

		if semuaSama(nomorStr) {
			fmt.Println("Hadiah A")
			hadiahA++
		} else if semuaBerbeda(nomorStr) {
			fmt.Println("Hadiah B")
			hadiahB++
		} else {
			fmt.Println("Hadiah C")
			hadiahC++
		}
	}

	fmt.Println()
	fmt.Println("Jumlah yang memperoleh Hadiah A:", hadiahA)
	fmt.Println("Jumlah yang memperoleh Hadiah B:", hadiahB)
	fmt.Println("Jumlah yang memperoleh Hadiah C:", hadiahC)
}
