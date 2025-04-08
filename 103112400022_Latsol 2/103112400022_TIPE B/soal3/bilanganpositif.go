package main 
import "fmt"

func main (){
	fmt.Println("Damanik. Yohanes Geovan Ondova\n NIM : 103112400022")
	fmt.Print("Masukan bilangan (negatif untuk behenti):\n")
	total := hasil(0)
	fmt.Println("Jumlah bilangan ke 4:", total)
}

func hasil(total int) int {
	var a int
	fmt.Scan(&a)
	if a < 0 {
		return total
	}
	if a > 0 && a%4 == 0 {
		total += a
	}
	return hasil(total)
}