// Felix Pedrosa V

package main

import "fmt"

// Fungsi rekursif untuk menjumlahkan bilangan positif kelipatan 4
func jumlahKelipatanEmpat(angka []int, indeks int) int {
	// Jika sudah mencapai akhir array atau bilangan negatif (sentinel), kembalikan 0
	if indeks >= len(angka) || angka[indeks] < 0 {
		return 0
	}

	// Cek jika bilangan adalah positif kelipatan 4, tambahkan ke hasil
	if angka[indeks]%4 == 0 && angka[indeks] > 0 {
		return angka[indeks] + jumlahKelipatanEmpat(angka, indeks+1)
	} else {
		// Jika bukan kelipatan 4, lanjutkan ke bilangan berikutnya
		return jumlahKelipatanEmpat(angka, indeks+1)
	}
}

func main() {
	var input_103112400056 int
	angka := []int{}

	fmt.Println("Masukkan bilangan (negatif untuk berhenti):")

	// Ambil input dari pengguna hingga bilangan negatif dimasukkan
	for {
		fmt.Scan(&input_103112400056)
		if input_103112400056 < 0 {
			break
		}
		angka = append(angka, input_103112400056)
	}

	// Hitung jumlah bilangan kelipatan 4 menggunakan fungsi rekursif
	hasil := jumlahKelipatanEmpat(angka, 0)

	// Tampilkan hasil
	fmt.Printf("Jumlah bilangan kelipatan 4: %d\n", hasil)
}