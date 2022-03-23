package electionday

import (
	"fmt"
)

/*
NewVoteCounter returns a new vote counter with a given number of inital votes.
Accepts the number of initial votes for a candidate.
Returns a pointer referring to an int, initialized with the given number of initial votes. */
func NewVoteCounter(initialVotes int) *int {
	var voteCounter *int
	voteCounter = &initialVotes
	return voteCounter
}

/*
VoteCount extracts the number of votes from a counter.
Takes a counter (*int) as an argument and returns the number of votes in the counter.
If the counter is nil you should assume the counter has no votes. */
func VoteCount(counter *int) int {
	var votes int
	if counter == nil {
		return 0
	}
	votes = *counter
	return votes
}

/*
IncrementVoteCount increments the value in a vote counter.
Takes a counter (*int) as an argument and a number of votes,
and will increment the counter by that number of votes.
You can assume the pointer passed will never be nil. */
func IncrementVoteCount(counter *int, increment int) {
	*counter += increment
}

// NewElectionResult creates a new election result.
// Receives the name of a candidate and their number of votes and returns a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	var result *ElectionResult
	result = &ElectionResult{candidateName, votes}
	return result
}

/*
DisplayResult creates a message with the result to be displayed
The message should show the name of the new president
and the votes it had, in the following format: <candidate_name> (<votes>). */
func DisplayResult(result *ElectionResult) string {
	var message string
	message = fmt.Sprintf("%s (%d)", result.Name, result.Votes)
	return message
}

/*
DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map
Receives the final results and the name of a candidate for which you should decrement its vote count.
The final results are given in the form of a map[string]int,
where the keys are the names of the candidates and the values are its total votes. */
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	results[candidate]--
}
