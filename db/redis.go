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
	result := make([]string, 0)
	var elements []string
	var cursor uint64
	for exec := true; exec; {
		tmp := cache.Scan(cursor, "word_*"+str+"*", 10)
		elements, cursor, _ = tmp.Result()
		for _, v := range elements {
			result = append(result, v)
		}
		if cursor == 0 {
			exec = false
		}
	}
	return result
}

// GetCache Obtiene la cache
func GetCache() *Cache {
	db := redis.NewClient(&redis.Options{
		Addr: config.REDIS_HOST + ":" + config.REDIS_PORT,
	})
	return &Cache{Client: db}
}
