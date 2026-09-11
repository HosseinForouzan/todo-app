package redistask

import (
	"fmt"

	goredis "github.com/redis/go-redis/v9"
)

type Config struct {
	Host     string 
	Port     int    
	Password string 
	DB       int    

}

type Adapter struct {
	client *goredis.Client
}

func New(config Config) Adapter {
	rdb := goredis.NewClient(&goredis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
	})

	return Adapter{client: rdb}
}

func (a Adapter) Client() *goredis.Client {
	return a.client
}