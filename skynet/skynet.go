package skynet

import (
	"errors"
	"fmt"
)

type node int

type link struct {
	n1, n2 node
}

func (l link) getLinked(n node) (node, error) {
	if l.n1 == n {
		return l.n2, nil
	}

	if l.n2 == n {
		return l.n1, nil
	}

	return 0, errors.New("Invalid link")
}

//func (l link) isSame(targetLink link) bool {
//	return (targetLink.n1 == l.n1 && targetLink.n2 == l.n2) || (targetLink.n2 == l.n1 && targetLink.n1 == l.n2)
//}

type graph struct {
	links []link
}

func (g *graph) getLinkedNodes(n node, excludedNodes []node) (result []node) {
outer:
	for _, l := range g.links {
		if linkedNode, err := l.getLinked(n); err == nil {
			// exclude linked nodes
			for _, en := range excludedNodes {
				if en == linkedNode {
					continue outer
				}
			}

			result = append(result, linkedNode)
		}
	}

	return
}

type treeNode struct {
	payload        node
	parentTreeNode *treeNode
}

//func (g *graph) unlink(targetLink link) {
//	for i, l := range g.links {
//		if (l.isSame(targetLink)) {
//			g.links = append(g.links[:i])
//		}
//	}
//}

func solve() {
	var skynetNode node = 0
	var virusNode node = 3
	g := graph{
		[]link{
			{0, 1},
			{0, 2},
			{1, 3},
			{2, 3},
		},
	}

	visitedNodes := []node{virusNode}
	queue := make(chan treeNode, 500)
	rootTreeNode := treeNode{virusNode, nil}
	queue <- rootTreeNode

	var foundDstTreeNode treeNode
	for currentTreeNode := range queue {
		if currentTreeNode.payload == skynetNode {
			foundDstTreeNode = currentTreeNode
			close(queue)
			break
		}

		linkedNotVisitedNodes := g.getLinkedNodes(currentTreeNode.payload, visitedNodes)
		for _, l := range linkedNotVisitedNodes {
			s := currentTreeNode
			queue <- treeNode{l, &s}
			visitedNodes = append(visitedNodes, l)
		}
	}

	fmt.Printf("%#v", foundDstTreeNode)
}
