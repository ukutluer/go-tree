package main

import (
	"fmt"

	"github.com/ukutluer/go-tree/tree"
)

func main() {

	root := tree.NewBinaryTree(5)
	root.Insert(3)
	root.Insert(1)
	root.Insert(4)

	root.Insert(80)
	root.Insert(100)
	root.Insert(110)
	root.Insert(90)
	root.Insert(95)
	root.Insert(96)

	fmt.Println("max item in tree: ", root.FindMaximum())
	fmt.Println("min item in tree: ", root.FindMinimum())

	_, hasValue := root.Find(100)
	if hasValue {
		fmt.Println("Item founded in tree")
	}
	fmt.Println("\n\nTree is printing in-order\n")
	root.PrintTreeInOrder()

	fmt.Println("\n\nTree is printing pre-order\n")
	root.PrintTreePreOrder()

	fmt.Println("\nTree is printing post-order\n")
	root.PrintTreePostOrder()

	root.Delete(100)
	fmt.Println("\n\nTree is printing in-order\n")
	root.PrintTreeInOrder()

}
