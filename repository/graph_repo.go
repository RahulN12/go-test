package repository

import (
	"github.com/RahulN12/go-test/model"
	"sync"
)

type GraphRepository struct {
	mu     sync.RWMutex
	graphs map[string]*model.Graph
}

func (g *GraphRepository) GetGraphLength() int {
	return len(g.graphs)
}

func NewGraphRepository() *GraphRepository {
	return &GraphRepository{
		mu:     sync.RWMutex{},
		graphs: make(map[string]*model.Graph),
	}
}

func (r *GraphRepository) Save(id string, graph *model.Graph) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.graphs[id] = graph
}

func (r *GraphRepository) Get(id string) (*model.Graph, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	graph, exists := r.graphs[id]
	return graph, exists
}

func (r *GraphRepository) Delete(id string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.graphs, id)
}
