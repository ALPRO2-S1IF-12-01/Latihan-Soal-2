package main

import "fmt"

func main() {
	var jumlahPeserta, nomor, hadiahA, hadiahB, hadiahC int
	fmt.Println("NAMA: MULIA AKBAR NANDA PRATAMA\nNIM; 103112400034")

	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&jumlahPeserta)

	for i := 1; i <= jumlahPeserta; i++ {
		fmt.Printf("Masukkan nomor kartu peserta ke-%d: ", i)
		fmt.Scan(&nomor)

		angkaSudahAda := [10]bool{}
		digitSama := true
		sisa := nomor % 10
		angkaSudahAda[sisa] = true
		nomor /= 10

		for nomor > 0 {
			d := nomor % 10
			if d != sisa {
				digitSama = false
			}
			if angkaSudahAda[d] {
				angkaSudahAda[d] = true
			}
			angkaSudahAda[d] = true
			sisa = d
			nomor /= 10
		}

		jumlahUnik := 0
		for _, ada := range angkaSudahAda {
			if ada {
				jumlahUnik++
			}
		}

		if digitSama {
			fmt.Println("Hadiah A")
			hadiahA++
		} else if jumlahUnik >= 3 {
			fmt.Println("Hadiah B")
			hadiahB++
		} else {
			fmt.Println("Hadiah C")
			hadiahC++
		}
	}

	fmt.Println("\nJumlah peserta yang mendapat:")
	fmt.Println("Hadiah A:", hadiahA)
	fmt.Println("Hadiah B:", hadiahB)
	fmt.Println("Hadiah C:", hadiahC)
}
