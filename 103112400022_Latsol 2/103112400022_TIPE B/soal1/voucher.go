package main 
import "fmt"

func main(){
	fmt.Println("Nama: Damanik, Yohanes Geovan Ondova\n NIM: 103112400022")
	var a, b int 
	fmt.Print("Masukan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukan nilai b: ")
	fmt.Scan(&b)

	jumlah := 0
	for i := a; i <= b; i++ {
		if i%2 != 0 {
			jumlah++
		}
	}
	fmt.Println("Banyaknya angka ganjil:", jumlah)
}