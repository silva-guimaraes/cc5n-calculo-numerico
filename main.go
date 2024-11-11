package main

import (
    "fmt"
    "calculoNumerico/dataset"
)

func main() {
    fmt.Println("hello world")
    dataset := read_dataset.Read()
    fmt.Println(dataset[0])
}
