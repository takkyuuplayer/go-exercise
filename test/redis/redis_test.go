package redis_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
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

type StringSlice []string

func (s StringSlice) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *StringSlice) ScanRedis(val string) error {
	return json.Unmarshal([]byte(val), s)
}

func TestRedisHash(t *testing.T) {
	t.Parallel()

	var ctx = context.Background()
	rdb := redisDb(t)

	t.Run("when exists", func(t *testing.T) {
		t.Parallel()

		rdb.HSet(ctx, "hashkey", map[string]interface{}{"foo": "bar"})

		res, err := rdb.HGetAll(ctx, "hashkey").Result()
		assert.Len(t, res, 1)
		assert.NoError(t, err)

		t.Log(rdb.Exists(ctx, "hashkey").Result())
		t.Log(rdb.Exists(ctx, "hashkey2").Result())

		deleted, err := rdb.Del(ctx, "hashkey").Result()
		assert.Equal(t, deleted, int64(1))
		assert.NoError(t, err)
	})

	t.Run("struct", func(t *testing.T) {
		type Sample struct {
			Foo   string      `json:"foo" redis:"foo"`
			Slice StringSlice `json:"slice" redis:"slice"`
		}

		s := Sample{
			Foo:   "bar",
			Slice: []string{"a", "b", "c"},
		}

		assert.NoError(t, rdb.HSet(ctx, "hashkey", s).Err())

		var sample Sample
		err := rdb.HGetAll(ctx, "hashkey").Scan(&sample)
		assert.Equal(t, s, sample)
		assert.NoError(t, err)

		var sample2 Sample
		err = rdb.HGetAll(ctx, "key-does-not-exists").Scan(&sample2)
		assert.Equal(t, Sample{}, sample2)
		assert.NoError(t, err)

		var sample3 Sample
		res := rdb.HGetAll(ctx, "hashkey")
		hash, err := res.Result()
		t.Log(hash, err)
		assert.NoError(t, err)
		assert.NoError(t, res.Scan(&sample3))
		assert.Equal(t, s, sample3)
	})

	t.Run("when not exists", func(t *testing.T) {
		t.Parallel()

		res, err := rdb.HGetAll(ctx, "key-does-not-exist").Result()
		assert.Len(t, res, 0)
		assert.NoError(t, err)

		deleted, err := rdb.Del(ctx, "key-does-not-exist").Result()
		assert.Equal(t, deleted, int64(0))
		assert.NoError(t, err)
	})
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
