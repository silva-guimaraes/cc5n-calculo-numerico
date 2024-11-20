package functions

import (
	"fmt"
)

// NewtonInterpolation realiza a interpolação de Newton para um conjunto de pontos dados
func NewtonInterpolation(xValues []float64, yValues []float64, x float64) float64 {
	n := len(xValues)

	// Verificar e remover pontos duplicados
	uniqueXValues := make([]float64, 0, n)
	uniqueYValues := make([]float64, 0, n)

	for i := 0; i < n; i++ {
		// Verificar se o valor de x já existe na lista de valores únicos
		isDuplicate := false
		for j := 0; j < len(uniqueXValues); j++ {
			if xValues[i] == uniqueXValues[j] {
				isDuplicate = true
				break
			}
		}

		// Se não for duplicado, adicionar ao array de valores únicos
		if !isDuplicate {
			uniqueXValues = append(uniqueXValues, xValues[i])
			uniqueYValues = append(uniqueYValues, yValues[i])
		}
	}

	// Verificar se há pontos suficientes após remover duplicados
	if len(uniqueXValues) < 2 {
		panic("Erro: Não há pontos suficientes para interpolação!")
	}

	// Atualizar o número de pontos após remoção de duplicatas
	n = len(uniqueXValues)

	// Criar a tabela de diferenças divididas
	divDifferences := make([][]float64, n)
	for i := range divDifferences {
		divDifferences[i] = make([]float64, n)
	}

	// Preencher a primeira coluna com yValues
	for i := 0; i < n; i++ {
		divDifferences[i][0] = uniqueYValues[i]
	}

	// Calcular as diferenças divididas
	for j := 1; j < n; j++ {
		for i := 0; i < n-j; i++ {
			// Calcular a diferença dividida
			denom := uniqueXValues[i+j] - uniqueXValues[i]
			if denom == 0 {
				panic(fmt.Sprintf("Erro: Ponto duplicado em x = %f", uniqueXValues[i]))
			}
			divDifferences[i][j] = (divDifferences[i+1][j-1] - divDifferences[i][j-1]) / denom
		}
	}

	// Calcular o valor interpolado usando a tabela de diferenças divididas
	interpolatedValue := divDifferences[0][0]
	productTerm := 1.0
	for i := 1; i < n; i++ {
		productTerm *= (x - uniqueXValues[i-1]) // Produto acumulado
		interpolatedValue += divDifferences[0][i] * productTerm
	}

	return interpolatedValue
}
