package electionday

import "strconv"

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	var pv *int;

    pv = &initialVotes;
    return pv;
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {

    if counter == nil {
        return (0);
    } else {
    	return *counter;
    }
}

// IncrementVoteCount increments the value in a vote counter
func IncrementVoteCount(counter *int, increment int) {
	*counter += increment;
}

// NewElectionResult creates a new election result
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	var result *ElectionResult;

    result = &ElectionResult{Name: candidateName, Votes: votes};
    return (result);
}

// DisplayResult creates a message with the result to be displayed
func DisplayResult(result *ElectionResult) string {
	var out string;
    var res ElectionResult;

    res = *result;
    out = res.Name + " (" + strconv.Itoa(res.Votes)+ ")";
    return out;
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
    results[candidate] -= 1;
}

