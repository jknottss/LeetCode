/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    res := &ListNode{}
    sum := 0
    for current := res; l1 != nil || l2 != nil || sum != 0; {
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
            
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }
        current.Next = &ListNode{Val: sum % 10}
        sum /= 10
        current = current.Next
    }
    return res.Next
}
