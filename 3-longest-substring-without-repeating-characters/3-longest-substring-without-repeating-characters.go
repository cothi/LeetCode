func lengthOfLongestSubstring(s string) int {
   	a := s
	var ans, tmp int = 0, 0
	var hMap = make(map[rune]int)

	for j, i := range a {
		hMap = make(map[rune]int)
		tmp = 1
		hMap[i] = 1

		for _, k := range a[j+1:] {
			_, exists := hMap[k]
			if !exists {
				tmp++
				hMap[k] = 1
			} else {
				break
			}
		}

		if ans < tmp {
			ans = tmp
		}

	} 
    return ans
}	