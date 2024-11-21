package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

type Row struct {
	PacienteID          int
	Idade               int
	Sexo                string
	Diagnostico         string
	GravidadeSintomas   int
	Humor               int
	SonoQualidade       int
	AtividadeFisica     int
	Medicacao           string
	TerapiaTipo         string
	TratamentoInicio    time.Time
	TratamentoDuracao   int
	Estresse            int
	Resultado           string
	TratamentoProgresso int
	EstadoEmocional     string
	TratamenetoAdesao   int
}

type Matriz struct {
	Altura, Largura int
	elementos       []float64
}

// isso é praticamente herança em Go
type Vetor struct {
	Matriz
}

func NovaMatriz(largura, altura int, elementos []float64) Matriz {
	if largura*altura != len(elementos) {
		panic("largura e altura não conferem com a quantidade de elementos na lista")
	}
	return Matriz{
		Altura:    altura,
		Largura:   largura,
		elementos: elementos,
	}
}

func NovoVetor(elementos []float64) Vetor {
	return Vetor{
		Matriz: NovaMatriz(1, len(elementos), elementos),
	}
}

func NovaMatrizVazia(largura, altura int) Matriz {
	return Matriz{
		Largura:   largura,
		Altura:    altura,
		elementos: make([]float64, largura*altura),
	}
}

func (m Matriz) Copia() Matriz {
    copia := make([]float64, m.Largura*m.Altura)
    _ = copy(copia, m.elementos)
    // if copied == 0 {
    //     panic("zero copy")
    // }
	return Matriz{
		Largura:   m.Largura,
		Altura:    m.Altura,
		elementos: copia,
	}
}

func (m Matriz) Get(i, j int) float64 {
	return m.elementos[i*m.Largura+j]
}
func (m *Matriz) Set(i, j int, x float64) {
	m.elementos[i*m.Largura+j] = x
}
func (v Vetor) Get(i int) float64 {
	return v.elementos[i]
}
func (v *Vetor) Set(i int, x float64) {
	v.elementos[i] = x
}

// para debug
func (m Matriz) Print() {
	for i := range m.Altura {
		for j := range m.Largura {
			fmt.Printf("%10f ", m.Get(i, j))
		}
		fmt.Println()
	}
	fmt.Println()
}

// para debug
func Eq(a, b Matriz) bool {
	if a.Largura != b.Largura || a.Altura != b.Altura {
		return false
	}
	for i := range a.Largura * a.Altura {
		if a.elementos[i] != b.elementos[i] {
			return false
		}
	}
	return true
}

func (m *Matriz) Inverso() Matriz {
	resultado := m.Copia()
	for i := range m.Altura * m.Largura {
        x := 1 / m.elementos[i]
        if x == math.Inf(1) {
            panic("divisão por 0")
        }
		resultado.elementos[i] = x
	}
	return resultado
}

// multiplicação de matrizes
func (a Matriz) ProdEscalar(b Matriz) Matriz {
	if a.Largura != b.Altura {
		log.Panicf(
			"numero de colunas da matriz A é diferente do numero " +
            "de linhas da matriz B: %d != %d\n", a.Largura, b.Altura)
	}
	resultado := NovaMatrizVazia(b.Largura, a.Altura)
	for i := range resultado.Altura {
		for j := range resultado.Largura {
			var soma float64 = 0
			for r := range a.Largura {
				ax := a.Get(i, r)
				bx := b.Get(r, j)
				soma += ax * bx
			}
			resultado.Set(i, j, soma)
		}
	}
	return resultado
}

// subtração de matrizes
func (a Matriz) Sub(b Matriz) Matriz {
	if len(a.elementos) != len(b.elementos) {
		log.Panic("tamanho de matriz A é diferente do tamanho do matriz B")
	}
	resultado := NovaMatrizVazia(a.Largura, a.Altura)

	for i := range a.Altura * a.Largura {
		resultado.elementos[i] = a.elementos[i] - b.elementos[i]
	}
	return resultado
}

// soma de matrizes
func (a Matriz) Sum(b Matriz) Matriz {
	if len(a.elementos) != len(b.elementos) {
		log.Panic("tamanho de matriz A é diferente do tamanho do matriz B")
	}
	resultado := NovaMatrizVazia(a.Largura, a.Altura)

	for i := range a.Altura * a.Largura {
		resultado.elementos[i] = a.elementos[i] + b.elementos[i]
	}
	return resultado
}

func (a Matriz) T() Matriz {
    resultado := a.Copia()
    for i := range a.Altura {
        for j := range a.Largura {
            temp := resultado.Get(i,j)
            resultado.Set(i,j, resultado.Get(j,i))
            resultado.Set(j,i, temp)
        }
    }
    return resultado
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func LerDataset(path string) []Row {
	csvfile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	records, err := csv.NewReader(csvfile).ReadAll()
	if err != nil {
		panic(err)
	}

	var rows []Row

	for _, r := range records[1:] {
		if len(r) != 17 {
			panic(r)
		}
		tratamentoInicio, err := time.Parse(time.DateOnly, r[10])
		if err != nil {
			panic(err)
		}
		linha := Row{
			// Paciente ID
			PacienteID: mustAtoi(r[0]),
			// Idade
			Idade: mustAtoi(r[1]),
			// Sexo
			Sexo: r[2],
			// Diagnóstico
			Diagnostico: r[3],
			// Gravidade dos sintomas (1-10)
			GravidadeSintomas: mustAtoi(r[4]),
			// Nível do humor (1-10)
			Humor: mustAtoi(r[5]),
			// Qualidade do sono (1-10)
			SonoQualidade: mustAtoi(r[6]),
			// Atividade Física (hrs/semana)
			AtividadeFisica: mustAtoi(r[7]),
			// Medicação
			Medicacao: r[8],
			// Tipo de Terapia
			TerapiaTipo: r[9],
			// Início do Tratamento
			TratamentoInicio: tratamentoInicio,
			// Duração do tratamento (semanas)
			TratamentoDuracao: mustAtoi(r[11]),
			// Nível de Estresse (1-10)
			Estresse: mustAtoi(r[12]),
			// Resultado
			Resultado: r[13],
			// Progresso do tratamento (1-10)
			TratamentoProgresso: mustAtoi(r[14]),
			// Estado emocional detectado por IA
			EstadoEmocional: r[15],
			// Adesão ao tratamento (%)
			TratamenetoAdesao: mustAtoi(r[16]),
		}
		rows = append(rows, linha)
	}
	return rows
}
