package raft

type AppendEntriesArgs struct {
	Term         int
	LeaderId     int
	PrevLogIndex int
	PrevLogTerm  int

	// 需要复制的日志条目，用于发送心跳消息时候，Entries 为空
	Entries []LogEntry
	// 领导者已提交的最大日志索引，用于跟随者提交
	LeaderCommit int
}

type AppendEntriesReply struct {
	Term    int
	Success bool
}

func (rf *Raft) AppendEntries(args *AppendEntriesArgs, reply *AppendEntriesReply) {

	rf.mu.Lock()
	defer rf.mu.Unlock()

	reply.Term = rf.currentTerm
	reply.Success = false

	if args.Term < rf.currentTerm {
		return
	}

	if args.Term > rf.currentTerm {
		reply.Term = args.Term
		rf.currentTerm = args.Term
	}

	// 主要为了重置选举超时时间
	rf.setState(Follower)

	// 日志一致性检查
	// lastLogIndex := rf.getLastIndex()
	lastLogIndex := len(rf.log) - 1

	// （1）如果leader发送的日志前一条的 index > 当前状态机最后一条的index
	// （2）如果当前状态机的前一条日志的任期 != leader发送的日志的前一条的任期
	if args.PrevLogIndex > rf.log[lastLogIndex].Index || rf.log[args.PrevLogIndex].Term != args.PrevLogTerm {
		return
	}

	// args.PrevLogIndex <= rf.log[lastLogIndex].Index &&
	// rf.log[args.PrevLogIndex].Term == args.PrevLogTerm

	reply.Success = true

	// 需要处理重复的RPC请求
	// 比较日志条目的任期，已确认是否能够安全的追加日志
	// 否则会导致重复应用命令

	index := args.PrevLogIndex

	for i, entry := range args.Entries {

		index++
		if index < len(rf.log) {

		}
	}
}

func (rf *Raft) setState(s int) {

}
