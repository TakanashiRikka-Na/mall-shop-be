package data

import (
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"testing"
)

func TestNewRedisConn(t *testing.T) {
	if err := redis.NewClient(&redis.Options{DB: 0, Addr: "127.0.0.1:6379"}).Ping().Err(); err != nil {
		t.Error(err)
	}
}

func TestNewDB(t *testing.T) {
	_, err := gorm.Open(mysql.Open(fmt.Sprintf("admin:%v@tcp(43.143.227.115:3306)/mall?%v",
		os.Getenv("mall_password"),
		os.Getenv("mysql.param"))))
	if err != nil {
		t.Error(err)
	}
}
