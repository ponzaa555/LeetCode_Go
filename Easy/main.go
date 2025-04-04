package main

import "fmt"

func main() {
	// columnNumber := 1020
	// ans := convertToTitle(columnNumber)
	// println("Column number", columnNumber, "corresponds to title", ans)

	nums := []int{3, 2, 3}
	ans := majorityElement(nums)
	fmt.Printf("List of number %v have Majornumber is %v ", nums, ans)
}
