// Package bridge 桥接模式
// 定义：将抽象和实现解耦，让它们可以独立变化
//      一个类存在两个（或多个）独立变化的维度，我们通过组合的方式，让这两个（或多个）维度可以独立进行扩展
// 例子：监控告警，有不同的告警类别，有不同的通知类型
//      将通知类型和告警类别进行拆分成两个类，将通知类型作为参数传递给通知类即可
package bridge

import "testing"

// IMsgSender 消息接口
type IMsgSender interface {
	Send(msg string) error
}

// EmailMsgSend 发送邮件
// 可能还有电话、短信等各种实现
type EmailMsgSender struct {
	emails []string
}

func NewEmailMsgSender(emails []string) *EmailMsgSender {
	return &EmailMsgSender{emails: emails}
}

func (s *EmailMsgSender) Send(msg string) error {
	// Send
	return nil
}

// INotification 通知接口
type INotification interface {
	Notify(msg string) error
}

type ErrorNotification struct {
	sender IMsgSender
}

func NewErrorNotification(sender IMsgSender) *ErrorNotification {
	return &ErrorNotification{sender: sender}
}

func (r *ErrorNotification) Notify(msg string) error {
	return r.sender.Send(msg)
}

func TestErrorNotification_Notify(t *testing.T) {
	sender := NewEmailMsgSender([]string{"test@qq.com"})
	n := NewErrorNotification(sender)
	err := n.Notify("test msg")
	if err != nil {
		t.Error(err)
	}
}
