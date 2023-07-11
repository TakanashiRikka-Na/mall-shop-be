package data

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"time"
)

func Cache[T interface{}](key string, data *T, rdb *redis.Client, time time.Duration) error {
	marshal, err := json.Marshal(&data)
	if err != nil {
		return err
	}
	if err = rdb.Set(key, string(marshal), time).Err(); err != nil {
		return err
	}
	return nil
}

func ReadCache[T interface{}](key string, rdb *redis.Client, data *T) error {
	result, err := rdb.Get(key).Result()
	if err != nil {
		return err
	}
	if err = json.Unmarshal([]byte(result), &data); err != nil {
		return err
	}
	return nil
}

func DeleteCache(key string, rdb *redis.Client) error {
	return rdb.Del(key).Err()
}
