// M.HANIF AL FAIZ
// 103112400042
package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&N)

	nomorKartu := make([]string, N)

	countA, countB, countC := 0, 0, 0

	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan nomor kartu peserta ke-%d: ", i+1)
		fmt.Scan(&nomorKartu[i])

		hadiah := tentukanHadiah(nomorKartu[i])
		fmt.Println("Hadiah", hadiah)

		switch hadiah {
		case "A":
			countA++
		case "B":
			countB++
		case "C":
			countC++
		}
	}

	fmt.Println("\nJumlah yang memperoleh Hadiah A:", countA)
	fmt.Println("Jumlah yang memperoleh Hadiah B:", countB)
	fmt.Println("Jumlah yang memperoleh Hadiah C:", countC)
}

func tentukanHadiah(nomor string) string {

	semuaSama := true
	for i := 1; i < len(nomor); i++ {
		if nomor[i] != nomor[0] {
			semuaSama = false
			break
		}
	}
	if semuaSama {
		return "A"
	}

	semuaBeda := true
	digitMap := make(map[byte]bool)
	for i := 0; i < len(nomor); i++ {
		if digitMap[nomor[i]] {
			semuaBeda = false
			break
		}
		digitMap[nomor[i]] = true
	}
	if semuaBeda {
		return "B"
	}

	return "C"
}
