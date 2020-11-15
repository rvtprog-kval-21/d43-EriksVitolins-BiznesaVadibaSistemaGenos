package online

import (
	"fmt"
	"time"
)

var online map[interface{}]time.Time

func Ping(id interface{}) {
	initiate()
	test := online
	fmt.Println(test)
	online[id] = time.Now()
}

func initiate() {
	if online == nil {
		online = make(map[interface{}]time.Time)
	}
}

func ClearOldOnes() {
	for key, date := range online {
		if time.Since(date) > 1*time.Minute {
			delete(online, key)
		}
	}

	time.Sleep(1 * time.Minute)
	ClearOldOnes()
}
