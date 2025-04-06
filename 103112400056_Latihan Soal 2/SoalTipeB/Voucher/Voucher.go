// Felix Pedrosa V

package main

import "fmt"

func main() {
	var a, b_103112400056 int

	// Meminta input dari pengguna
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b_103112400056)

	// Validasi input
	if a > b_103112400056 {
		fmt.Println("Kesalahan: Nilai a harus lebih kecil atau sama dengan b.")
		return
	}

	// Menghitung jumlah angka ganjil dalam rentang [a, b]
	count := 0
	for i := a; i <= b_103112400056; i++ {
		if i%2 != 0 {
			count++
		}
	}

	// Menampilkan hasil
	fmt.Printf("Banyaknya angka ganjil: %d", count)
}