func twoSum(nums []int, target int) []int {
    resultValue := make([]int, 2)
 
    for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
            if nums[i]+nums[j] == target {
                resultValue[0] = i
                resultValue[1] = j
                return resultValue
            }
		}
	}
	return resultValue
}