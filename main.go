package main

import (
	_ "recitewords/routers"
	"github.com/astaxie/beego"
	"fmt"
)

func main() {
	beego.Run()
	fmt.Println(beego.AppConfig.String("redishost"))
}

