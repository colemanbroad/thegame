package main

import (
	"github.com/gin-gonic/gin"
)

func startServer() {
	alldata := load_data_and_jsonify()
	r := gin.Default()

	r.GET("/api", func(c *gin.Context) {
		c.JSON(200, alldata)
		// c.JSON(200, gin.H{
		// 	// "message": "pong",
		// 	"message": randSeq(20),
		// })
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func main() {
	init_logs()
	defer logfile.Close()
	startServer()
}
