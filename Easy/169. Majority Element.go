package main

func majorityElement(nums []int) int {
	majorityElement := 0
	mapCount := make(map[int]int)

	if len(nums) == 1 {
		return nums[0]
	}
	for _, num := range nums {
		_, exist := mapCount[num]

		if exist {
			mapCount[num] += 1
			if mapCount[num] > len(nums)/2 {
				majorityElement = num
				break
			}
		} else {
			mapCount[num] = 1
		}
	}
	return majorityElement
}
