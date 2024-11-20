# cc5n-calculo-numerico
## Setup
## Baixar projeto
```sh
git clone https://github.com/silva-guimaraes/cc5n-calculo-numerico
cd cc5n-calculo-numerico
go mod tidy
```
## Rodar
```sh
cd etapa2
go run .
```
```sh
cd etapa3
go run .
```
## Etapas
## Etapa 1 - Coleta e Estruturação dos Dados
Analisamos um banco de dados relacionado ao tratamento e exames de pacientes com problemas mentais, disponível [neste link](https://www.kaggle.com/datasets/uom190346a/mental-health-diagnosis-and-treatment-monitoring).
Os dados estão formatados em CSV e incluem as seguintes colunas:
- Paciente ID
- Idade
- Sexo
- Diagnóstico
- Gravidade dos sintomas (1-10)
- Nível do humor (1-10)
- Qualidade do sono (1-10)
- Atividade Física (hrs/semana)
- Medicação
- Tipo de Terapia
- Início do Tratamento	Duração do tratamento (semanas)
- Nível de Estresse (1-10)
- Resultado
- Progresso do tratamento (1-10)
- Estado emocional detectado por IA
- Adesão ao tratamento (%)

## Etapa 2 - Interpolação Polinomial para Estimar Variáveis

## Etapa 3 - Resolução de um Sistema Linear com Métodos Iterativos
Com o objetivo de estimar a adesão ao tratamento para futuros pacientes, desenvolvemos um sistema linear utilizando informações relacionadas à idade, humor e estresse. A relação entre esses fatores e a porcentagem de adesão foi modelada como uma soma de produtos:

$$
adesão = idade \cdot x_1 + humor \cdot x_2 + estresse \cdot x_3
$$

Para resolver este sistema, utilizamos o Método de Jacobi, um método iterativo aplicado para solucionar sistemas lineares. Nosso problema foi formulado como um sistema linear de duas matrizes, permitindo a aplicação do método. Após a execução, encontramos os seguintes valores para os coeficientes:

$$x1 = 0.013850245$$

$$x2 = 0.353243993$$

$$x3 = 0.652950318$$

## Relatório Final

## To Do
 - [x] Criar o repositório
 - [x] Encontrar um dataset (Etapa 1)
 - [x] Testes para ambas as etapas
 - [x] Etapa 2
 - [x] Etapa 3
 - [ ] Etapa 4
