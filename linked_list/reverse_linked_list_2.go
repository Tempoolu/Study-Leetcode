package main

// 中等难度，大致可以做出来，边界条件无法完成
// https://leetcode.cn/problems/reverse-linked-list-ii/

// 技巧：使用穿针引线法，虚拟头

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	// border condition
	if head == nil || head.Next == nil || left == right {
		return head
	}

	// find start node
	startNode := head
	var preNode *ListNode
	for i := 1; i < left; i += 1 {
		preNode = startNode
		startNode = startNode.Next
	}

	//reverse scoped chain
	current := startNode
	next := startNode.Next
	for i := left; i < right; i += 1 {
		tmp := next.Next
		next.Next = current
		current = next
		next = tmp
	}

	// reconnect to the chain
	if preNode != nil {
		preNode.Next = current
	}
	startNode.Next = next

	// find the real head
	if left == 1 {
		return current
	} else {
		return head
	}

}

// func reverseBetween(head *ListNode, left int, right int) *ListNode {

//     if left == right {
//         return head
//     }
//     cur := head
//     var seenL bool = false
//     var pre, a, b, c, d *ListNode
//     var count int = 1
//     ans := &ListNode{Val:-1, Next:head}

//     for cur != nil {

//         if left == 1 {
//             a, b = cur, cur
//             seenL = true
//         }

//         if seenL == false {
//             pre = cur
//             cur = cur.Next
//             count++

//             if count == left {
//                 a = pre
//                 b = cur
//                 seenL = true
//             }
//             continue
//         }

//         if seenL == true {

//             next := cur.Next
//             cur.Next = pre
//             pre = cur
//             cur = next
//             count++

//             if count == right {
//                 c = cur
//                 d = cur.Next
//                 c.Next = pre
//                 break
//             }
//         }
//     }
// 穿针引线
//     if left != 1 {
//         a.Next = c
//         b.Next = d
//     }

//     if left == 1 {
//         return c
//     }

//     return ans.Next
// }
