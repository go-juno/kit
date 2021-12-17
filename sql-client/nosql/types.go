package nosql

import "time"

type Options struct {
	Network string `json:"network"`
	Addr    string `json:"addr"`

	Username string `json:"username"`

	Password string `json:"password"`

	DB int `json:"db"`

	MaxRetries int `json:"max_retries"`

	MinRetryBackoff time.Duration

	MaxRetryBackoff time.Duration

	DialTimeout time.Duration

	ReadTimeout time.Duration

	WriteTimeout time.Duration

	PoolFIFO bool

	PoolSize int

	PoolTimeout time.Duration
}

type ZRangeBy struct {
	Min, Max      string
	Offset, Count int64
}

type Z struct {
	Score  float64
	Member interface{}
}
