package main

import "fmt"

func main() {
	var peserta, nomor, giftA, giftB, giftC int

	fmt.Print("masukkan jumlah peserta : ")
	fmt.Scan(&peserta)

	for i := 1; i <= peserta; i++ {
		fmt.Printf("masukkan nomor kartu peserta ke-%d : ", i)
		fmt.Scan(&nomor)

		switch jenisGift(nomor) {
		case "A":
			fmt.Println("Hadiah A")
			giftA++
		case "B":
			fmt.Println("Hadiah B")
			giftB++
		default:
			fmt.Println("Hadiah C")
			giftC++
		}
	}

	fmt.Println()
	fmt.Printf("jumlah yang memperoleh Hadiah A : %d\n", giftA)
	fmt.Printf("jumlah yang memperoleh Hadiah B : %d\n", giftB)
	fmt.Printf("jumlah yang memperoleh Hadiah C : %d\n", giftC)

	fmt.Print("Sheila Stephanie A [103112400086]\n")
}

func jenisGift(digit int) string {
	//cek digit sama semua
	digitTerakhir := digit % 10
	samaSemua := true
	salinan := digit

	for salinan > 0 {
		if salinan%10 != digitTerakhir {
			samaSemua = false
			break
		}
		salinan = salinan / 10
	}
	if samaSemua {
		return "A"
	}

	//cek digit beda semua
	var digitCheck int
	salinan = digit
	for salinan > 0 {
		digit2 := salinan % 10
		if digitCheck&(1<<digit2) != 0 {
			return "C" // ada yang sama
		}
		digitCheck = digitCheck | 1<<digit2
		salinan = salinan / 10
	}

	return "B" // semua digit beda
}
