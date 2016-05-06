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

    primes := euler.Sieve(n)

    fmt.Printf("There are %d primes less than %d.\n", len(primes), n)
}


func usage() {
    fmt.Println("Usage: sieve number")
    os.Exit(1)
}
