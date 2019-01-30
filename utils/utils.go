package utils

import (
	"fmt"
	"time"
)

func TimeCost(start time.Time) {
	terminal := time.Since(start)
	fmt.Println(terminal)
}
