package repository

import "github.com/vanWezel/request-count-server/internal/model"

type Mock struct {
	Interface
}

func NewMock() *Mock {
	return &Mock{}
}

func (m *Mock) Ping() error {
	return nil
}

func (m *Mock) Increment() (*model.Stats, error) {
	return &model.Stats{
		Instance: 1,
		Total:    1,
	}, nil
}

func (m *Mock) GetStats() (*model.Stats, error) {
	return &model.Stats{
		Instance: 1,
		Total:    1,
	}, nil
}
