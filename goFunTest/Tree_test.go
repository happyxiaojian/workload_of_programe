package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Tree struct {
	Root *TreeNode
	Len int
}

func NewTree()*Tree{
	return &Tree{
		Root: new(TreeNode),
		Len:  0,
	}
}


