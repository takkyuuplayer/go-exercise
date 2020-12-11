package redis_test

import (
	"context"
	"os"
	"testing"

	"github.com/go-redis/redis/v8"
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

	res, err := rdb.Get(ctx, "key-does-not-exist").Result()
	if res != "" {
		t.Error("key-does-not-exist should not exist")
	}
	if err != redis.Nil {
		t.Error("key-does-not-exist should not exist")
	}
}

func TestRedisHash(t *testing.T) {
	t.Parallel()

	var ctx = context.Background()
	rdb := redisDb(t)

	rdb.HSet(ctx, "hashkey", map[string]interface{}{"foo": "bar"})

	res, err := rdb.HGetAll(ctx, "hashkey").Result()
	if len(res) != 1 {
		t.Error("hashkey should exist")
	}
	if err != nil {
		t.Error("error should not exist")
	}

	t.Log(rdb.Exists(ctx, "hashkey").Result())
	t.Log(rdb.Exists(ctx, "hashkey2").Result())

	deleted, err := rdb.Del(ctx, "hashkey").Result()
	if deleted != 1 {
		t.Error("hashkey must be deleted")
	}
	if err != nil {
		t.Error("err should not exist")
	}

	res, err = rdb.HGetAll(ctx, "key-does-not-exist").Result()
	if len(res) != 0 {
		t.Error("key-does-not-exist should not exist")
	}
	if err != nil {
		t.Error("key-does-not-exist should not exist")
	}

	t.Log(rdb.Del(ctx, "key-does-not-exist").Result())
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
