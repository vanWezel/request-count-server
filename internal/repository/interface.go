package repository

import "github.com/vanWezel/request-count-server/internal/model"

type Interface interface {
	Ping() error
	Increment() (*model.Stats, error)
	GetStats() (*model.Stats, error)
}
