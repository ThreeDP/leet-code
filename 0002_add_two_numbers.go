/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var newl *ListNode = new(ListNode)
    l := newl
    carrying := 0
    for l1 != nil || l2 != nil {
        num1 := 0
        num2 := 0
        if l1 != nil {
            num1 = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            num2 = l2.Val
            l2 = l2.Next
        }
        num := num1 + num2 + carrying
        carrying = num / 10
        num = num % 10
        l.Val = num
        if l1 != nil || l2 != nil || carrying > 0 { l.Next = new(ListNode) }
        l = l.Next
    }
    if carrying > 0 {
        l.Val = carrying
    }
    return newl
}
