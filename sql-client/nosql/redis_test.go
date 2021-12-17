package nosql

import (
	"context"
	"fmt"
	"testing"
)

func TestRedisNoSqlClient(t *testing.T) {
	O := &Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
	}
	r, err := RedisNoSQLClient(O)
	if err != nil {
		return
	}
	set, err := r.Set(context.Background(), "w", "xx", 0)
	if err != nil {
		return
	}
	fmt.Println(set)
	get, err := r.Get(context.Background(), "w")
	if err != nil {
		return
	}
	fmt.Println(get)

}
