package main
import "fmt"
func Kelipatan4(n int) int {
	fmt.Scan(&n)
	if n < 0 {
		return 0
	}
	if n % 4 == 0 {
		return n + Kelipatan4(n)
	}
	return Kelipatan4(n)
}
func main() {
	var n_103112400041 int
	fmt.Print("Masukkan bilangan (negatif untuk berhenti): ")
	fmt.Println()
	fmt.Print("Jumlah bilangan kelipatan 4: ",Kelipatan4(n_103112400041))
}