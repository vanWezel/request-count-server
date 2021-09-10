package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vanWezel/request-count-server/pkg/helper"
)

func TestRedis_IncrementSuccess(t *testing.T) {
	host := helper.Getenv("TEST_REDIS_HOST", "127.0.0.1:6379")
	db := helper.GetenvToi("TEST_REDIS_DB", "1")
	r := NewRedis(host, db, "TEST")

	err := r.FlushDb()
	assert.NoError(t, err)

	stats, err := r.Increment()
	assert.NoError(t, err)
	assert.Equal(t, 1, stats.Total)
	assert.Equal(t, 1, stats.Instance)

	stats, err = r.Increment()
	assert.NoError(t, err)
	assert.Equal(t, 2, stats.Total)
	assert.Equal(t, 2, stats.Instance)
}

func TestRedis_GetStatsSuccess(t *testing.T) {
	host := helper.Getenv("TEST_REDIS_HOST", "127.0.0.1:6379")
	db := helper.GetenvToi("TEST_REDIS_DB", "1")
	r := NewRedis(host, db, "TEST")

	err := r.FlushDb()
	assert.NoError(t, err)

	stats, err := r.Increment()
	assert.NoError(t, err)
	assert.Equal(t, 1, stats.Total)
	assert.Equal(t, 1, stats.Instance)

	stats, err = r.GetStats()
	assert.NoError(t, err)
	assert.Equal(t, 1, stats.Total)
	assert.Equal(t, 1, stats.Instance)
}
