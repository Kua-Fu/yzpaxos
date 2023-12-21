package raft

type LogEntry struct {
	Index   int
	Term    int
	Command interface{}
}

// 通过Index 和 term唯一标识一个日志条目
// （1） 如果两个节点的日志在相同索引位置上的任期号相同，则认为它们具有一样的命令，
// 并且从日志开头到这个索引位置之间的日志也完全相同
// （2） 如果给定的记录已经提交，那么所有前面的记录也提交
