/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ansNode := &ListNode{}
	calNode := ansNode
	var sumNum int

	for l1 != nil || l2 != nil {
        if l1 != nil {
			sumNum += l1.Val
            l1 = l1.Next
		}
		if l2 != nil {
			sumNum += l2.Val
            l2 = l2.Next
		}


		calNode.Next = &ListNode{
			Val:  sumNum % 10,
			Next: nil,
		}
		calNode = calNode.Next 
        
		sumNum /= 10
	}
    
    if sumNum > 0 {
        calNode.Next = &ListNode{sumNum, nil}
    }
	return ansNode.Next
}