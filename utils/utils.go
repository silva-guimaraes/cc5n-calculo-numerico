package utils

import (
	"encoding/csv"
	"os"
	"strconv"
)

type Row struct {
    pacienteID int
    idade int
    sexo string
    diagnostico string
    gravidadeSintomas int
    humor int
    sonoQualidade int
    atividadeFisica int
    medicacao string
    terapiaTipo string
    tratamentoInicio string
    tratamentoDuracao int
    estresse int
    resultado string
    tratamentoProgresso int
    estadoEmocional string
    tratamenetoAdesao int
}

func mustAtoi(s string) int {
    i, err := strconv.Atoi(s)
    if err != nil {
        panic(err)
    }
    return i
}

func ReadDataset(path string) []Row {
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
        linha := Row {
            // Paciente ID
            pacienteID: mustAtoi(r[0]),
            // Idade
            idade: mustAtoi(r[1]),
            // Sexo
            sexo: r[2],
            // Diagnóstico
            diagnostico: r[3],
            // Gravidade dos sintomas (1-10)
            gravidadeSintomas: mustAtoi(r[4]),
            // Nível do humor (1-10)
            humor: mustAtoi(r[5]),
            // Qualidade do sono (1-10)
            sonoQualidade: mustAtoi(r[6]),
            // Atividade Física (hrs/semana)
            atividadeFisica: mustAtoi(r[7]),
            // Medicação
            medicacao: r[8],
            // Tipo de Terapia
            terapiaTipo: r[9],
            // Início do Tratamento
            tratamentoInicio: r[10],
            // Duração do tratamento (semanas)
            tratamentoDuracao: mustAtoi(r[11]),
            // Nível de Estresse (1-10)
            estresse: mustAtoi(r[12]),
            // Resultado
            resultado: r[13],
            // Progresso do tratamento (1-10)
            tratamentoProgresso: mustAtoi(r[14]),
            // Estado emocional detectado por IA
            estadoEmocional: r[15],
            // Adesão ao tratamento (%)
            tratamenetoAdesao: mustAtoi(r[16]),
        }
        rows = append(rows, linha)
    }
    return rows
}
