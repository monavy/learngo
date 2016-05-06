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

    factors := euler.Factor(n)

    fmt.Printf("The factors of %d are %v.\n", n, factors)
    fmt.Printf("The product of %v is %d.\n", factors, product(factors))
}


func usage() {
    fmt.Println("Usage: factor number")
    os.Exit(1)
}


func product(nums []uint64) uint64 {
    product := uint64(1)

    for i := 0; i < len(nums); i++ {
        product = product * nums[i]
    }

    return product
}
