/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func findLast(l *ListNode) *ListNode {
    var prevLast *ListNode
    if l == nil {
        return l
    }
    for l.Next != nil {
        prevLast = l
        l = l.Next
    }
    if prevLast != nil {
        prevLast.Next = nil
    }
    return l
}

func reorderListAux(head, last *ListNode) {
    if head == nil {
        return
    }
    aux := head.Next
    head.Next = last
    last.Next = aux
    reorderListAux(aux, findLast(aux))
}

func reorderList(head *ListNode)  {
    reorderListAux(head, findLast(head))
}
