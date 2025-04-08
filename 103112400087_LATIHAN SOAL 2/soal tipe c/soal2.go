package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&n)

	a, b, c := 0, 0, 0

	for i := 0; i < n; i++ {
		var num_103112400087 int
		fmt.Printf("Masukkan nomor kartu peserta ke-%d: ", i+1)
		fmt.Scan(&num_103112400087)

		digits := make(map[rune]bool)
		for _, d := range fmt.Sprintf("%d", num_103112400087) {
			digits[d] = true
		}

		switch len(digits) {
		case 1:
			fmt.Println("Hadiah A")
			a++
		case len(fmt.Sprintf("%d", num_103112400087)):
			fmt.Println("Hadiah B")
			b++
		default:
			fmt.Println("Hadiah C")
			c++
		}
	}

	fmt.Printf("\nJumlah yang memperoleh Hadiah A: %d\n", a)
	fmt.Printf("Jumlah yang memperoleh Hadiah B: %d\n", b)
	fmt.Printf("Jumlah yang memperoleh Hadiah C: %d\n", c)
}
