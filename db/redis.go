package db

import (
	"github.com/juliotorresmoreno/searchbar/config"
	redis "gopkg.in/redis.v5"
)

//Cache Almacen temporal de datos
type Cache struct {
	*redis.Client
}

//FindKeys busca y encuentra los keys
func (cache Cache) FindKeys(str string) []string {
	c := cache.Scan(0, "word_*"+str+"*", 10)
	r, _, _ := c.Result()
	return r
}

// GetCache Obtiene la cache
func GetCache() *Cache {
	db := redis.NewClient(&redis.Options{
		Addr: config.REDIS_HOST + ":" + config.REDIS_PORT,
	})
	return &Cache{Client: db}
}
