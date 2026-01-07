package saga_pattern

import (
	"context"
	"sync"
)

type MemoryStore struct {
	mu   sync.Mutex
	data map[string][]Log
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{data: make(map[string][]Log)}
}
func (m *MemoryStore) AppendLog(ctx context.Context, sagaName string, log Log) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[sagaName] = append(m.data[sagaName], log)
	return nil
}
func (m *MemoryStore) GetLogs(ctx context.Context, sagaName string) ([]Log, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	logs := m.data[sagaName]
	copied := make([]Log, len(logs))
	copy(copied, logs)
	return copied, nil
}
func (m *MemoryStore) Clear(ctx context.Context, sagaName string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, sagaName)
	return nil
}
