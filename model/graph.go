package model

type Graph struct {
	Nodes []string   `json:"nodes"`
	Edges [][]string `json:"edges"` // unweighted graph as required by problem statement
}

func (g *Graph) ContainsNode(e string) bool {
	for _, a := range g.Nodes {
		if a == e {
			return true
		}
	}
	return false
}

type ShortestPathResponse struct {
	Path []string `json:"path"`
}
