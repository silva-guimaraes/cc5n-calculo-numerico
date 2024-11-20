package main

import (
	"calculoNumerico/utils"
	"fmt"
	"math"
	"path/filepath"
)

// faz a decomposição da matriz
func lud(a utils.Matriz) (l, u, d utils.Matriz, e error) {

    // criar novas matrizes do mesmo tamanho que 'a'
    l = utils.NovaMatrizVazia(a.Largura, a.Altura)
    u = utils.NovaMatrizVazia(a.Largura, a.Altura)
    d = utils.NovaMatrizVazia(a.Largura, a.Altura)

    for i := range a.Altura {
        for j := range a.Largura {

            x := a.Get(i, j)

            if i == j {
                d.Set(i, j, x)

            } else if j > i {
                u.Set(i, j, x)

            } else {
                l.Set(i, j, x)
            }
        }
    }
    return l, u, d, nil
}

func verificarDiagonal(A utils.Matriz) {
    for i := range min(A.Largura, A.Altura) {
        soma := float64(0)
        for j := range min(A.Largura, A.Altura) {
            if i != j {
                soma += math.Abs(A.Get(i,j))
            }
        }
        if math.Abs(A.Get(i,i)) < soma {
            panic(fmt.Errorf(
                "matriz não é diagonalmente dominante! ii = %f soma = %f",
                A.Get(i,i), soma,
                ))
        }
    }
}

// inverte apenas os elementos da diagonal
func diagonalInversa(d utils.Matriz) utils.Matriz {
    resultado := d.Copia()
    for i := range min(d.Largura, d.Altura) {
        x := resultado.Get(i,i)
        resultado.Set(i,i, 1 / x)
    }
    return resultado
}

func novoJacobi(iter int) func (A, b utils.Matriz) utils.Matriz {
    return func (A, b utils.Matriz) utils.Matriz {

        // importante que a soma dos elementos da diagonal seja maior que todo o restante, caso o contrario, a solução não irá convergir.
        verificarDiagonal(A)

        l, u, d, _ := lud(A)
        r := u.Sum(l)
        d_1 := diagonalInversa(d) // inversa da diagonal
        x := utils.NovaMatrizVazia(1, b.Altura)
        // At := A.T()

        for range iter {
            // d_1 * (b - r * x)
            x = d_1.ProdEscalar(b.Sub(r.ProdEscalar(x)))
        }
        return x
    }
}

// não funciona!
// func novoGaussSeidel(iter int) func (A, b utils.Matriz) utils.Matriz {
//     return func (A, b utils.Matriz) utils.Matriz {
//         verificarDiagonal(A)
//
//         l, u, d, _ := lud(A)
//         l = l.Sum(d)
//         l_1 := diagonalInversa(l)
//         x := utils.NovaMatrizVazia(1, b.Altura)
//         // At := A.T()
//
//         for range iter {
//             // l_1 * (b - u * x)
//             x = l_1.ProdEscalar(b.Sub(u.ProdEscalar(x)))
//         }
//         return x
//     }
// }

func main() {

    dataset := utils.LerDataset(filepath.Join("..", "dataset.csv"))
    columns := 3
    rows := 3

    A := utils.NovaMatrizVazia(columns, rows)
    b := utils.NovaMatrizVazia(1, rows)

    // pegar os campos que estamos interessados.
    // nesse caso estamos em interessados em descobrir como o estresse, idade e humor afetao a adesão ao tratamento de um paciente
    for i, paciente := range dataset[:rows] {
        A.Set(i, 1, float64(paciente.Estresse))
        A.Set(i, 0, float64(paciente.Idade))
        A.Set(i, 2, float64(paciente.Humor))

        b.Set(i, 0, float64(paciente.TratamenetoAdesao))
    }
    // pequena alteração apenas para fins de demonstração
    A.Set(0,0, A.Get(0,0)*100)
    A.Set(1,1, A.Get(1,1)*30)
    A.Set(2,2, A.Get(2,2)*30)


    // criar um método de jacobi com até no máximo 100 iterações
    metodo := novoJacobi(100)
    // metodo := novoGaussSeidel(100)

    // calcula e retorna o resultado
    x := utils.Vetor{ 
        Matriz: metodo(A, b),
    }

    // pretty printing adiante:

    fmt.Println()
    fmt.Println("equação:")
    for i := range A.Altura {
        soma := float64(0)
        for j := range A.Largura {
            fmt.Printf("%8.2f * x%d ", A.Get(i,j), j+1)
            soma += A.Get(i,j) * x.Get(j)
        }
        fmt.Printf("= %9.4f = %.0f\n", soma, b.Get(i, 0))
    }
    fmt.Println()

    fmt.Println("variáveis encontradas:")
    x.Print()

    fmt.Println("equação resultante:")
    for i := range A.Altura {
        soma := float64(0)
        for j := range A.Largura {
            fmt.Printf("%8.2f * %1.9f ", A.Get(i,j), x.Get(j))
            soma += A.Get(i,j) * x.Get(j)
        }
        fmt.Printf("= %9.4f ~= %.0f\n", soma, b.Get(i, 0))
    }
    
}
