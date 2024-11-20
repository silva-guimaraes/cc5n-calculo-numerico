package main

import (
	"calculoNumerico/etapa2/functions"
	"calculoNumerico/utils"
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	fmt.Println("hello etapa 2")
	// Lê o dataset do arquivo CSV
	dataset := utils.LerDataset(filepath.Join("..", "dataset.csv"))

	if len(dataset) == 0 {
		log.Fatal("O arquivo CSV está vazio!")
	}

	// Inicializa os slices xValues e yValues com o tamanho do dataset
	xValues := make([]float64, len(dataset))
	yValues := make([]float64, len(dataset))

	// Preenche os valores de x (GravidadeSintomas) e y (Humor) nos slices
	for i := 0; i < len(dataset); i++ {
		xValues[i] = float64(dataset[i].GravidadeSintomas)
		yValues[i] = float64(dataset[i].Humor)

	}

	// xValues := []float64{1, 2, 2, 4}
	// yValues := []float64{1, 4, 4, 16}

	// Teste de interpolação para um valor de x (pode ser ajustado conforme necessário)
	interpolatedValue := functions.NewtonInterpolation(xValues, yValues, 9.8)

	// Exibe o valor interpolado
	fmt.Printf("Valor interpolado para x = 2.5: %.2f\n", interpolatedValue)
}
