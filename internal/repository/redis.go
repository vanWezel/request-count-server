package repository

import (
	"fmt"

	"github.com/vanWezel/request-count-server/internal/model"
	"github.com/vanWezel/request-count-server/internal/redis"
)

type Redis struct {
	Interface
	client *redis.Client
	instance string
}

func NewRedis(host string, db int, instance string) *Redis {
	return &Redis{
		client:   redis.New(host, db),
		instance: instance,
	}
}

func (r *Redis) Ping() error {
	return r.client.Ping()
}

func (r *Redis) FlushDb() error {
	return r.client.FlushDb()
}

func (r *Redis) Increment() (*model.Stats, error) {
	instance, err := r.client.Increment(fmt.Sprintf("instance:%v", r.instance))
	if err != nil {
		return nil, err
	}

	total, err := r.client.Increment("total")
	if err != nil {
		return nil, err
	}

	return &model.Stats{
		Instance: instance,
		Total:    total,
	}, nil
}

func (r *Redis) GetStats() (*model.Stats, error) {
	instance, err := r.client.GetInt(fmt.Sprintf("instance:%v", r.instance))
	if err != nil {
		return nil, err
	}

	total, err := r.client.GetInt("total")
	if err != nil {
		return nil, err
	}

	return &model.Stats{
		Instance: instance,
		Total:    total,
	}, nil
}
