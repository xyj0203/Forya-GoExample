package leetcode

//   遍历 树 如果节点相等， 则遍历子节点
// 			  节点不相等， 则继续遍历子节点， 重置指针为头指针直到 head == nil

var h *ListNode

func isSubPath(head *ListNode, root *TreeNode) bool {
	h = head
	return isSubPath1(head, root, nil)
}

// 中间处不想等
func isSubPath1(head *ListNode, root *TreeNode, parent *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	// 左右
	if root.Val == head.Val {
		if parent == nil {
			parent = root
		}
		return isSubPath1(head.Next, root.Left, parent) || isSubPath1(head.Next, root.Right, parent)
	}
	if parent != nil {
		left := parent.Left
		right := parent.Right
		parent = nil
		return isSubPath1(h, left, parent) || isSubPath1(h, right, parent)
	} else {
		// 从第一次匹配后的下一个重新匹配
		return isSubPath1(h, root.Left, parent) || isSubPath1(h, root.Right, parent)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
