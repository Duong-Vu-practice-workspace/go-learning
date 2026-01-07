package saga_pattern

import (
	"context"
	"encoding/json"
)

type Log struct {
	StepName string          `json:"step_name"`
	Output   json.RawMessage `json:"output"`
}
type Store interface {
	AppendLog(ctx context.Context, sagaName string, log Log) error
	GetLogs(ctx context.Context, sagaName string) ([]Log, error)
	Clear(ctx context.Context, sagaName string) error
}
