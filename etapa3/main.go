package main

import (
	"calculoNumerico/utils"
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println("hello etapa 3")
	dataset := utils.LerDataset(filepath.Join("..", "dataset.csv"))
	fmt.Println(dataset[1])
}
