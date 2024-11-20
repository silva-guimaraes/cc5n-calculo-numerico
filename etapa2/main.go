package main

import (
	"calculoNumerico/etapa2/functions"
	"calculoNumerico/utils"
	"fmt"
	"log"
	"math/rand"
	"path/filepath"
	"time"
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

	var x float64

	// Preenche os valores de x (GravidadeSintomas) e y (Humor) nos slices
	for i := 0; i < len(dataset); i++ {
		xValues[i] = float64(dataset[i].GravidadeSintomas)
		yValues[i] = float64(dataset[i].Humor)

	}

	result := make([]float64, 20)

	// xValues := []float64{1, 2, 2, 4}
	// yValues := []float64{1, 4, 4, 16}
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	//com valores aleatorios
	for i := 0; i < 20; i++ {
		x = r.Float64() * 10
		interpolatedValue := functions.NewtonInterpolation(xValues, yValues, x)
		result[i] = interpolatedValue

		fmt.Println(result[i])
	}

	arr := []float64{4.5, 9.8, 7.1, 3.3, 1.4, 8.9, 6.6, 2.7, 5.5, 10.0, 4.2, 7.8, 3.9, 1.2, 8.3, 6.1, 2.5, 5.7, 9.3, 7.0}

	for i := range arr {
		interpolatedValue := functions.NewtonInterpolation(xValues, yValues, arr[i])
		fmt.Println(interpolatedValue)
	}

	/*
		--> resultados:
			arr := []float64{
				6.289062499999999,
				4.347328000000001,
				6.244433499999999,
				29.7579405,
				129.60230399999998,
				5.248616499999999,
				4.828095999999999,
				52.24935949999997,
				2.2109375,
				5,
				9.847871999999997,
				7.081728,
				14.830241500000003,
				144.89747199999994,
				6.621315499999998,
				3.269758499999999,
				61.570312499999986,
				2.3873845,
				4.370990499999999,
				6,
			}

	*/
}
