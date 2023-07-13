package data

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"log"
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

func DeleteCache(rdb *redis.Client, keys ...string) error {
	for _, key := range keys {
		if err := rdb.Del(key).Err(); err != nil {
			//如果错误原因是空值就 continue
			if err.Error() == "redis: nil" {
				continue
			}
			log.Println(err)
			return err
		}
	}
	return nil
}

// 游标分页

func StoreCursor(rdb *redis.Client, cursor string, id int) error {
	err := rdb.Set("cursor", cursor, 0).Err()
	if err != nil {
		return err
	}
	err = rdb.Set(cursor, id, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func StorePageResult[T interface{}](rdb *redis.Client, pages []T, key string, pageId int) error {
	var z []redis.Z
	for _, page := range pages {
		z = append(z, redis.Z{Score: float64(pageId), Member: page})
	}
	return rdb.ZAdd(key, z...).Err()

}

func GetPagesResult[T interface{}](rdb *redis.Client, key string) ([]T, error) {
	results, err := rdb.ZRange(key, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	var ts []T
	for _, result := range results {
		var t T
		err = json.Unmarshal([]byte(result), &t)
		if err != nil {
			return nil, err
		}
		ts = append(ts, t)
	}
	return ts, nil
}
