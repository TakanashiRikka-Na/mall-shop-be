package data

import (
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mall-shop-be/internal/conf"
	"mall-shop-be/internal/errors"
	"os"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewRedisConn, NewDB, NewUserRepo, errors.NewErrorController)

// Data .github.com/go-redis/redis
type Data struct {
	DB  *gorm.DB
	Rdb *redis.Client
}

const (
	UserError = "user error"
)

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, rdb *redis.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{DB: db, Rdb: rdb}, cleanup, nil
}

func NewRedisConn(c *conf.Data) *redis.Client {
	return redis.NewClient(&redis.Options{
		DB:       0,
		Addr:     c.Redis.Addr,
		Password: "",
	})
}

func NewDB(c *conf.Data) *gorm.DB {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?%v",
		c.Database.User,
		os.Getenv("mall_password"),
		c.Database.Host,
		c.Database.Port,
		c.Database.DBName,
		c.Database.Param,
	)), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		panic(err)
	}
	return db
}
