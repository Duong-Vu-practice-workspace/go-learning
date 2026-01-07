package saga_pattern

import (
	"context"
	"encoding/json"
)

type Result struct {
	Success         bool
	Err             error
	Completed       []string
	CompensateError error
}
type Coordinator struct {
	store Store
}

func NewCoordinator(store Store) *Coordinator {
	return &Coordinator{store: store}
}
func (c *Coordinator) Execute(ctx context.Context, s *Saga) Result {
	var completed []string
	for _, step := range s.Steps() {
		out, err := step.Action(ctx)
		if err != nil {
			// fetch previous logs and compensate
			prev, _ := c.store.GetLogs(ctx, s.Name)
			compErr := c.compensate(ctx, s, prev)
			return Result{
				Success:         false,
				Err:             err,
				Completed:       completed,
				CompensateError: compErr,
			}
		}
		raw, _ := json.Marshal(out)
		_ = c.store.AppendLog(ctx, s.Name, Log{StepName: step.Name, Output: raw})
		completed = append(completed, step.Name)
	}
	return Result{Success: true, Completed: completed}
}

func (c *Coordinator) compensate(ctx context.Context, s *Saga, logs []Log) error {
	var firstErr error
	// map steps by name
	stepMap := map[string]Step{}
	for _, st := range s.Steps() {
		stepMap[st.Name] = st
	}
	for i := len(logs) - 1; i >= 0; i-- {
		l := logs[i]
		st, ok := stepMap[l.StepName]
		if !ok || st.Compensate == nil {
			continue
		}
		if err := st.Compensate(ctx, l.Output); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}
