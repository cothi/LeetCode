func twoSum(nums []int, target int) []int {
    var map1 = make(map[int]int)
	for idx, numi := range nums {
		num := target - numi
		if j, ok := map1[num]; ok {
			return []int{j, idx}
		}
		map1[numi] = idx
	}
	return []int{0, 0}
}