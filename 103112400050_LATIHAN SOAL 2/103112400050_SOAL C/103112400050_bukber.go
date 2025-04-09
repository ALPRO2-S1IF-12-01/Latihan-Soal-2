package main

import "fmt"

//103112400050
func main() {
	var peserta, nomor int
	var a, b, c int

	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&peserta)

	for i := 1; i <= peserta; i++ {
		fmt.Printf("Masukkan nomor kartu peserta ke-%d: ", i)
		fmt.Scan(&nomor)

		d1 := nomor / 100
		d2 := (nomor / 10) % 10
		d3 := nomor % 10

		switch {
		case d1 == d2 && d2 == d3:
			fmt.Println("Hadiah A")
			a++
		case d1 != d2 && d2 != d3 && d1 != d3:
			fmt.Println("Hadiah B")
			b++
		default:
			fmt.Println("Hadiah C")
			c++
		}
	}

	fmt.Println("Jumlah yang memperoleh Hadiah A:", a)
	fmt.Println("Jumlah yang memperoleh Hadiah B:", b)
	fmt.Println("Jumlah yang memperoleh Hadiah C:", c)
}
