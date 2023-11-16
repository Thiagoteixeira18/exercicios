package exercicios

import "fmt"

// NewVoteCounter cria um contador de votos e retorna um ponteiro para ele.
func NewVoteCounter(initialVotes int) *int {
	counter := initialVotes
	return &counter
}

// VoteCount retorna o número de votos no contador.
// Se o contador for nil, assume-se que não há votos.
func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	}
	return *counter
}

// IncrementVoteCount incrementa o contador de votos pelo valor especificado.
// Supõe-se que o ponteiro passado nunca seja nil.
func IncrementVoteCount(counter *int, increment int) {
	*counter += increment
}

// ElectionResult representa o resultado de uma eleição.
type ElectionResult struct {
	Name  string
	Votes int
}

// NewElectionResult cria um novo resultado de eleição para um candidato com um número especificado de votos.
func NewElectionResult(name string, votes int) *ElectionResult {
	return &ElectionResult{Name: name, Votes: votes}
}

// DisplayResult gera uma string para exibir o resultado da eleição de um candidato.
func DisplayResult(result *ElectionResult) string {
	return fmt.Sprintf("%s (%d)", result.Name, result.Votes)
}

// DecrementVotesOfCandidate decrementa o número de votos para um candidato específico nos resultados finais.
func DecrementVotesOfCandidate(finalResults map[string]int, candidate string) {
	if votes, exists := finalResults[candidate]; exists {
		finalResults[candidate] = votes - 1
	}
}

func main10() {
	// Exemplos de uso das funções:

	// Criar um contador de votos
	var initialVotes int = 2
	var counter *int
	counter = NewVoteCounter(initialVotes)
	fmt.Println("*counter == initialVotes:", *counter == initialVotes)

	// Contar votos
	var votes int = 3
	var voteCounter *int
	voteCounter = &votes

	voteCount := VoteCount(voteCounter)
	fmt.Println("VoteCount(voteCounter):", voteCount)

	var nilVoteCounter *int
	nilVoteCount := VoteCount(nilVoteCounter)
	fmt.Println("VoteCount(nilVoteCounter):", nilVoteCount)

	// Incrementar votos
	var voteCounter2 *int
	voteCounter2 = &votes

	IncrementVoteCount(voteCounter2, 2)

	fmt.Println("votes == 5:", votes == 5)
	fmt.Println("*voteCounter2 == 5:", *voteCounter2 == 5)

	// Criar resultado de eleição
	var result *ElectionResult
	result = NewElectionResult("Peter", 3)

	fmt.Println("result.Name == 'Peter':", result.Name == "Peter")
	fmt.Println("result.Votes == 3:", result.Votes == 3)

	// Exibir resultado da eleição
	resultDisplay := DisplayResult(result)
	fmt.Println("DisplayResult(result):", resultDisplay)

	// Decrementar votos de um candidato nos resultados finais
	finalResults := map[string]int{
		"Mary": 10,
		"John": 51,
	}

	DecrementVotesOfCandidate(finalResults, "Mary")

	fmt.Println("finalResults['Mary'] == 9:", finalResults["Mary"] == 9)
}