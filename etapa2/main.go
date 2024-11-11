package main

import (
	"calculoNumerico/utils"
	"fmt"
	"path/filepath"
)

func main() {
    fmt.Println("hello etapa 2")
    dataset := utils.ReadDataset(filepath.Join("..", "dataset.csv"))
    fmt.Println(dataset[0])
}
