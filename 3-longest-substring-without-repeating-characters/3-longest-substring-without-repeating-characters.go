func lengthOfLongestSubstring(s string) int {
	a := s

	var left, right, res int
	ch := make(map[byte]int)

	for right < len(a) {

		// 존재하는지 확인!
		index, exist := ch[a[right]]

		// 존재하면, 존재하는 index 인덱스에서 right 다시 증명 시작
		if exist && index >= left {
			left = index + 1
		}

		if res < right-left+1 {
			res = right - left + 1
		}

		ch[a[right]] = right

		right += 1
	}
    
    return res
}	