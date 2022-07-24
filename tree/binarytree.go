package tree

import (
	"errors"
	"fmt"
)

type BinaryTreeNode struct {
	data  int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func (t *BinaryTreeNode) Insert(data int) error {
	if t == nil {
		return errors.New("Tree is nil")
	}
	if t.data == data {
		return errors.New("Value already exist")
	}
	if data < t.data {
		if t.left == nil {
			t.left = &BinaryTreeNode{data: data}
			return nil
		}
		t.left.Insert(data)
	}
	if data > t.data {
		if t.right == nil {
			t.right = &BinaryTreeNode{data: data}
		}
		t.right.Insert(data)
	}
	return nil

}

func (t *BinaryTreeNode) Delete(data int) *BinaryTreeNode {
	if t == nil {
		return nil
	}
	if data < t.data {
		t.left = t.left.Delete(data)
		return t
	}
	if data > t.data {
		t.right = t.right.Delete(data)
		return t
	}
	if data == t.data {
		//there is no child in current node
		if t.left == nil && t.right == nil {
			t = nil
			return t
		}
		//there is only right child in current node
		if t.left == nil && t.right != nil {
			t = t.right
			return t
		}
		//there is only left child in current node
		if t.left != nil && t.right == nil {
			t = t.left
			return t
		}

		//current node has left and right childs,
		//Solution 1: we can find smallest value on right child, replace to current node then remove smallest item in right child.
		//Solution 2: we can find biggest value on left child, replace to current node then remove biggest item in left child.

		//We prefer solution 1
		smallestNodeOnRight := t.right
		for {
			//find smallest value on the right side
			if smallestNodeOnRight != nil && smallestNodeOnRight.left != nil {
				smallestNodeOnRight = smallestNodeOnRight.left
			} else {
				break
			}
		}
		t.data = smallestNodeOnRight.data
		t.right = t.right.Delete(t.data)
		return t
	}

	return nil

}

func (t *BinaryTreeNode) Find(data int) (BinaryTreeNode, bool) {
	if t == nil {
		return BinaryTreeNode{}, false
	}
	if t.data == data {
		return *t, true
	}
	if data < t.data {
		return t.left.Find(data)
	}
	if data > t.data {
		return t.right.Find(data)
	}
	return BinaryTreeNode{}, false
}

func (t *BinaryTreeNode) FindMinimum() int {
	if t == nil {
		return -1
	}
	if t.left == nil {
		return t.data
	}
	return t.left.FindMinimum()
}

func (t *BinaryTreeNode) FindMaximum() int {
	if t == nil {
		return -1
	}
	if t.right == nil {
		return t.data
	}
	return t.right.FindMaximum()
}

func (t *BinaryTreeNode) PrintTreeInOrder() {
	if t == nil {
		return
	}
	t.left.PrintTreeInOrder()
	fmt.Print(t.data, " ")
	t.right.PrintTreeInOrder()

}

func (t *BinaryTreeNode) PrintTreePreOrder() {
	if t == nil {
		return
	}
	fmt.Print(t.data, " ")
	t.left.PrintTreePreOrder()
	t.right.PrintTreePreOrder()

}

func (t *BinaryTreeNode) PrintTreePostOrder() {
	if t == nil {
		return
	}
	t.left.PrintTreePostOrder()
	t.right.PrintTreePostOrder()
	fmt.Print(t.data, " ")
}

func NewBinaryTree(data int) *BinaryTreeNode {
	btn := &BinaryTreeNode{data: data}
	return btn
}
