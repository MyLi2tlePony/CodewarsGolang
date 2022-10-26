package highest_rank_number_in_an_array

func solution(nums []int) int {
	numberFrequency := make(map[int]int)

	number := 0
	frequency := 0

	for _, num := range nums {
		numberFrequency[num]++

		if numberFrequency[num] > frequency || (numberFrequency[num] == frequency && num > number) {
			number = num
			frequency = numberFrequency[num]
		}
	}

	return number
}
