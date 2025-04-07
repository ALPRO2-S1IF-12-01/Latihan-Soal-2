// MUHAMMAD GAMEL AL GHIFARI
// 103112400028
package main

import "fmt"

func main() {
	peserta := inputDataPeserta()
	hadiahA, hadiahB, hadiahC := hitungHadiah(peserta)
	tampilkanRingkasan(hadiahA, hadiahB, hadiahC)
}

func inputDataPeserta() []string {
	var jumlahPeserta int
	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&jumlahPeserta)

	peserta := make([]string, jumlahPeserta)
	for i := 0; i < jumlahPeserta; i++ {
		fmt.Printf("Masukkan nomor kartu peserta ke-%d: ", i+1)
		fmt.Scan(&peserta[i])
	}
	return peserta
}

func hitungHadiah(peserta []string) (int, int, int) {
	var a, b, c int
	for _, nomor := range peserta {
		hadiah := tentukanJenisHadiah(nomor)
		fmt.Println("Hadiah", hadiah)

		switch hadiah {
		case "A":
			a++
		case "B":
			b++
		case "C":
			c++
		}
	}
	return a, b, c
}

func tentukanJenisHadiah(nomor string) string {
	if semuaDigitSama(nomor) {
		return "A"
	}
	if semuaDigitBeda(nomor) {
		return "B"
	}
	return "C"
}

func semuaDigitSama(nomor string) bool {
	for i := 1; i < len(nomor); i++ {
		if nomor[i] != nomor[0] {
			return false
		}
	}
	return true
}

func semuaDigitBeda(nomor string) bool {
	seen := make(map[byte]bool)
	for i := 0; i < len(nomor); i++ {
		if seen[nomor[i]] {
			return false
		}
		seen[nomor[i]] = true
	}
	return true
}

func tampilkanRingkasan(a, b, c int) {
	fmt.Println("\nJumlah yang memperoleh Hadiah A:", a)
	fmt.Println("Jumlah yang memperoleh Hadiah B:", b)
	fmt.Println("Jumlah yang memperoleh Hadiah C:", c)
}
