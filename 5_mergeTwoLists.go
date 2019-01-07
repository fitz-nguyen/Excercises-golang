package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var l3 *ListNode
    var head *ListNode

    for l1 != nil && l2 != nil {
        var temp *ListNode
        if l1.Val <= l2.Val {
            temp = l1
            l1 = l1.Next
        } else {
            temp = l2
            l2 = l2.Next
        }
        if l3 == nil {
            l3 = &ListNode{temp.Val, nil}
            head = l3
        } else {
            newl3 := &ListNode{temp.Val, nil}
            l3.Next = newl3
            l3 = newl3
        }
    }
    return head
}

func (list *ListNode) show() {
    head := list
    fmt.Println("--------")
    for head.Next != nil {
        fmt.Println(head.Val)
        head = head.Next
    }
}



func main() {
    list1 := &ListNode{1, nil}
    head1 := list1
    list1.Next = &ListNode{3, nil}
    list1 = list1.Next
    list1.Next = &ListNode{5, nil}
    list1 = list1.Next
    list1.Next = &ListNode{7, nil}
    list1 = list1.Next

    list2 := &ListNode{2, nil}
    head2 := list2
    list2.Next = &ListNode{4, nil}
    list2 = list2.Next
    list2.Next = &ListNode{6, nil}
    list2 = list2.Next
    list2.Next = &ListNode{8, nil}
    list2 = list2.Next

    result := mergeTwoLists(head1, head2)
    result.show()
}