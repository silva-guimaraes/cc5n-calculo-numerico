package utils_test

import (
	. "calculoNumerico/utils"
	"testing"
)

func TestProdEscalar(t *testing.T) {
	a := NovaMatriz(3, 2, []float64{1, 2, 3, 4, 5, 6})
	b := NovaMatriz(2, 3, []float64{7, 8, 9, 10, 11, 12})

	esperado := NovaMatriz(2, 2, []float64{58, 64, 139, 154})
	if !Eq(a.ProdEscalar(b), esperado) {
		t.Fatalf("oops!")
	}
}

func TestCriarMatriz(t *testing.T) {
	a := NovaMatriz(3, 3, []float64{0, 0, 0, 0, 0, 0, 0, 0, 0})
	b := a.Copia()
	if !Eq(a, b) {
		t.Fatalf("oops!")
	}
}

func TestInverso(t *testing.T) {
    a := NovoVetor([]float64{ 1.0/2.0, 1.0/4.0, 1.0/5.0, 1.0/8.0, 1.0/10.0 })
    esperado := NovoVetor([]float64{2,4,5,8,10})
    resultado := a.Inverso()

    if !Eq(resultado, esperado.Matriz) {
        resultado.Print()
        esperado.Print()
		t.Fatalf("oops!")
    }
}

func TestSub(t *testing.T) {
	a := NovoVetor([]float64{1, 2, 3, 4, 5, 6})
	b := NovoVetor([]float64{1, 2, 3, 4, 5, 6})

	esperado := NovoVetor([]float64{0, 0, 0, 0, 0, 0})
	resultado := a.Sub(b.Matriz)

	if !Eq(esperado.Matriz, resultado) {
		t.Fatalf("oops!")
	}
}
