package main
//103112400078
//Mohammad Reyhan Aretha Fatin
import "fmt"

func perkalianIteratif(n, m int) int {
	var hasil int
	for i := n; i >= 1; i-- {
		hasil += m
	}
	return hasil
}

func perkalianRekursif(n, m, hasil int) int {
	if n == 0 {
		return hasil
	}
	return perkalianRekursif(n-1, m, hasil+m)
}

func main() {
	var n, m int
	fmt.Print("Masukkan bilangan n: ")
	fmt.Scan(&n)
	fmt.Print("Masukkan bilangan m: ")
	fmt.Scan(&m)
	hasilIteratif := perkalianIteratif(n, m)
	hasilRekursif := perkalianRekursif(n, m, 0)
	fmt.Printf("Iteratif---Hasil dari %d x %d =%d\n", n, m, hasilIteratif)
	fmt.Printf("Rekursif---Hasil dari %d x %d =%d\n", n, m, hasilRekursif)
}