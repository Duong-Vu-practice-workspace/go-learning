package saga_pattern

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
)

func TestAllStepsSucceed(t *testing.T) {
	store := NewMemoryStore()
	coord := NewCoordinator(store)
	s := NewSaga("all-success")
	var calls []string

	s.AddStep(Step{
		Name: "step1",
		Action: func(ctx context.Context) (interface{}, error) {
			calls = append(calls, "a1")
			return map[string]string{"id": "1"}, nil
		},
		Compensate: func(ctx context.Context, out json.RawMessage) error {
			calls = append(calls, "c1")
			return nil
		},
	})
	s.AddStep(Step{
		Name: "step2",
		Action: func(ctx context.Context) (interface{}, error) {
			calls = append(calls, "a2")
			return "ok", nil
		},
		Compensate: func(ctx context.Context, out json.RawMessage) error {
			calls = append(calls, "c2")
			return nil
		},
	})
	res := coord.Execute(context.Background(), s)
	if !res.Success {
		t.Fatalf("expected success, got error: %v", res.Err)
	}
	if len(res.Completed) != 2 {
		t.Fatalf("expected 2 completed steps, got %d", len(res.Completed))
	}
}

func TestCompensateCalledWhenError(t *testing.T) {
	store := NewMemoryStore()
	coord := NewCoordinator(store)
	s := NewSaga("compensate")

	var compensated bool
	s.AddStep(Step{
		Name: "step1",
		Action: func(ctx context.Context) (interface{}, error) {
			return map[string]int{"id": 99}, nil
		},
		Compensate: func(ctx context.Context, out json.RawMessage) error {
			var v map[string]int
			_ = json.Unmarshal(out, &v)
			if v["id"] == 99 {
				compensated = true
			}
			return nil
		},
	})
	s.AddStep(Step{
		Name: "step2",
		Action: func(ctx context.Context) (interface{}, error) {
			return nil, errors.New("boom")
		},
	})
	res := coord.Execute(context.Background(), s)
	if res.Success {
		t.Fatalf("expected failure, got success")
	}
	if !compensated {
		t.Fatalf("expected compensation to be called")
	}
}

func TestCompensationErrorIsReturned(t *testing.T) {
	store := NewMemoryStore()
	coord := NewCoordinator(store)
	s := NewSaga("comp-err")

	s.AddStep(Step{
		Name: "step1",
		Action: func(ctx context.Context) (interface{}, error) {
			return "foo", nil
		},
		Compensate: func(ctx context.Context, out json.RawMessage) error {
			return errors.New("comp-regret")
		},
	})
	s.AddStep(Step{
		Name: "step2",
		Action: func(ctx context.Context) (interface{}, error) {
			return nil, errors.New("boom")
		},
	})
	res := coord.Execute(context.Background(), s)
	if res.Success {
		t.Fatalf("expected failure, got success")
	}
	if res.CompensateError == nil || res.CompensateError.Error() != "comp-regret" {
		t.Fatalf("expected compensation error, got %v", res.CompensateError)
	}
}
