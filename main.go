// main.go
package main

import (
    "fmt"
    "example.com/new-go/mascot"  // Import your custom mascot package
)

func main() {
    fmt.Println("Hello, World!")
    mascot.PrintMascot()  // Call the function from the mascot package
}
