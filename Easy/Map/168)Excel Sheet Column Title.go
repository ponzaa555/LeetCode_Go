package main

func convertToTitle(columnNumber int) string {
	/*
		For example:
			A -> 1
			B -> 2
			C -> 3
			...
			Z -> 26
			AA -> 27
			AB -> 28
			...
	*/
	mapstring := map[int]string{1: "A", 2: "B", 3: "C", 4: "D", 5: "E", 6: "F", 7: "G", 8: "H", 9: "I", 10: "J", 11: "K", 12: "L", 13: "M", 14: "N", 15: "O", 16: "P", 17: "Q", 18: "R", 19: "S", 20: "T", 21: "U", 22: "V", 23: "W", 24: "X", 25: "Y", 26: "Z"}
	ans := ""
	current := columnNumber
	for current > 26 {
		if current%26 == 0 {
			current = current/26 - 1
			ans = mapstring[26] + ans
		} else {
			ans = mapstring[current%26] + ans
			current = current / 26
		}
	}
	if current != 0 {
		ans = mapstring[current] + ans
	}
	return ans
}
