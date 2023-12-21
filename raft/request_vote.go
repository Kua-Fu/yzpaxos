package raft

import "time"

type RequestVoteArgs struct {
	// 候选者任期
	Term int
	// 候选者id
	CandidateId int
}

type RequestVoteReply struct {

	// 处理请求节点的任期号，用于候选者更新自己的任期
	Term int

	// 候选者获得选票时候为 true
	// 否则为false
	VoteGranted bool
}

// 投票选举逻辑
func (rf *Raft) RequestVote(
	args *RequestVoteArgs,
	reply *RequestVoteReply) {

	rf.mu.Lock()
	defer rf.mu.Unlock()

	reply.Term = rf.currentTerm
	reply.VoteGranted = false

	if args.Term < rf.currentTerm {
		return
	}

	// 如果收到来自更大任期
	if args.Term > rf.currentTerm {
		rf.currentTerm = args.Term
		rf.state = Follower
		rf.votedFor = -1
	}

	if rf.votedFor == -1 || rf.votedFor == args.CandidateId {
		rf.votedFor = args.CandidateId
		reply.VoteGranted = true
		rf.heartbeatTime = time.Now()
	}

	return
}
