package main
//103112400078
//Mohammad Reyhan Aretha Fatin
import "fmt"

func kelipatanIteratif() {
	var bilangan, hasil int = 0, 0
	fmt.Println("Masukkan bilangan (negatif untuk berhenti): ")
	for bilangan >= 0 {
		fmt.Scan(&bilangan)
		if bilangan%4 == 0 {
			hasil += bilangan
		}
	}
	fmt.Println("Iteratif--Jumlah bilangan kelipatan 4:", hasil)
}

func kelipatanRekursif(a int) int {
	var bil int
	fmt.Scan(&bil)
	if bil < 0 {
		return a
	}
	if bil >= 0 && bil%4 == 0 {
		a += bil
	}
	return kelipatanRekursif(a)
}

func main() {
	kelipatanIteratif()
	fmt.Println("Masukkan bilangan (negatif untuk berhenti): ")
	hasil := kelipatanRekursif(0)
	fmt.Print("Rekursif--Jumlah bilangan kelipatan 4: ", hasil)
}