package main

import "fmt"

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func createSampleTree () *Tree {

	root := Tree{4, nil, nil}
	root.Left = &Tree{2, nil , nil}
	root.Right = &Tree{8, nil , nil}

	root.Left.Left = &Tree{1, nil, nil}
	root.Left.Right = &Tree{3, nil, nil}

	root.Right.Left = &Tree{5, nil, nil}
	root.Right.Right = &Tree{13, nil, nil}

	return &root
}

func (t *Tree) Inorder( root *Tree)  {
	if root != nil {
		t.Inorder(root.Left)
		Visit(root)
		t.Inorder(root.Right)
	}
}

func (t *Tree) PreOrder( root *Tree)  {
	if root != nil {
		Visit(root)
		t.PreOrder(root.Left)
		t.PreOrder(root.Right)
	}
}

func (t *Tree) PostOrder( root *Tree)  {
	if root != nil {
		t.PostOrder(root.Left)
		t.PostOrder(root.Right)
		Visit(root)
	}
}

func Visit(tree *Tree) {
	fmt.Print(tree.Value, " ")
}

func main() {
	t := createSampleTree()
	t.Inorder(t)
	fmt.Println()
	t.PreOrder(t)
}
