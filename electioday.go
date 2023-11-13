package exercicios

import "fmt"

// NewVoteCounter creates a new vote counter initialized with the given number of votes.
func NewVoteCounter(initialVotes int) *int {
	return &initialVotes
}

// VoteCount returns the number of votes in the counter. If the counter is nil, it returns 0.
func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	}
	return *counter
}

// IncrementVoteCount increments the counter by the given number of votes.
func IncrementVoteCount(counter *int, votes int) {
	*counter += votes
}

// DecrementVotesOfCandidate decrements the vote count of the specified candidate in the final results map.
func DecrementVotesOfCandidate(finalResults map[string]int, candidateName string) {
	if votes, exists := finalResults[candidateName]; exists {
		finalResults[candidateName] = votes - 1
	}
}

// ElectionResult represents the result of an election for a candidate.
type ElectionResult struct {
	Name  string
	Votes int
}

// NewElectionResult creates a new election result with the given candidate name and votes.
func NewElectionResult(name string, votes int) *ElectionResult {
	return &ElectionResult{Name: name, Votes: votes}
}

// DisplayResult returns a string with the message to display for the election result.
func DisplayResult(result *ElectionResult) string {
	return fmt.Sprintf("%s (%d)", result.Name, result.Votes)
}