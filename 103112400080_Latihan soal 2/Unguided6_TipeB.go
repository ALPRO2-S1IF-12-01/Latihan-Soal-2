// JESIKA METANIA RAHMA ARIFIN
// 103112400080

package main

import "fmt"

func Kelipatan4(nums []int, index int) int {
	if index == len(nums) {
		return 0
	}
	if nums[index]%4 == 0 {
		return nums[index] + Kelipatan4(nums, index+1)
	}
	return Kelipatan4(nums, index+1)
}

func main() {
	var nums []int
	var num int

	fmt.Println("Masukkan bilangan (negatif untuk berhenti):")
	for {
		fmt.Scan(&num)
		if num < 0 {
			break
		}
		nums = append(nums, num)
	}

	total := Kelipatan4(nums, 0)
	fmt.Printf("Jumlah bilangan kelipatan 4: %d\n", total)
}