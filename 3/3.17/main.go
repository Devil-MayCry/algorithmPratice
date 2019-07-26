/*
现在有一种新的二叉树节点类型如下：

public class Node{
	public int value;
	public Node left;
	public Node right;
	public Node parent;
	public Node(int data){
		this.value = data;
	}
}
该结构比普通二叉树节点结构多了一个指向父节点的parent指针。
假设有一棵Node类型的节点组成的二叉树，树中每个节点的parent指针都正确地指向自己的父节点，
头节点的parent指向null。只给一个在二叉树中的某个节点node，请实现返回node的后继节点的函数。
在二叉树的中序遍历的序列中。node的下一个节点叫作node的后继节点。
*/

package main

type PNode struct {
	V      int64
	Left   *PNode
	Right  *PNode
	Parent *PNode
}

func main() {
	h := &PNode{}
	searchNext(h)
}

// 寻找下一个节点
func searchNext(n *PNode) *PNode {
	if n == nil {
		return nil
	}
	if n.Right != nil {
		return getLeftMost(n.Right)
	} else {
		p := n.Parent
		for p != nil && p.Left != n {
			n = p
			p = n.Parent
		}
		return p
	}
}

func getLeftMost(n *PNode) *PNode {
	if n == nil {
		return n
	}
	for n.Left != nil {
		n = n.Left
	}
	return n
}
