// RYAN AKEYLA NOVIANTO WIDODO
// 103112400081

package main

import (
	"fmt"
	"strconv"
)

func tentukanHadiah(n int) string {
	nStr := strconv.Itoa(n)
	digitMap := make(map[rune]bool)
	allSame := true
	allDifferent := true

	firstDigit := rune(nStr[0])
	for _, digit := range nStr {
		if digit == firstDigit {
			allSame = false
		}
		if digitMap[digit] {
			allDifferent = false
		}
		digitMap[digit] = true
	}

	if allSame {
		return "Hadiah A"
	} else if allDifferent {
		return "Hadiah B"
	} else {
		return "Hadiah C"
	}
}
func main() {
	var n int
	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&n)

	hadiahA, hadiahB, hadiahC := 0, 0, 0
	hadiahS := make([]string, n)

	for i := 0; i < n; i++ {
		var kartu int
		fmt.Print("Masukkan nomor kartu peserta ke-%d: ", i+1)
		fmt.Scan(&kartu)

		hadiah := tentukanHadiah(kartu)
		hadiahS[i] = hadiah
		switch hadiah {
		case "Hadiah A":
			hadiahA++
		case "Hadiah B":
			hadiahB++
		case "Hadiah C":
			hadiahC++
		}
	}
	for _, hadiah := range hadiahS {
		fmt.Println(hadiah)
	}

	fmt.Printf("Jumlah yang memperoleh Hadiah A: %d\n", hadiahA)
	fmt.Printf("Jumlah yang memperoleh Hadiah B: %d\n", hadiahB)
	fmt.Printf("Jumlah yang memperoleh Hadiah C: %d\n", hadiahC)
}
