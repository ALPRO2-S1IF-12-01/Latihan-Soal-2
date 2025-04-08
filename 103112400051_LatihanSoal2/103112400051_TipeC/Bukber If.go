//Pratama Bintang Daniswara 103112400051
package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&n)

	var a, b, c int

	for i := 1; i <= n; i++ {
		var nomor int
		fmt.Printf("Masukkan nomor kartu peserta ke-%d: ", i)
		fmt.Scan(&nomor)

		if cekSama(nomor) {
			fmt.Println("Hadiah A")
			a++
		} else if cekBeda(nomor) {
			fmt.Println("Hadiah B")
			b++
		} else {
			fmt.Println("Hadiah C")
			c++
		}
	}

	fmt.Println("Jumlah yang memperoleh Hadiah A:", a)
	fmt.Println("Jumlah yang memperoleh Hadiah B:", b)
	fmt.Println("Jumlah yang memperoleh Hadiah C:", c)
}
func cekSama(n int) bool {
	d := n % 10
	for n > 0 {
		if n%10 != d {
			return false
		}
		n = n / 10
	}
	return true
}
func cekBeda(n int) bool {
	angka := [10]int{}
	for n > 0 {
		d := n % 10
		if angka[d] == 1 {
			return false
		}
		angka[d] = 1
		n = n / 10
	}
	return true
}
