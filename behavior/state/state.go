// Package state 状态模式
// 定义：状态模式通过将事件触发的状态转移和动作执行，拆分到不同的状态中，来避免分支判断逻辑
// 常用于实现状态机（游戏、工作流引擎）
//
// 有限状态机(FSM）- 组成：1）状态  2) 事件  3）动作
// 实现方式：1）分支逻辑法：当状态和逻辑比较少的时候可以用到，否则会出现大量 if/else
//   		2）查表法：对于状态很多、状态转移比较复杂的状态机来说，查表法比较适合
//			3）状态模式：对于状态不多、状态转移也比较简单，但事件触发执行的动作包含的业务逻辑
// 				可能比较复杂的状态机

package state

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 这是一个工作流的例子，在企业内部或者是学校我们经常会看到很多审批流程
// 假设我们有一个报销的流程: 员工提交报销申请 -> 直属部门领导审批 -> 财务审批 -> 结束
// 在这个审批流中，处在不同的环节就是不同的状态
// 而流程的审批、驳回就是不同的事件

// IState 状态
type IState interface {
	// 审批通过
	Approval(m *Machine)
	// 驳回
	Reject(m *Machine)
	// 获取当前状态名称
	GetName() string
}

type Machine struct {
	state IState
}

func (m *Machine) SetState(state IState)  {
	m.state = state
}

func (m *Machine) GetStateName() string {
	return m.state.GetName()
}
func (m *Machine) Approval() {
	m.state.Approval(m)
}

func (m *Machine) Reject() {
	m.state.Reject(m)
}

// leaderApproveState 直属领导审批
type leaderApproveState struct{}

func (s leaderApproveState) GetName() string  {
	return "LeaderApproveState"
}

func (s leaderApproveState) Approval(m *Machine) {
	fmt.Println("leader 审批成功")
	m.SetState(GetFinanceApproveState())
}

func (s leaderApproveState) Reject(m *Machine) {
}

func GetLeaderApproveState() IState {
	return &leaderApproveState{}
}


// financeApproveState 财务审批
type financeApproveState struct{}

// Approval 审批通过
func (f financeApproveState) Approval(m *Machine) {
	fmt.Println("财务审批成功")
	fmt.Println("出发打款操作")
}

// 拒绝
func (f financeApproveState) Reject(m *Machine) {
	m.SetState(GetLeaderApproveState())
}

// GetName 获取名字
func (f financeApproveState) GetName() string {
	return "FinanceApproveState"
}

// GetFinanceApproveState GetFinanceApproveState
func GetFinanceApproveState() IState {
	return &financeApproveState{}
}


func TestMachine_GetStateName(t *testing.T) {
	m := &Machine{state: GetLeaderApproveState()}
	assert.Equal(t, "LeaderApproveState", m.GetStateName())
	m.Approval()
	assert.Equal(t, "FinanceApproveState", m.GetStateName())
	m.Reject()
	assert.Equal(t, "LeaderApproveState", m.GetStateName())
	m.Approval()
	assert.Equal(t, "FinanceApproveState", m.GetStateName())
	m.Approval()
}