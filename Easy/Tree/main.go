package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	numNode := 4
	numberOFUniqueTree := numTrees(numNode)
	fmt.Printf("%v nodes can make %d unique trees", numNode, numberOFUniqueTree)
}
