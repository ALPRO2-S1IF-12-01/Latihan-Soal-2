package main

import "fmt"

func main() {
	var jumlahPeserta_103112400077, nomorKartu, jumlahA, jumlahB, jumlahC int
	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&jumlahPeserta_103112400077)

	for i := 1; i <= jumlahPeserta_103112400077; i++ {
		fmt.Printf("Masukkan nomor kartu peserta ke-%d: ", i)
		fmt.Scan(&nomorKartu)

		semuaSama, semuaBerbeda := true, true
		angkaTerakhir := nomorKartu % 10
		angkaYangSudahDilihat := 0

		for sementara := nomorKartu; sementara > 0; sementara /= 10 {
			angka := sementara % 10

			if angka != angkaTerakhir {
				semuaSama = false
			}

			angkaDitemukan := false
			for pengecekan := angkaYangSudahDilihat; pengecekan > 0; pengecekan /= 10 {
				if pengecekan%10 == angka {
					angkaDitemukan = true
					break
				}
			}
			if angkaDitemukan {
				semuaBerbeda = false
			}

			angkaYangSudahDilihat = angkaYangSudahDilihat*10 + angka
		}

		if semuaSama {
			fmt.Println("Hadiah A")
			jumlahA++
		} else if semuaBerbeda {
			fmt.Println("Hadiah B")
			jumlahB++
		} else {
			fmt.Println("Hadiah C")
			jumlahC++
		}
	}

	fmt.Printf("\nJumlah yang memperoleh Hadiah A: %d\n", jumlahA)
	fmt.Printf("Jumlah yang memperoleh Hadiah B: %d\n", jumlahB)
	fmt.Printf("Jumlah yang memperoleh Hadiah C: %d\n", jumlahC)
}
