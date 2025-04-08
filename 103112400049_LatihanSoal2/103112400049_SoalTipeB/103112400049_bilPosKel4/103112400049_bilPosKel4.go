package main

//HISYAM NURDIATMOKO - 103112400049
import "fmt"

func bilPosKel4(nomor []int, idx, total int) int {
	if idx >= len(nomor) {
		return total
	}
	if nomor[idx] > 0 && nomor[idx]%4 == 0 {
		total += nomor[idx]
	}
	return bilPosKel4(nomor, idx+1, total)
}

func main() {
	fmt.Println("Masukkan bilangan (negatif untuk berhenti):")
	var nomor []int
	var n_103112400049 int
	for {
		fmt.Scan(&n_103112400049)
		if n_103112400049 < 0 {
			break
		}
		nomor = append(nomor, n_103112400049)
	}
	total := bilPosKel4(nomor, 0, 0)
	fmt.Printf("Jumlah bilangan kelipatan 4: %d\n", total)
}
