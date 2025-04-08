package main  
import "fmt"

func main(){
	fmt.Println("Nama : Damanik, Yohanes Geovan Ondova\n NIM : 103112400022")
	var n, nomor, a, b, c int
	fmt.Print("Masukan jumlah peserta: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Printf("Nomor kartu peserta ke-%d: ", i)
		fmt.Scan(&nomor)
		switch cekNomor(nomor) {
		case "A":
			fmt.Println("Hadiah A")
			a++
		case "B":
			fmt.Println("Hadiah B")
			b++
		default:
			fmt.Println("Hadiah C")
			c++
		}
	}

	fmt.Printf("\nJumlah Hadiah A: %d\nJumlah Hadiah B: %d\nJumlah Hadiah C: %d\n", a, b, c)
}

func cekNomor(n int) string {
	d := n % 10
	sama := true
	temp := n
	for temp > 0 {
		if temp%10 != d {
			sama = false
			break
		}
		temp /= 10
	} 
	if sama {
		return "A"
	}
	
	var cek int
	temp = n
	for temp > 0 {
		digit := temp % 10
		if cek&(1<<digit) != 0 {
			return "C" 
		}
		cek |= 1 << digit
		temp /= 10
	}
	return "B"
	
}