package online

import (
	"fmt"
	"time"
)

var Online map[interface{}]time.Time

func Ping(id interface{}) {
	initiate()
	test := Online
	fmt.Println(test)
	Online[id] = time.Now()
}

func initiate() {
	if Online == nil {
		Online = make(map[interface{}]time.Time)
	}
}

func ClearOldOnes() {
	for key, date := range Online {
		if time.Since(date) > 1*time.Minute {
			delete(Online, key)
		}
	}

	time.Sleep(1 * time.Minute)
	ClearOldOnes()
}
