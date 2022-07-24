# Binary Search Tree
Binary Search Tree is a node-based binary tree data structure which has the following properties:

The left subtree of a node contains only nodes with keys lesser than the node’s key.
The right subtree of a node contains only nodes with keys greater than the node’s key.
The left and right subtree each must also be a binary search tree.

## Functionalities
- Insert value
- Delete value
- Find value in tree
- Find Minimum value in tree
- Find Maximum value in tree
- Print tree in-order
- Print tree pre-order
- Print tree post-order

## Installation

```sh
go get https://github.com/ukutluer/go-tree
go mod tidy
```
After installation you can import below. 
import (
	"github.com/ukutluer/go-tree/tree"
)

## License
MIT
