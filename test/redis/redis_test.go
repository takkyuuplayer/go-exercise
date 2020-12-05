package redis_test

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
	"testing"
)

func TestRedis(t *testing.T) {
	t.Parallel()

	var ctx = context.Background()

	rdb := redisDb(t)
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		t.Fatal(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		t.Fatal(err)
	}
	if val != "value" {
		t.Errorf("Result() got = %v, want %v", val, "value")
	}

	_, err = rdb.Get(ctx, "key2").Result()
	if err == nil {
		t.Error("key2 should not exist")
	} else if err != redis.Nil {
		t.Fatal(err)
	}
}

func redisDb(t *testing.T) *redis.Client {
	t.Helper()

	var db *redis.Client
	if host, ok := os.LookupEnv("REDIS_HOST"); ok && host != "" {
		db = redis.NewClient(&redis.Options{
			Addr: host,
		})
	} else {
		db = redis.NewClient(&redis.Options{})
	}

	return db
}
