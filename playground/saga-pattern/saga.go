package saga_pattern

import (
	"context"
	"encoding/json"
	"fmt"
)

type Step struct {
	Name       string
	Action     func(context.Context) (interface{}, error)
	Compensate func(context.Context, json.RawMessage) error
}
type Saga struct {
	Name  string
	steps []Step
}

func NewSaga(name string) *Saga {
	return &Saga{
		Name: name,
	}
}
func (s *Saga) AddStep(step Step) {
	if step.Name == "" {
		fmt.Errorf("step name must not be empty")
	}
	s.steps = append(s.steps, step)
}
func (s *Saga) Steps() []Step {
	return s.steps
}
