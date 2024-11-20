package utils

import (
	"encoding/csv"
	"fmt"
	"log"
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
	var copia []float64
	copy(copia, m.elementos)
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
			fmt.Printf("%f ", m.Get(i, j))
		}
		fmt.Println()
	}
	fmt.Println()
}

func (m Matriz) Inverso() Matriz {
	resultado := m.Copia()
	for i := range m.Altura * m.Largura {
		resultado.elementos[i] = 1 / resultado.elementos[i]
	}
	return resultado
}

// multiplicação de matrizes
func (a Matriz) ProdEscalar(b Matriz) Matriz {
	if a.Largura != b.Altura {
		log.Panic(
			"numero de colunas da matriz A é diferente do numero do" +
				"numero de linhas da matriz B")
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

func (a Vetor) Sub(b Vetor) Vetor {
	if a.Altura != b.Altura {
		log.Panic("tamanho de vetor A é diferente do tamanho do vetor B")
	}
	resultado := NovoVetor(a.elementos)
	for i := range a.Altura {
		resultado.Set(i, a.Get(i)-b.Get(i))
	}
	return resultado
}

func ParseFloat64(value string) float64 {
	parsedValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		// Se falhar, retorna 0
		return 0
	}
	return parsedValue
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
