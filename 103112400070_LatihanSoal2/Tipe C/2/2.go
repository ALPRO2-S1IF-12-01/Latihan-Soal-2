// Achmad Zulvan Nur Hakim 103112400070
package main

import "fmt"
func angkaSama(n int) bool {
	angka := n % 10
	for n > 0 {
		if n%10 != angka {
			return false
		}
		n /= 10
	}
	return true
}
func angkabeda(n int) bool {
	angkasudahAda := [10]bool{}
	for n > 0 {
		sisa := n % 10
		if angkasudahAda[sisa] {
			return false
		}
		angkasudahAda[sisa] = true
		n /= 10
	}
	return true
}

func main() {
	var n, nomor int
	var a, b, c int
	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Printf("Masukkan nomor kartu peserta ke-%d: ", i)
		fmt.Scan(&nomor)
		if angkaSama(nomor) {
			fmt.Println("Hadiah A")
			a++
		} else if angkabeda(nomor) {
			fmt.Println("Hadiah B")
			b++
		} else {
			fmt.Println("Hadiah C")
			c++
		}
	}
	fmt.Printf("Jumlah yang memperoleh Hadiah A: %d\n", a)
	fmt.Printf("Jumlah yang memperoleh Hadiah B: %d\n", b)
	fmt.Printf("Jumlah yang memperoleh Hadiah C: %d\n", c)
}
