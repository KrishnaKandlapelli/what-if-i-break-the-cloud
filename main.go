package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// NodeRPC wraps our Node struct to expose methods over the network
//
type NodeRPC struct {
	ServerNode *Node
}

// RequestVote handles incoming vote requests 
//
func (n *NodeRPC) RequestVote(args *VoteRequestArgs, reply *VoteResponseReply) error {
	fmt.Printf("[Node %d] Received VoteRequest from Node %d for Term %d\n", 
		n.ServerNode.NodeID, args.CandidateID, args.Term)

	// Stub response
	//
	reply.VoterID = n.ServerNode.NodeID
	reply.Term = n.ServerNode.CurrentTerm
	reply.Granted = true 

	return nil
}

func main() {
	nodeID := 1
	port := 8001

	// Initialising the node 
	//
	myNode := &Node{
		NodeID:        nodeID,
		Nodes:         []int{1, 2, 3},
		CurrentTerm:   0,
		VotedFor:      nil,
		Log:           make([]LogEntry, 0),
		CommitLength:  0,
		CurrentRole:   "follower",
		CurrentLeader: nil,
		VotesReceived: make(map[int]bool),
		SentLength:    make(map[int]int),
		AckedLength:   make(map[int]int),
	}

	// Register RPC server
	//
	nodeRPC := &NodeRPC{ServerNode: myNode}
	rpc.Register(nodeRPC)
	rpc.HandleHTTP()

	addr := fmt.Sprintf("localhost:%d", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to bind to port %d: %v", port, err)
	}

	fmt.Printf("[Node %d] Initialised as %s. RPC Server listening on %s\n", 
		myNode.NodeID, myNode.CurrentRole, addr)

	if err := http.Serve(listener, nil); err != nil {
		log.Fatalf("server error: %v", err)
	}
}