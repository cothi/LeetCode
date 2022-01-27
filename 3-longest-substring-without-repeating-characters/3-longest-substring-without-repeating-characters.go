func lengthOfLongestSubstring(s string) int {

	var left, right, res int
	ch := make(map[byte]int)

	for right < len(s) {

		// 존재하는지 확인!
		index, exist := ch[s[right]]

		// 존재하면, 존재하는 index 인덱스에서 right 다시 증명 시작
        // left가 존재하는 다음 자리에서 다시 증명을 시작하는 이유는
        // 이미 right까지 중복한 숫자가 없다는 것을 증명했기 떄문
		if exist && index >= left {
			left = index + 1
		}

		if res < right-left+1 {
			res = right - left + 1
		}

		ch[s[right]] = right

		right += 1
	}
    
    return res
}	