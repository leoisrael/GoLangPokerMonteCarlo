package prob

import (
	"encoding/json"
	"fmt"
	"sort"
)

// Estrutura representando uma carta
type Carta struct {
	Naipe string `json:"naipe"`
	Valor string `json:"valor"`
}

// Estrutura representando a entrada para os cálculos de probabilidade
type EntradaProbabilidade struct {
	Mao  []Carta `json:"mao"`
	Mesa []Carta `json:"mesa"`
}

// ResultadoProbabilidades ajustado para corresponder às classes CSS do front-end
type ResultadoProbabilidades struct {
	RoyalFlush    float64 `json:"Royal-flush"`
	StraightFlush float64 `json:"Straight-flush"`
	FourOfAKind   float64 `json:"Four-of-a-kind"`
	FullHouse     float64 `json:"Full-house"`
	Flush         float64 `json:"Flush"`
	Straight      float64 `json:"Straight"`
	ThreeOfAKind  float64 `json:"Three-of-a-kind"`
	TwoPair       float64 `json:"Two-pair"`
	Pair          float64 `json:"Pair"`
}

// ProbabilidadeFlop calcula e retorna as probabilidades de formar mãos específicas após o flop.
func ProbabilidadeFlop(entrada EntradaProbabilidade) string {
	resultado := ResultadoProbabilidades{
		RoyalFlush:    calculaProbabilidadeRoyalFlush(entrada),
		StraightFlush: calculaProbabilidadeStraightFlush(entrada),
		FourOfAKind:   calculaProbabilidadeFourOfAKind(entrada),
		FullHouse:     calculaProbabilidadeFullHouse(entrada),
		Flush:         calculaProbabilidadeFlush(entrada),
		Straight:      calculaProbabilidadeStraight(entrada),
		ThreeOfAKind:  calculaProbabilidadeThreeOfAKind(entrada),
		TwoPair:       calculaProbabilidadeTwoPair(entrada),
		Pair:          calculaProbabilidadePair(entrada),
	}

	// Serializa o resultado para JSON
	jsonResultado, err := json.Marshal(resultado)
	if err != nil {
		fmt.Println("Erro ao serializar o resultado:", err)
		return "{}"
	}
	return string(jsonResultado)
}

func ProbabilidadeTurn(entrada EntradaProbabilidade) string {
	resultado := ResultadoProbabilidades{
		RoyalFlush:    calculaProbabilidadeRoyalFlush(entrada),
		StraightFlush: calculaProbabilidadeStraightFlush(entrada),
		FourOfAKind:   calculaProbabilidadeFourOfAKind(entrada),
		FullHouse:     calculaProbabilidadeFullHouse(entrada),
		Flush:         calculaProbabilidadeFlush(entrada),
		Straight:      calculaProbabilidadeStraight(entrada),
		ThreeOfAKind:  calculaProbabilidadeThreeOfAKind(entrada),
		TwoPair:       calculaProbabilidadeTwoPair(entrada),
		Pair:          calculaProbabilidadePair(entrada),
	}

	// Serializa o resultado para JSON
	jsonResultado, err := json.Marshal(resultado)
	if err != nil {
		// Tratar erro
		return "{}"
	}
	return string(jsonResultado)
}

func ProbabilidadeRiver(entrada EntradaProbabilidade) string {
	resultado := ResultadoProbabilidades{
		RoyalFlush:    calculaProbabilidadeRoyalFlush(entrada),
		StraightFlush: calculaProbabilidadeStraightFlush(entrada),
		FourOfAKind:   calculaProbabilidadeFourOfAKind(entrada),
		FullHouse:     calculaProbabilidadeFullHouse(entrada),
		Flush:         calculaProbabilidadeFlush(entrada),
		Straight:      calculaProbabilidadeStraight(entrada),
		ThreeOfAKind:  calculaProbabilidadeThreeOfAKind(entrada),
		TwoPair:       calculaProbabilidadeTwoPair(entrada),
		Pair:          calculaProbabilidadePair(entrada),
	}

	// Serializa o resultado para JSON
	jsonResultado, err := json.Marshal(resultado)
	if err != nil {
		// Tratar erro
		return "{}"
	}
	return string(jsonResultado)
}

func calculaProbabilidadePair(entrada EntradaProbabilidade) float64 {
	// Contador para as cartas
	valoresVistos := make(map[string]int)
	for _, carta := range entrada.Mao {
		valoresVistos[carta.Valor]++
	}
	for _, carta := range entrada.Mesa {
		valoresVistos[carta.Valor]++
	}

	// Verifica se já existe um par
	for _, contagem := range valoresVistos {
		if contagem >= 2 {
			// Se um par já existe entre a mão e o flop, a probabilidade é 100%
			return 100.0
		}
	}

	cartasRestantes := 52 - 5
	chanceFormarPar := 1 - float64(cartasRestantes-3)/float64(cartasRestantes)

	return chanceFormarPar * 100
}

func calculaProbabilidadeTwoPair(entrada EntradaProbabilidade) float64 {
	// Contador para as cartas
	valoresVistos := make(map[string]int)
	for _, carta := range append(entrada.Mao, entrada.Mesa...) {
		valoresVistos[carta.Valor]++
	}

	paresEncontrados := 0
	for _, contagem := range valoresVistos {
		if contagem == 2 {
			paresEncontrados++
		}
	}

	// Se já existem dois pares entre a mão e o flop, a probabilidade é 100%
	if paresEncontrados >= 2 {
		return 100.0
	}

	// Se existe apenas um par, considere a chance de formar outro par com as cartas restantes.
	// Este cálculo é uma simplificação e não considera especificamente quais cartas restantes podem formar o par.
	if paresEncontrados == 1 {
		cartasRestantes := 52 - 5 // 5 cartas já vistas (2 na mão + 3 no flop)
		// Simplificação: considera a chance de a próxima carta (turn) formar um par com qualquer uma das cartas restantes.
		chanceFormarOutroPar := float64(3) / float64(cartasRestantes)
		return chanceFormarOutroPar * 100
	}

	// Se não há pares, a chance é consideravelmente mais baixa e requer um cálculo mais detalhado.
	// Este retorno é apenas um placeholder e não reflete cálculos precisos.
	return 0.0
}

// calculaProbabilidadeThreeOfAKind calcula a probabilidade simplificada de formar um "Three of a Kind".
func calculaProbabilidadeThreeOfAKind(entrada EntradaProbabilidade) float64 {
	valoresVistos := make(map[string]int)
	for _, carta := range append(entrada.Mao, entrada.Mesa...) {
		valoresVistos[carta.Valor]++
	}

	// Verifica se já existe uma trinca
	for _, contagem := range valoresVistos {
		if contagem == 3 {
			// Se uma trinca já existe, a probabilidade é 100%
			return 100.0
		}
	}

	// Se existe um par, considere a chance de formar uma trinca com as cartas restantes.
	for _, contagem := range valoresVistos {
		if contagem == 2 {
			cartasRestantes := 52 - len(entrada.Mao) - len(entrada.Mesa)
			// Considera a chance de a próxima carta (turn) ser uma das duas que podem formar uma trinca.
			chanceFormarTrinca := float64(2) / float64(cartasRestantes)
			return chanceFormarTrinca * 100
		}
	}

	// Se não há pares, a probabilidade de formar uma trinca diretamente é ainda mais complexa
	// e requer considerar duas cartas adicionais do mesmo valor nas próximas duas ruas (turn e river),
	// o que é bastante improvável e não será calculado aqui de forma simplificada.
	return 0.0
}

func calculaProbabilidadeStraight(entrada EntradaProbabilidade) float64 {
	// Converter cartas para valores numéricos para facilitar a verificação de sequências.
	valores := map[string]int{
		"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7,
		"8": 8, "9": 9, "10": 10, "J": 11, "Q": 12, "K": 13, "A": 14,
	}
	// Contar a ocorrência de cada valor numérico nas cartas da mão + flop.
	valorOcorrencias := make(map[int]int)
	for _, carta := range append(entrada.Mao, entrada.Mesa...) {
		valorOcorrencias[valores[carta.Valor]]++
	}

	// Verificar a maior sequência possível formada até agora.
	maxSequencia := 0
	sequenciaAtual := 0
	for i := 2; i <= 14; i++ { // Inclui a possibilidade de sequência A-2-3-4-5 tratando o Ás como '1'.
		if _, ok := valorOcorrencias[i]; ok {
			sequenciaAtual++
			if sequenciaAtual > maxSequencia {
				maxSequencia = sequenciaAtual
			}
		} else {
			sequenciaAtual = 0
		}
	}

	// Se já existe um straight, a probabilidade é 100%.
	if maxSequencia >= 5 {
		return 100.0
	}

	// A chance de completar um straight depende das cartas específicas necessárias
	// e das cartas restantes no baralho. Esta é uma simplificação.
	if maxSequencia == 4 {
		// Considerando que faltam 2 cartas para serem reveladas (turn e river),
		// e simplificando a chance de uma dessas completar um straight.
		return 10.0 // Valor exemplo, não baseado em cálculos reais.
	}

	return 0.0 // Se não está próximo de um straight, a chance é considerada nula nesta simplificação.
}

func calculaProbabilidadeFlush(entrada EntradaProbabilidade) float64 {
	// Contador para os naipes
	naipesContados := make(map[string]int)
	for _, carta := range append(entrada.Mao, entrada.Mesa...) {
		naipesContados[carta.Naipe]++
	}

	// Verifica o número máximo de cartas de um mesmo naipe
	maxNaipe := 0
	for _, contagem := range naipesContados {
		if contagem > maxNaipe {
			maxNaipe = contagem
		}
	}

	// Se já existem 5 ou mais cartas do mesmo naipe, a probabilidade é 100%
	if maxNaipe >= 5 {
		return 100.0
	}

	// Para simplificar, vamos considerar que faltam 2 cartas para serem reveladas (turn e river)
	// e que você precisa de (5 - maxNaipe) cartas do mesmo naipe para completar o Flush.
	cartasNecessarias := 5 - maxNaipe

	// Se precisar de mais cartas do que as que estão para ser reveladas, a probabilidade é 0%
	if cartasNecessarias > 2 {
		return 0.0
	}

	// Este cálculo simplificado não reflete as probabilidades reais, que seriam mais complexas e
	// dependeriam das cartas específicas restantes no deck.
	// Como exemplo, vamos retornar uma probabilidade simplificada baseada no número de cartas necessárias.
	switch cartasNecessarias {
	case 1:
		// Considerando uma chance simplificada de pegar uma carta necessária no turn ou river
		return 19.6 // Exemplo: chance aproximada de pegar uma carta específica de um naipe no turn ou river
	case 2:
		// Considerando uma chance muito simplificada e não precisa para dois draws
		return 4.0 // Valor totalmente exemplificativo
	default:
		return 0.0
	}
}

func calculaProbabilidadeFullHouse(entrada EntradaProbabilidade) float64 {
	valoresVistos := make(map[string]int)
	for _, carta := range append(entrada.Mao, entrada.Mesa...) {
		valoresVistos[carta.Valor]++
	}

	// Verifica a presença de trinca e pares.
	temTrinca := false
	temPar := false
	segundoPar := false // Para o caso de dois pares.

	for _, contagem := range valoresVistos {
		if contagem == 3 {
			temTrinca = true
		} else if contagem == 2 {
			if temPar {
				segundoPar = true // Já tinha um par, agora tem dois.
			} else {
				temPar = true
			}
		}
	}

	// cartasRestantes := 52 - len(entrada.Mao) - len(entrada.Mesa)

	// Caso já tenha uma trinca, calcula a chance de pegar um par.
	if temTrinca && !temPar {
		// Chance simplificada de formar um par nas próximas duas cartas.
		return 20.0 // Valor exemplificativo.
	}

	// Caso tenha dois pares, a chance de uma das próximas cartas formar uma trinca.
	if segundoPar {
		return 4.3 // Chance de qualquer carta das 47 restantes transformar um par em trinca.
	}

	// Caso tenha apenas um par, a chance é mais complexa pois precisa formar uma trinca e outro par.
	if temPar {
		// Considerando a chance simplificada de formar uma trinca no turn e um par no river.
		return 2.0 // Valor completamente exemplificativo.
	}

	return 0.0 // Se não tem nenhum par ou trinca, a chance de formar um Full House é mínima.
}

func calculaProbabilidadeFourOfAKind(entrada EntradaProbabilidade) float64 {
	valoresVistos := make(map[string]int)
	for _, carta := range append(entrada.Mao, entrada.Mesa...) {
		valoresVistos[carta.Valor]++
	}

	// Verifica a presença de trinca ou par.
	temTrinca := false
	temPar := false

	for _, contagem := range valoresVistos {
		if contagem == 3 {
			temTrinca = true
		} else if contagem == 2 {
			temPar = true
		}
	}

	cartasRestantes := 52 - len(entrada.Mao) - len(entrada.Mesa)

	// Se já tem uma trinca, calcula a chance de pegar a quarta carta.
	if temTrinca {
		// Existem apenas 2 cartas no baralho que poderiam completar a Quadra.
		chancePegarQuartaCarta := float64(2) / float64(cartasRestantes)
		return chancePegarQuartaCarta * 100
	}

	// Se tem um par, necessita que as próximas duas cartas (turn e river) sejam as duas específicas restantes.
	if temPar {
		// Para simplificar, essa probabilidade é muito baixa, então consideramos uma chance simplificada.
		// A chance real envolveria calcular a probabilidade de pegar uma das duas cartas específicas no turn
		// e então a última carta específica no river.
		return 0.25 // Valor altamente simplificado e não baseado em cálculos reais.
	}

	return 0.0 // Se não tem um par ou trinca, a chance de formar uma Quadra diretamente é extremamente baixa.
}

func calculaProbabilidadeStraightFlush(entrada EntradaProbabilidade) float64 {
	// Converte cartas para valores numéricos para facilitar a verificação de sequências e naipes.
	valores := map[string]int{
		"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7,
		"8": 8, "9": 9, "10": 10, "J": 11, "Q": 12, "K": 13, "A": 14,
	}

	// Mapeamento de cartas por naipe para verificar sequências dentro do mesmo naipe.
	cartasPorNaipe := make(map[string][]int)
	for _, carta := range append(entrada.Mao, entrada.Mesa...) {
		cartasPorNaipe[carta.Naipe] = append(cartasPorNaipe[carta.Naipe], valores[carta.Valor])
	}

	// Verifica se há uma sequência potencial em algum naipe.
	for _, cartas := range cartasPorNaipe {
		if len(cartas) < 3 { // Precisamos de pelo menos 3 cartas para considerar um potencial Straight Flush.
			continue
		}

		// Ordena as cartas do naipe atual para verificar sequências.
		sort.Ints(cartas)

		maxSequencia := 1
		sequenciaAtual := 1
		for i := 1; i < len(cartas); i++ {
			if cartas[i] == cartas[i-1]+1 { // Sequência continua.
				sequenciaAtual++
				if sequenciaAtual > maxSequencia {
					maxSequencia = sequenciaAtual
				}
			} else if cartas[i] != cartas[i-1] { // Sequência interrompida, mas evita contar cartas duplicadas.
				sequenciaAtual = 1
			}
		}

		if maxSequencia >= 3 { // Se temos uma sequência de pelo menos 3 cartas.
			// Esta é uma simplificação. A probabilidade real dependeria das cartas específicas restantes no baralho.
			cartasRestantes := 52 - len(entrada.Mao) - len(entrada.Mesa)
			chance := 100.0 * (float64(maxSequencia) / float64(cartasRestantes))
			return min(chance, 100.0) // Garante que a probabilidade não exceda 100%.
		}
	}

	// Se não há sequência potencial, a chance de formar um Straight Flush é extremamente baixa.
	return 0.0
}

// Função auxiliar para encontrar o mínimo entre dois valores.
func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func calculaProbabilidadeRoyalFlush(entrada EntradaProbabilidade) float64 {
	// Necessário ter ao menos duas cartas do "Royal Flush" na mão ou no flop para considerar.
	royalFlushComponents := map[string]bool{"10": true, "J": true, "Q": true, "K": true, "A": true}
	naipesContados := make(map[string]int)
	royalFlushCount := make(map[string]int)

	for _, carta := range append(entrada.Mao, entrada.Mesa...) {
		if royalFlushComponents[carta.Valor] {
			naipesContados[carta.Naipe]++
			royalFlushCount[carta.Naipe]++
		}
	}

	// Verifica se algum naipe tem 3 ou mais componentes do Royal Flush.
	for naipe, count := range royalFlushCount {
		if count >= 3 {
			// Considerando cartas restantes e a especificidade do Royal Flush, a chance ainda é muito baixa.
			// Simplificando drasticamente, se você já tem 3 das 5 cartas necessárias em sequência,
			// e todas do mesmo naipe, você tem duas chances (turn e river) para conseguir as duas faltantes.
			cartasRestantes := 52 - len(entrada.Mao) - len(entrada.Mesa)
			if naipesContados[naipe] == 3 {
				// Chance simplificada para obter as duas cartas específicas restantes no turn e river.
				return (2.0 / float64(cartasRestantes)) * (1.0 / float64(cartasRestantes-1)) * 100
			} else if naipesContados[naipe] == 4 {
				// Se já tem 4, precisa de apenas uma carta específica restante.
				return (1.0 / float64(cartasRestantes)) * 100
			}
			// Nota: Este cálculo é extremamente simplificado e não reflete a probabilidade real.
		}
	}

	// Se não está perto de um Royal Flush, a chance é essencialmente zero.
	return 0.0
}
