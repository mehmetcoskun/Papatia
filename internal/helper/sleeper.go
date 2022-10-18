package helper

import (
	"fmt"
	"time"
)

func Sleeper(seconds, key int) {
	if key%100 == 0 && key != 0 {
		fmt.Println("Sleeping for 15 seconds...")
		time.Sleep(15 * time.Second)
	}
}
