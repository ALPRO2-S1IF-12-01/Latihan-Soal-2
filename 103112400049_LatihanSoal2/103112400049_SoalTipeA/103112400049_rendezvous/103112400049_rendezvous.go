package main

//HISYAM NURDIATMOKO - 103112400049
import "fmt"

func main() {
	var x_103112400049, y int
	fmt.Print("masukkan x: ")
	fmt.Scan(&x_103112400049)
	fmt.Print("masukkan y: ")
	fmt.Scan(&y)
	count := 0
	for hari := 1; hari <= 365; hari++ {
		if hari%x_103112400049 == 0 && hari%y != 0 {
			count++
		}
	}
	fmt.Println("jumlah pertemuan dalam setahun:", count)
}
