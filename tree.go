package main

import "fmt"

// Node represents a single node in the tree
type Node struct {
	value  int
	left   *Node
	right  *Node
	object map[string]any
}

// BST represents the binary search tree
type BST struct {
	root        *Node
	resultOrder *[]map[string]any
}

func CreateBST(array []map[string]any, key string) *BST {
	bst := &BST{}
	for _, arr := range array {
		bst.Insert(arr[key].(int), arr)
	}
	return bst
}

// Insert inserts a new value into the BST
func (bst *BST) Insert(value int, object map[string]any) {
	newNode := &Node{value: value, object: object}
	if bst.root == nil {
		bst.root = newNode
		bst.resultOrder = &[]map[string]any{}
	} else {
		insertNode(bst.root, newNode)
	}
}

func insertNode(node, newNode *Node) {
	if newNode.value < node.value {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

// Search searches for a value in the BST
func (bst *BST) Search(value int) map[string]any {
	return searchNode(bst.root, value)
}

func searchNode(node *Node, value int) map[string]any {
	if node == nil {
		return nil
	}
	if value < node.value {
		return searchNode(node.left, value)
	} else if value > node.value {
		return searchNode(node.right, value)
	} else {
		return node.object
	}
}

// InOrderTraversal traverses the tree in-order
func (bst *BST) InOrderTraversal() []map[string]any {
	inOrder(bst.root, bst.resultOrder)
	return *bst.resultOrder
}

func inOrder(node *Node, resultOrder *[]map[string]any) {
	if node != nil {
		inOrder(node.left, resultOrder)
		*resultOrder = append(*resultOrder, node.object)
		inOrder(node.right, resultOrder)
	}
}

// PreOrderTraversal traverses the tree pre-order
func (bst *BST) PreOrderTraversal() {
	preOrder(bst.root)
}

func preOrder(node *Node) {
	if node != nil {
		fmt.Printf("%d ", node.value)
		preOrder(node.left)
		preOrder(node.right)
	}
}

// PostOrderTraversal traverses the tree post-order
func (bst *BST) PostOrderTraversal() {
	postOrder(bst.root)
}

func postOrder(node *Node) {
	if node != nil {
		postOrder(node.left)
		postOrder(node.right)
		fmt.Printf("%d ", node.value)
	}
}
