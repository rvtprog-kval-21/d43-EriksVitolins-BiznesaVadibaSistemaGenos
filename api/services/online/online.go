package online

import (
	"fmt"
	"time"
	"github.com/patrickmn/go-cache"
)

var cacheStorage *cache.Cache

func Ping(id interface{}) {
	fmt.Println(cacheStorage)
	cacheStorage.Set(fmt.Sprintf("%v", id), "true",cache.DefaultExpiration)
}

func Initiate() {
	cacheStorage = cache.New(3*time.Minute, 5*time.Minute)
}

func UsersOnline() map[string]cache.Item {
	return cacheStorage.Items()
}
