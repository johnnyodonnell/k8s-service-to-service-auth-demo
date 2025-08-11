package main

import (
    "fmt"
    "time"
)


func checkService() {
    fmt.Println("Checking service")
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

