package main

//HISYAM NURDIATMOKO - 103112400049
import "fmt"

func main() {
	var n_103112400049 int
	fmt.Print("masukkan bilangan bulat positif (>10): ")
	fmt.Scan(&n_103112400049)
	n := n_103112400049
	panjang := 0
	for n > 0 {
		n /= 10
		panjang++
	}
	tengah := panjang / 2
	a := 1
	for i := 0; i < panjang-tengah; i++ {
		a *= 10
	}
	bil1 := n_103112400049 / a
	bil2 := n_103112400049 % a
	fmt.Printf("bilangan 1: %d\n", bil1)
	fmt.Printf("bilangan 2: %d\n", bil2)
	fmt.Printf("hasil penjumlahan: %d\n", bil1+bil2)
}
