package main

func numTrees(n int) int {
	listNumberUniqueNodeCreatedByNode := []int{1, 1, 2} // 0 node 1 node 2 node
	if n <= 2 {
		return listNumberUniqueNodeCreatedByNode[n]
	}
	// right node make ต้องกับ for ไว้เนื่องจากถ้า [1,2,3,4,5,6] ตอนแรกเราจะไม่ที่ 5 โหนดทำไดแค่รอบแรก
	for i := len(listNumberUniqueNodeCreatedByNode); i <= n; i++ {
		uniqueNode := 0
		for j := 1; j <= i; j++ {
			leftNode := j - 1
			rightNode := i - j
			nodeMake := listNumberUniqueNodeCreatedByNode[leftNode] * listNumberUniqueNodeCreatedByNode[rightNode]
			uniqueNode += nodeMake
		}
		listNumberUniqueNodeCreatedByNode = append(listNumberUniqueNodeCreatedByNode, uniqueNode)
	}
	return listNumberUniqueNodeCreatedByNode[len(listNumberUniqueNodeCreatedByNode)-1]
}
