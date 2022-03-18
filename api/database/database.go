package database

import(
	"context"

	"github.com/go-redis/redis/v8"
	"os"
)

var Ctx = context.Background() /* create a conext */

func CreateClient(dbNo int) *redis.CLient { /*  */
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("DB_ADDR")
		Password: os.Getenv("DB_PASS")
		DB: dbNo, 
	})
}
