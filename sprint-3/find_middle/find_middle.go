package find_middle

type Node struct {
	Val  int
	Next *Node
}

func FindMiddle(head *Node) *Node {
	if head == nil {
		return nil
	}

	middle := head
	node := head.Next

	for node != nil {
		middle = middle.Next

		if node.Next == nil || node.Next.Next == nil {
			break
		}

		node = node.Next.Next
	}
	
	return middle
}
