package main

import "fmt"

func main() {
	var (
		x, n_103112400061 int
		jumlahHadiahA, jumlahHadiahB, jumlahHadiahC int
	)
	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&n_103112400061)
	for i := 1; i <= n_103112400061; i++ {
		fmt.Printf("Masukkan nomor kartu peserta ke-%v: ", i)
		fmt.Scan(&x)

		result := cekHadiah(x)

		if result == "Hadiah A" {
			jumlahHadiahA++
		} else if result == "Hadiah B" {
			jumlahHadiahB++
		} else if result == "Hadiah C" {
			jumlahHadiahC++
		}

		fmt.Println(result)
	}
	
	fmt.Println()
	fmt.Println("Jumlah yang memperoleh Hadiah A:", jumlahHadiahA)
	fmt.Println("Jumlah yang memperoleh Hadiah B:", jumlahHadiahB)
	fmt.Println("Jumlah yang memperoleh Hadiah C:", jumlahHadiahC)
}

func cekHadiah(x int) string {
	var digits []int
	var temp int = x
	var allSama bool = true
	var allDiff bool = true

	for temp > 0 {
		digits = append([]int{temp%10}, digits...)
		temp /= 10
	}

	for i := 1; i < len(digits); i++ {
		if digits[i] != digits[0] {
			allSama = false
			break
		}
	}

	if allSama {
		return "Hadiah A"
	}
		for i := 0; i < len(digits); i++ {
			for j := i + 1; j < len(digits); j++ {
				if digits[i] == digits[j] {
					allDiff = false
					break
				}
			}
			if !allDiff {
				break
			}
		}

	if allDiff {
		return "Hadiah B"
	} else {
		return "Hadiah C"
	}
}

