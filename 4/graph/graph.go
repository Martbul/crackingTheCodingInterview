package graph

type Node struct {
	Value    string
	Adjacent []*Node
	Visited  bool
}

func CreateGraph() *Node {
	// Create nodes
	a := &Node{Value: "a"}
	b := &Node{Value: "b"}
	c := &Node{Value: "c"}
	d := &Node{Value: "d"}
	e := &Node{Value: "e"}
	f := &Node{Value: "f"}
	g := &Node{Value: "g"}
	h := &Node{Value: "h"}

	// Set up adjacency list
	a.Adjacent = []*Node{b, c}
	b.Adjacent = []*Node{d, e}
	c.Adjacent = []*Node{f}
	e.Adjacent = []*Node{g}
	f.Adjacent = []*Node{h}

	return a // Root node is 'a'
}

func CreateDirectedGraph() (*Node, *Node) {
	// Create nodes
	a := &Node{Value: "a"}
	b := &Node{Value: "b"}
	c := &Node{Value: "c"}
	d := &Node{Value: "d"}
	e := &Node{Value: "e"}
	f := &Node{Value: "f"}

	// Set up adjacency list (directed edges)
	a.Adjacent = []*Node{b, d}
	b.Adjacent = []*Node{c}
	c.Adjacent = []*Node{e}
	d.Adjacent = []*Node{e}
	e.Adjacent = []*Node{f}

	return b, f
}
