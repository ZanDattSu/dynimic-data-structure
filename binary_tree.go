package main

import (
	"dynimic-data-structure/figures"
	"fmt"
	"math"
)

type TreeNode struct {
	value figures.Shape
	left  *TreeNode
	right *TreeNode
}

type BinaryTree struct {
	root *TreeNode
}

func (t *BinaryTree) Insert(shape figures.Shape) {
	t.root = insertTreeNode(t.root, shape)
}

func insertTreeNode(node *TreeNode, shape figures.Shape) *TreeNode {
	if node == nil {
		return &TreeNode{value: shape}
	}

	if shape.Area() < node.value.Area() {
		node.left = insertTreeNode(node.left, shape)
	} else {
		node.right = insertTreeNode(node.right, shape)
	}

	return node
}

func (t *BinaryTree) Search(area float64) figures.Shape {
	return searchNode(t.root, area)
}

func searchNode(node *TreeNode, area float64) figures.Shape {
	if node == nil {
		return nil
	}

	nodeArea := node.value.Area()

	if math.Abs(nodeArea-area) < 1e-9 {
		return node.value
	}

	if area < nodeArea {
		return searchNode(node.left, area)
	}

	return searchNode(node.right, area)
}

func (t *BinaryTree) FindMin() figures.Shape {
	if t.root == nil {
		return nil
	}
	return findMin(t.root).value
}

func findMin(node *TreeNode) *TreeNode {
	if node.left == nil {
		return node
	}
	return findMin(node.left)
}

func (t *BinaryTree) FindMax() figures.Shape {
	if t.root == nil {
		return nil
	}
	return findMax(t.root).value
}

func findMax(node *TreeNode) *TreeNode {
	if node.right == nil {
		return node
	}
	return findMax(node.right)
}

func (t *BinaryTree) Height() int {
	return height(t.root)
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftHeight := height(node.left)
	rightHeight := height(node.right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func (t *BinaryTree) Size() int {
	return size(t.root)
}

func size(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + size(node.left) + size(node.right)
}

func (t *BinaryTree) FindInRange(min, max float64) []figures.Shape {
	var result []figures.Shape
	findInRange(t.root, min, max, &result)
	return result
}

func findInRange(node *TreeNode, min, max float64, result *[]figures.Shape) {
	if node == nil {
		return
	}

	nodeArea := node.value.Area()

	if nodeArea > min {
		findInRange(node.left, min, max, result)
	}

	if nodeArea >= min && nodeArea <= max {
		*result = append(*result, node.value)
	}

	if nodeArea < max {
		findInRange(node.right, min, max, result)
	}
}

func (t *BinaryTree) ToSlice() []figures.Shape {
	var result []figures.Shape
	toSlice(t.root, &result)
	return result
}

func toSlice(node *TreeNode, result *[]figures.Shape) {
	if node == nil {
		return
	}
	toSlice(node.left, result)
	*result = append(*result, node.value)
	toSlice(node.right, result)
}

func (t *BinaryTree) PrintInOrder() {
	printInOrder(t.root)
}

func printInOrder(node *TreeNode) {
	if node == nil {
		return
	}

	printInOrder(node.left)
	fmt.Printf("%s area = %.2f\n", node.value.GetName(), node.value.Area())
	printInOrder(node.right)
}
