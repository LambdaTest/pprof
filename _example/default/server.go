package main

import (
	"github.com/LambdaTest/gin"
	"github.com/LambdaTest/pprof"
)

func main() {
	router := gin.Default()
	pprof.Register(router)
	router.Run(":8080")
}
