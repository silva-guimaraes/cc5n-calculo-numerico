package functions

import (
	"fmt"
	"math"
)

// NewtonInterpolation realiza a interpolação de Newton para um conjunto de pontos dados
func NewtonInterpolation(xValues []float64, yValues []float64, x float64) float64 {
	n := len(xValues)

	// Verificar se existem pontos duplicados (para evitar divisão por zero)
	for i := 0; i < n-1; i++ {
		if xValues[i] == xValues[i+1] {
			fmt.Printf("Erro: xValues[%d] = xValues[%d] = %.2f, pontos duplicados encontrados.\n", i, i+1, xValues[i])
			return math.NaN() // Retorna NaN para indicar erro
		}
	}

	// Criar a tabela de diferenças divididas
	divDifferences := make([][]float64, n)
	for i := range divDifferences {
		divDifferences[i] = make([]float64, n)
	}

	// Preencher a primeira coluna com yValues
	for i := 0; i < n; i++ {
		divDifferences[i][0] = yValues[i]
	}

	// Calcular as diferenças divididas
	for j := 1; j < n; j++ {
		for i := 0; i < n-j; i++ {
			// Calcular a diferença dividida
			denom := xValues[i+j] - xValues[i]
			if denom == 0 {
				// Se o denominador for zero, há um problema com os dados
				fmt.Printf("Erro: diferença entre x[%d] e x[%d] é zero. Não é possível calcular a diferença dividida.\n", i, i+j)
				return math.NaN() // Retorna NaN para indicar erro
			}
			divDifferences[i][j] = (divDifferences[i+1][j-1] - divDifferences[i][j-1]) / denom
		}
	}

	// Calcular o valor interpolado usando a tabela de diferenças divididas
	interpolatedValue := divDifferences[0][0]
	productTerm := 1.0
	for i := 1; i < n; i++ {
		productTerm *= (x - xValues[i-1]) // Produto acumulado
		interpolatedValue += divDifferences[0][i] * productTerm
	}

	return interpolatedValue
}
