package mergelists

type Node struct {
	Val  int
	Next *Node
}

func compareNode(n1 *Node, n2 *Node) (*Node, *Node) {
	if n1.Val < n2.Val {
		return n1, n2
	}
	return n2, n1
}

func MergeLists(list1 *Node, list2 *Node) *Node {
	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	head, list := compareNode(list1, list2)
	tmp_head := head.Next
	tmp_list := list.Next

	prev_head := head
	prev_list := list

	for tmp_head != nil {
		if prev_list != nil && tmp_head.Val >= prev_list.Val {
			prev_head.Next = prev_list
			prev_list.Next = tmp_head
		}

		if prev_list != nil && tmp_head.Val < prev_list.Val {
			tmp_head.Next = prev_list
			prev_head = tmp_head
		}

		if tmp_list != nil && tmp_head.Next == nil {
			if tmp_head.Val < tmp_list.Val {
				tmp_head.Next = tmp_list
				tmp_list = tmp_head
			}
		}

		prev_head = tmp_head
		tmp_head = tmp_head.Next
		prev_list = tmp_list
		if tmp_list != nil {
			tmp_list = tmp_list.Next
		}
	}
	return head
}
