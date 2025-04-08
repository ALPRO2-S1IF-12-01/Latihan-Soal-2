package main

//HISYAM NURDIATMOKO - 103112400049
import "fmt"

func main() {
	var n_103112400049, nomor, hadiahA, hadiahB, hadiahC int
	fmt.Print("masukkan jumlah peserta: ")
	fmt.Scan(&n_103112400049)
	for i := 1; i <= n_103112400049; i++ {
		fmt.Printf("masukkan nomor kartu peserta ke-%d: ", i)
		fmt.Scan(&nomor)
		digit1 := nomor / 100
		digit2 := (nomor / 10) % 10
		digit3 := nomor % 10
		if digit1 == digit2 && digit2 == digit3 {
			fmt.Println("hadiah A")
			hadiahA++
		} else if digit1 != digit2 && digit2 != digit3 && digit1 != digit3 {
			fmt.Println("hadiah B")
			hadiahB++
		} else {
			fmt.Println("hadiah C")
			hadiahC++
		}
	}
	fmt.Println("jumlah yang memperoleh hadiah A:", hadiahA)
	fmt.Println("jumlah yang memperoleh hadiah B:", hadiahB)
	fmt.Println("jumlah yang memperoleh hadiah C:", hadiahC)
}
