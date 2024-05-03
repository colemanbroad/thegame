package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	r := gin.Default()

	val := map[string]any{"marmalade": "qualcom", "boogienight": []int{1, 2, 3}}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, val)
		// c.JSON(200, gin.H{
		// 	// "message": "pong",
		// 	"message": randSeq(20),
		// })
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
