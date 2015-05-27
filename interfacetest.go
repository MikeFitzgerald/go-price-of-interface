package main

import "fmt"

const (
    debug byte = iota
    prod
)

var logLevel = prod

func main() {
    logIface(debug, "Test interface")
    logString(debug, "Test string")
}

func logIface(level byte, msg interface{}) {
    if level >= logLevel {
        fmt.Println(msg)
    }
}

func logString(level byte, msg string) {
    if level >= logLevel {
        fmt.Println(msg)
    }
}
