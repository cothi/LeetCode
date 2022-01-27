func lengthOfLongestSubstring(s string) int {
	a := s
	var ans, tmp int = 0, 0

	for j, i := range a {
		var hMap = make(map[rune]int)
		tmp = 1
		hMap[i] = 1
		for _, k := range a[j+1:] {
			_, exists := hMap[k]
			if !exists {
				tmp += 1
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