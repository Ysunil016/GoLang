package main

import "fmt"

func main() {
	val := [][]int{{1, 0}, {2, 0}}
	fmt.Println(canFinish(3, val))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses == 0 || len(prerequisites) == 0 {
		return true
	}

	// Making Structure First
	type Node struct {
		inBound   int
		neighbour []int
	}
	NodeCollection := make([]Node, numCourses)
	for i := 0; i < numCourses; i++ {
		NodeCollection[i] = Node{0, []int{}}
	}
	// Populating Nodes W.R.T Neighbours
	for i := 0; i < len(prerequisites); i++ {
		NodeCollection[prerequisites[i][1]].inBound++
		NodeCollection[prerequisites[i][0]].neighbour = append(NodeCollection[prerequisites[i][0]].neighbour, prerequisites[i][1])
	}
	// Adding All Nodes with Zero InBound, Since These Nodes are Not Dependency for Any Node
	queueCollection := []Node{}
	CrossCount := 0
	for i := 0; i < numCourses; i++ {
		if NodeCollection[i].inBound == 0 {
			queueCollection = append(queueCollection, NodeCollection[i])
			CrossCount++
		}
	}
	// Now Looping for All Nodes with Zero Inbound
	for len(queueCollection) != 0 {
		currentNode := queueCollection[0]
		queueCollection = queueCollection[1:] // Shrinking Queue Collections
		for i := 0; i < len(currentNode.neighbour); i++ {
			NodeCollection[currentNode.neighbour[i]].inBound--
			if NodeCollection[currentNode.neighbour[i]].inBound == 0 {
				CrossCount++
				queueCollection = append(queueCollection, NodeCollection[currentNode.neighbour[i]])
			}
		}
	}

	return CrossCount == numCourses
}
