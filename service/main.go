package main

import (
    "fmt"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
)



func isHealthy(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, true)
}

func main() {
    fmt.Println("Starting up...")

    go StartCron()

    router := gin.Default()
    router.GET("/health", isHealthy)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    addr := fmt.Sprintf("localhost:%s", port)
    router.Run(addr)
}

