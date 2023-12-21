package raft

import (
	"sync"
	"time"
)

const (
	Follower = iota
	Condidate
	Leader
)

// RequestVoteRPC 领导者选举
// AppendEntriesRPC 领导者复制日志和发送心跳

// 共识算法
// （1）安全性 一个任期内只会有一个领导者被选举出来
// （2） 活性 最终能选出一个领导者

type Raft struct {
	mu sync.Mutex
	// 服务器当前状态
	state int

	// 服务器当前已知的最新任期
	currentTerm int

	// 当前任期内
	votedFor int

	// 心跳时间
	heartbeatTime time.Time

	// 增加成员变量，表示状态机日志
	log []LogEntry
}
