package containers

import (
	"github.com/Notserc/go-pixel/internal/pkg/ecs"
)

func NewStack() *Stack {
	return &Stack{}
}

type Stack struct {
	nodes []*ecs.Entity
	count int
}

func (s *Stack) Push(n *ecs.Entity) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *Stack) Pop() *ecs.Entity {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}
