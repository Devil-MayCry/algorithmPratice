package tree

// Node 二叉树节点
type Node struct {
	V     int64
	Left  *Node
	Right *Node
}

// MakeExampleTree 产生一颗特定的树
func MakeExampleTree() *Node {
	a := &Node{V: 1}
	b := &Node{V: 2}
	c := &Node{V: 3}
	d := &Node{V: 4}
	e := &Node{V: 5}
	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	return a
}

// PrintlnTree 打印整颗树
func (t *Node) PrintlnTree() {
	travesTree(t)
}

// 中序
func travesTree(t *Node) {
	if t == nil {
		return
	}
	println(t.V)
	if t.Left != nil {
		travesTree(t.Left)
	}
	if t.Right != nil {
		travesTree(t.Right)
	}
}
