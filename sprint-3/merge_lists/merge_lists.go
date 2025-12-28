package mergelists

type Node struct {
	Val  int
	Next *Node
}

func MergeLists(list1 *Node, list2 *Node) *Node {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var mergeList *Node

	if list1.Val < list2.Val {
		mergeList = list1
		list1 = list1.Next

	} else {
		mergeList = list2
		list2 = list2.Next
	}

	nodeMerdgeList := mergeList

	for list1 != nil || list2 != nil {

		if list1 == nil {
			nodeMerdgeList.Next = list2
			break
		}

		if list2 == nil {
			nodeMerdgeList.Next = list1
			break
			
		}

		if list1.Val < list2.Val {
			nodeMerdgeList.Next = list1
			list1 = list1.Next
			nodeMerdgeList = nodeMerdgeList.Next

		} else {
			nodeMerdgeList.Next = list2
			list2 = list2.Next
			nodeMerdgeList = nodeMerdgeList.Next

		}

	}

	return mergeList
}
