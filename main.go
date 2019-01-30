package main

import (
	"fmt"
	"time"

	_ "net/http/pprof"

	. "github.com/infoCollection/collection"
	. "github.com/infoCollection/database"
	. "github.com/infoCollection/utils"
)

func main() {
	OpenSQL()
	defer TimeCost(time.Now())
	fmt.Println("开始采集")
	Collection()
}
