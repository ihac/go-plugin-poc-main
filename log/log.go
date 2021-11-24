package log

import (
	"fmt"
	"time"
)

type Log struct{
	id int64
}

var log *Log

func init() {
	fmt.Println("init() being called")
	log = &Log{
		id: time.Now().UnixNano(),
	}
}

func Info(msg string) {
	fmt.Printf("[%d]: %s\n", log.id, msg)
}
