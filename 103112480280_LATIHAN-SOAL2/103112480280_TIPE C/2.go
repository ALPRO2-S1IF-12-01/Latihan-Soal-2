// Nama: Anggun Wahyu Widiyana
// NIM : 103112480280
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var jumlahPeserta int
	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&jumlahPeserta)

	hadiahA := 0
	hadiahB := 0
	hadiahC := 0

	for i := 1; i <= jumlahPeserta; i++ {
		var nomor_103112480280 int
		fmt.Printf("Masukkan nomor kartu peserta ke-%d: ", i)
		fmt.Scan(&nomor_103112480280)

		nomorStr := strconv.Itoa(nomor_103112480280)
		panjangNomor := len(nomorStr)

		semuaSama := true
		for j := 1; j < panjangNomor; j++ {
			if nomorStr[j] != nomorStr[0] {
				semuaSama = false
				break
			}
		}

		if semuaSama {
			fmt.Println("Hadiah A")
			hadiahA++
			continue
		}

		semuaBeda := true
		for j := 0; j < panjangNomor; j++ {
			for k := j + 1; k < panjangNomor; k++ {
				if nomorStr[j] == nomorStr[k] {
					semuaBeda = false
					break
				}
			}
			if !semuaBeda {
				break
			}
		}

		if semuaBeda {
			fmt.Println("Hadiah B")
			hadiahB++
		} else {
			fmt.Println("Hadiah C")
			hadiahC++
		}
	}

	fmt.Printf("Jumlah yang memperoleh Hadiah A: %d\n", hadiahA)
	fmt.Printf("Jumlah yang memperoleh Hadiah B: %d\n", hadiahB)
	fmt.Printf("Jumlah yang memperoleh Hadiah C: %d\n", hadiahC)
}