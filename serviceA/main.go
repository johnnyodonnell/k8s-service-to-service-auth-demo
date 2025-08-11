package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)


func isHealthy(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, true)
}

func main() {
    fmt.Println("ServiceA starting up...")

    router := gin.Default()
    router.GET("/health", isHealthy)

    router.Run("localhost:8080")
}

