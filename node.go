// Node representation

package main

// LogEntry represents a single message in the log w/ term
//
type LogEntry struct {
	Msg string
	Term int
} 

// Node represents a single consensus server instance
//
type Node struct {
	NodeID int
	Nodes  []int // IDs of all nodes in the system

	// Persistent State - stable storage/disk
	//
	CurrentTerm  int
	VotedFor     *int 
	Log          []LogEntry
	CommitLength int

	// Volatile State - memory
	//
	CurrentRole   string // "follower" / "candidate" / "leader"
	CurrentLeader *int
	VotesReceived map[int]bool
	SentLength    map[int]int
	AckedLength   map[int]int
}

//RPC arguements and reply struct
//
type VoteRequestArgs struct {
	CandidateID int
	Term        int
	LogLength   int
	LastLogTerm int
}

type VoteResponseReply struct {
	VoterID int
	Term    int
	Granted bool
}