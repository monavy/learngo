package main

import (
    "fmt"
    "os"
    "strconv"
    "euler"
)


func main() {
    if len(os.Args) != 2 { usage() }

    n, err := strconv.ParseUint(os.Args[1], 10, 64)
    if err != nil { usage() }

    prime, divisor := euler.IsPrime(n)

    if prime == true {
        fmt.Printf("The number %d is prime.\n", n)
    } else {
        fmt.Printf("The number %d is divisible by %d.\n", n, divisor)
    }
}


func usage() {
    fmt.Println("Usage: isprime number")
    os.Exit(1)
}
