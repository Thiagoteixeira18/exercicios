package exercicios

import "fmt"

// ElectionResult defines the structure for the result of an election.
type ElectionResult struct {
	Name  string
	Votes int
}

// NewVoteCounter creates a vote counter and returns a pointer to it.
func NewVoteCounter(initialVotes int) *int {
	counter := initialVotes
	return &counter
}

// VoteCount returns the number of votes in the counter.
// If the counter is nil, it is assumed that there are no votes.
func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	}
	return *counter
}

// IncrementVoteCount increments the vote counter by the specified value.
// It assumes the passed pointer is never nil.
func IncrementVoteCount(counter *int, increment int) {
	*counter += increment
}

// NewElectionResult creates a new election result for a candidate with a specified number of votes.
func NewElectionResult(name string, votes int) *ElectionResult {
	return &ElectionResult{Name: name, Votes: votes}
}

// DisplayResult generates a string to display the result of a candidate's election.
func DisplayResult(result *ElectionResult) string {
	return fmt.Sprintf("%s (%d)", result.Name, result.Votes)
}

// DecrementVotesOfCandidate decrements the number of votes for a specific candidate in the final results.
func DecrementVotesOfCandidate(finalResults map[string]int, candidate string) {
	if votes, exists := finalResults[candidate]; exists {
		finalResults[candidate] = votes - 1
	}
}

func main80() {
	// Examples of using the functions:

	// Create a vote counter
	var initialVotes int = 2
	var counter *int
	counter = NewVoteCounter(initialVotes)
	fmt.Println("*counter == initialVotes:", *counter == initialVotes)

	// Count votes
	var votes int = 3
	var voteCounter *int
	voteCounter = &votes

	voteCount := VoteCount(voteCounter)
	fmt.Println("VoteCount(voteCounter):", voteCount)

	var nilVoteCounter *int
	nilVoteCount := VoteCount(nilVoteCounter)
	fmt.Println("VoteCount(nilVoteCounter):", nilVoteCount)

	// Increment votes
	var voteCounter2 *int
	voteCounter2 = &votes

	IncrementVoteCount(voteCounter2, 2)

	fmt.Println("votes == 5:", votes == 5)
	fmt.Println("*voteCounter2 == 5:", *voteCounter2 == 5)

	// Create an election result
	var result *ElectionResult
	result = NewElectionResult("Peter", 3)

	fmt.Println("result.Name == 'Peter':", result.Name == "Peter")
	fmt.Println("result.Votes == 3:", result.Votes == 3)

	// Display election result
	resultDisplay := DisplayResult(result)
	fmt.Println("DisplayResult(result):", resultDisplay)

	// Decrement votes of a candidate in the final results
	finalResults := map[string]int{
		"Mary": 10,
		"John": 51,
	}

	DecrementVotesOfCandidate(finalResults, "Mary")

	fmt.Println("finalResults['Mary'] == 9:", finalResults["Mary"] == 9)
}