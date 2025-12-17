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
	tmp := head
	count := 1
	for tmp != nil {
		if count%2 == 0 {
			middle = middle.Next
		}
		tmp = tmp.Next
		count++
	}
	return middle
}
