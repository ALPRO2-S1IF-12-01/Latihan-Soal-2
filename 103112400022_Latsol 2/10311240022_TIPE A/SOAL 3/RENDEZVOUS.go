package main 
import "fmt"

func rendezvous(x, y int)int {
	j := 0
	for hari := 1; hari <= 365; hari++ {
		if hari%x == 0 && hari%y != 0 {
			j++
		}
	}
	return j
}
func main() {
	fmt.Println("Nama : Damanik, Yohanes Geovan Ondova\n NIM: 103112400022")
	var x, y int 
	fmt.Print("Masukan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukan nilai y: ")
	fmt.Scan(&y)
	fmt.Print("Jumlah pertemuan dalam setahun: ", rendezvous(x,y))
}