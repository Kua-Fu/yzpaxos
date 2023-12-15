package yzpaxos

import "net/rpc"

type MsgArgs struct {
	Number int         // 提案编号
	Value  interface{} // 提案值
	From   int         // 发送者id
	To     int         // 接受者id

}

type MsgReply struct {
	Ok     bool
	Number int
	Value  interface{}
}

func call(srv string, name string, args interface{}, reply interface{}) bool {
	c, err := rpc.Dial("tcp", srv)
	if err != nil {
		return false
	}
	defer c.Close()

	err = c.Call(name, args, reply)
	if err == nil {
		return true
	}
	return false
}
