package redis

type Interface interface {
	Ping() error
	FlushDb() error
	GetInt(key string) (int, error)
	Increment(key string) (int, error)
}
