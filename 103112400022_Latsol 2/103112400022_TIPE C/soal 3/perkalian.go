//Damanik, Yohanes Geovan Ondova
//NIM 103112400022
package main
import "fmt"

func main(){
	var n, m int
	fmt.Print("Masukan bilangan n: ")
	fmt.Scan(&n)
	fmt.Print("Masukan bilangan m: ")
	fmt.Scan(&m)

	hasil := 0
	for i := 0; i < m; i++ {
		hasil += n
	}

	fmt.Printf("Hasil dari %d x %d = %d\n", n, m, hasil)

}