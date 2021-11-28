package main

import (
	"github.com/gin-gonic/gin"
	"nacos-go-demo/naming"
	"strconv"
)

func main() {
	engine := gin.Default()
	// naming
	engine.GET("/naming/register", func(context *gin.Context) {
		values := context.Request.URL.Query()
		ip := values.Get("ip")
		port := values.Get("port")
		pr, _ := strconv.ParseUint(port, 0, 0)
		result := naming.Register(ip, pr)
		context.JSON(200, gin.H{"message": result})

	})
	engine.GET("/naming/get", func(context *gin.Context) {
		values := context.Request.URL.Query()
		serviceName := values.Get("serviceName")
		cluster := values.Get("cluster")
		group := values.Get("group")
		instance := naming.GetInstance(serviceName, cluster, group)
		context.JSON(200, gin.H{"message": instance})
	})

	engine.Run()
}
