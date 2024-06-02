package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Message struct {
    Message string `json:"message"`
}

func main() {
    router := gin.Default()

    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello from Backend Server! This is the main endpoint.\n")
    })

    router.GET("/health", func(c *gin.Context) {
        c.String(http.StatusOK, "OK")
    })

    router.POST("/message", func(c *gin.Context) {
        var message Message
        if err := c.BindJSON(&message); err != nil {
            c.String(http.StatusBadRequest, "Invalid request")
            return
        }
        c.JSON(http.StatusOK, gin.H{"response": "Client sent: " + message.Message})
    })

    router.Run(":5000")
}
