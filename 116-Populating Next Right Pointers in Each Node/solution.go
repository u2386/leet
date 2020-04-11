package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	p := root
	for p != nil && p.Left != nil {
		t := p
		for t != nil {
			t.Left.Next = t.Right
			if t.Next != nil {
				t.Right.Next = t.Next.Left
			}
			t = t.Next
		}

		p = p.Left
	}
	return root
}

func main() {

}
