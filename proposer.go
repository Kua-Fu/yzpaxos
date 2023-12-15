package yzpaxos

import "fmt"

type Proposer struct {
	id        int   // 服务器id
	round     int   // 当前提议者已知的最大轮次
	number    int   // 提案编号
	acceptors []int // 接受者id 列表

}

func (p *Proposer) propose(v interface{}) interface{} {

	p.round++
	p.number = p.proposalNumber()

	// 第一阶段 phase 1
	prepareCount := 0
	maxNumber := 0
	for _, aid := range p.acceptors {
		args := MsgArgs{
			Number: p.number,
			From:   p.id,
			To:     aid,
		}
		reply := new(MsgReply)
		fmt.Printf("--call prepare--, args: %v, reply: %v \n", args, reply)
		err := call(fmt.Sprintf("127.0.0.1:%d", aid), "Acceptor.Prepare", args, reply)
		if !err {
			continue
		}

		if reply.Ok {
			prepareCount++
			if reply.Number > maxNumber {
				maxNumber = reply.Number
				v = reply.Value
			}
		}

		if prepareCount == p.majority() {
			break
		}
	}

	// 第二阶段 phase 2
	acceptCount := 0
	if prepareCount >= p.majority() {

		for _, aid := range p.acceptors {
			args := MsgArgs{
				Number: p.number,
				Value:  v,
				From:   p.id,
				To:     aid,
			}

			reply := new(MsgReply)
			fmt.Printf("--call accept--, args: %v, reply: %v \n", args, reply)
			ok := call(fmt.Sprintf("127.0.0.1:%d", aid), "Acceptor.Accept", args, reply)
			if !ok {
				continue
			}

			if reply.Ok {
				acceptCount++
			}

		}
	}

	if acceptCount >= p.majority() {
		// 选择的提案值
		return v
	}

	return nil
}

// 提案编号=( 轮次，服务器id)
func (p *Proposer) proposalNumber() int {
	return p.round<<16 | p.id
}

func (p *Proposer) majority() int {
	return len(p.acceptors)/2 + 1
}
