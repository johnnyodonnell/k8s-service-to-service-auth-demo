package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"
    "time"
)


func runCheck(service string) bool {
    resp, err := http.Get(service + "/health")
    if err != nil {
        fmt.Println("HTTP request error:", err)
        return false
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Read body error:", err)
        return false
    }

    var result bool
    if err := json.Unmarshal(body, &result); err != nil {
        fmt.Println("JSON unmarshal error:", err)
        return false
    }

    return result
}

func checkService() {
    service := os.Getenv("SERVICE")
    if service == "" {
        service = "http://localhost:8080"
    }

    result := runCheck(service)

    if result {
        fmt.Println(fmt.Sprintf("Service %s is up", service))
    } else {
        fmt.Println(fmt.Sprintf("Unable to connect to %s", service))
    }
}

func StartCron() {
    ticket := time.NewTicker(5 * time.Second)
    defer ticket.Stop()

    checkService()

    for {
        select {
        case <- ticket.C:
            checkService()
        }
    }
}

