package main

//HISYAM NURDIATMOKO - 103112400049
import "fmt"

func kali(n, m int) int {
	if n == 0 {
		return 0
	}
	return m + kali(n-1, m)
}

func main() {
	var a_103112400049, b int
	fmt.Print("masukkan bilangan n: ")
	fmt.Scan(&a_103112400049)
	fmt.Print("masukkan bilangan m: ")
	fmt.Scan(&b)
	hasil := kali(a_103112400049, b)
	fmt.Printf("hasil dari %d x %d = %d", a_103112400049, b, hasil)
}
