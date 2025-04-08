package main

import "fmt"

func main() {
	var jumlahPeserta, hadiahA, hadiahB, hadiahC int

	fmt.Println("PROGRAM PENENTU HADIAH")
	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&jumlahPeserta)

	for i := 1; i <= jumlahPeserta; i++ {
		var nomorKartu int

		fmt.Printf("\nPeserta ke-%d\n", i)

		for {
			fmt.Scan(&nomorKartu)
			if nomorKartu >= 100 && nomorKartu <= 999 {
				break
			}

		}

		digit1 := nomorKartu / 100
		digit2 := (nomorKartu / 10) % 10
		digit3 := nomorKartu % 10

		if digit1 == digit2 && digit2 == digit3 {
			fmt.Println("Hadiah A ")
			hadiahA++
		} else if digit1 != digit2 && digit1 != digit3 && digit2 != digit3 {
			fmt.Println("Hadiah B ")
			hadiahB++
		} else {
			fmt.Println("Hadiah C ")
			hadiahC++
		}
	}

	fmt.Println("\nREKAP HADIAH:")
	fmt.Printf("Hadiah A: %d peserta\n", hadiahA)
	fmt.Printf("Hadiah B: %d peserta\n", hadiahB)
	fmt.Printf("Hadiah C: %d peserta\n", hadiahC)
}
