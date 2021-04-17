package graph

import (
	"fmt"
	"sync"
)

//节点
type Node struct {
	value int
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.value)
}

type Graph struct {
	nodes []*Node          //节点
	edges map[Node][]*Node //边, 邻接表
	sync.RWMutex
}

//添加结点
func (g *Graph) AddNode(n *Node) {
	g.Lock()
	g.nodes = append(g.nodes, n)
	g.Unlock()
}

//添加边
func (g *Graph) AddEdge(n1, n2 *Node) {
	g.Lock()
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
	g.Unlock()
}

func (g *Graph) String() {
	g.Lock()
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].String() + " -> "
		near := g.edges[*g.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
	g.Unlock()
}
