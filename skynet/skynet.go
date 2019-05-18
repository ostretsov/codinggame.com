package skynet

import (
	"errors"
	"fmt"
	"io"
	"log"
	"sort"
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

func (l link) isEqual(targetLink *link) bool {
	return (targetLink.n1 == l.n1 && targetLink.n2 == l.n2) || (targetLink.n2 == l.n1 && targetLink.n1 == l.n2)
}

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

func (t treeNode) getDepth() (depth int) {
	tn := t
	for depth = 0; tn.parentTreeNode != nil; depth++ {
		tn = *tn.parentTreeNode
	}
	return
}

func (t treeNode) getRootLink() (*link, error) {
	if t.parentTreeNode == nil {
		return nil, errors.New("Root treeNode does not have a parent!")
	}

	currentTreeNode := t
	parentTreeNode := currentTreeNode.parentTreeNode
	for parentTreeNode.parentTreeNode != nil {
		currentTreeNode = *parentTreeNode
		parentTreeNode = currentTreeNode.parentTreeNode
	}

	return &link{parentTreeNode.payload, currentTreeNode.payload}, nil
}

func (g *graph) unlink(targetLink *link) {
	for i, l := range g.links {
		if l.isEqual(targetLink) {
			g.links = append(g.links[:i], g.links[i+1:]...)
			break
		}
	}
}

func solve(in io.Reader, out io.Writer) {
	var nodeCount, linkCount, gwCount int
	fmt.Fscanf(in, "%d %d %d\n", &nodeCount, &linkCount, &gwCount)
	links := []link{}
	for i := 0; i < linkCount; i++ {
		var n1, n2 node
		fmt.Fscanf(in, "%d %d\n", &n1, &n2)
		links = append(links, link{n1, n2})
	}
	gateways := []node{}
	for i := 0; i < gwCount; i++ {
		var gw node
		fmt.Fscanf(in, "%d\n", &gw)
		gateways = append(gateways, gw)
	}

	g := graph{links}

	var skynetNode node
	for {
		_, err := fmt.Fscanf(in, "%d\n", &skynetNode)
		if err != nil {
			break
		}

		closestTreeNodes := []treeNode{}

		for gwi, gw := range gateways {
			// remove an isolated gateway
			linkedWithGwNodes := g.getLinkedNodes(gw, []node{})
			if len(linkedWithGwNodes) == 0 {
				gateways = append(gateways[:gwi], gateways[gwi+1:]...)
				continue
			}

			virusNode := gw
			rootTreeNode := treeNode{virusNode, nil}
			visitedNodes := []node{virusNode}
			queue := make(chan treeNode, 500)
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

			closestTreeNodes = append(closestTreeNodes, foundDstTreeNode)
		}

		sort.Slice(closestTreeNodes, func(i, j int) bool {
			return closestTreeNodes[i].getDepth() < closestTreeNodes[j].getDepth()
		})

		closestTreeNode := closestTreeNodes[0]
		rootLink, err := closestTreeNode.getRootLink()
		if err != nil {
			log.Fatal(err)
		}

		g.unlink(rootLink)

		fmt.Fprintf(out, "%d %d\n", rootLink.n1, rootLink.n2)
	}
}
