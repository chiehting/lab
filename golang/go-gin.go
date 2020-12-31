package main

import (
    "github.com/gin-gonic/gin"
    "fmt"
)

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.GET("/server", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "header": c.Request.Header,
        })
        fmt.Print(c.Request.Header);
    })
    r.Run("0.0.0.0:80")
}
