package service

import (
	"errors"
	"github.com/RahulN12/go-test/model"
	"github.com/RahulN12/go-test/repository"
)

type GraphService struct {
	repo *repository.GraphRepository
}

func NewGraphService(repo *repository.GraphRepository) *GraphService {
	return &GraphService{repo: repo}
}

func (s *GraphService) GetGraphLength() int {
	return s.repo.GetGraphLength()
}

func (s *GraphService) SaveGraph(id string, graph *model.Graph) error {
	for _, edge := range graph.Edges {
		if !graph.ContainsNode(edge[0]) || !graph.ContainsNode(edge[1]) {
			return errors.New("nodes in edge does not exists in graph")
		}
	}
	s.repo.Save(id, graph)
	return nil
}

func (s *GraphService) GetGraph(id string) (*model.Graph, error) {
	graph, exists := s.repo.Get(id)
	if !exists {
		return nil, errors.New("graph not found")
	}
	return graph, nil
}

func (s *GraphService) GetShortestPath(id, start, end string) (*model.ShortestPathResponse, error) {
	graph, exists := s.repo.Get(id)
	if !exists {
		return nil, errors.New("graph not found")
	}

	if !graph.ContainsNode(start) || !graph.ContainsNode(end) {
		return nil, errors.New("nodes does not exists in graph")
	}

	if start == end {
		return &model.ShortestPathResponse{Path: []string{start}}, nil
	}

	var (
		path []string
		err  error
	)
	if path, err = calculateSP(graph, start, end); err != nil {
		return nil, err
	}

	return &model.ShortestPathResponse{Path: path}, nil
}

func (s *GraphService) DeleteGraph(id string) error {
	_, exists := s.repo.Get(id)
	if !exists {
		return errors.New("graph not found")
	}

	s.repo.Delete(id)
	return nil
}

func calculateSP(graph *model.Graph, start, end string) ([]string, error) {
	return nil, errors.New("no path found")
}
