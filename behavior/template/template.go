// Package template 模板方法模式
// 定义: 在一个方法中定义了一个算法骨架，并将某些步骤推迟到子类中实现。
//		模板方法模式可以让子类在不改变算法的整体结构的情况下，重新定义算法中的某些步骤
// 应用场景：1）扩展 -  框架通过模板模式提供功能扩展点，让框架用户可以在不个性框架源码的情况下，
//               		基于扩展点，定制化框架的功能
// 			2) 复用 - 复用指的是，所有的子类可以复用父类中提供的模板方法的代码

package template

import (
	"errors"
	"fmt"
)

// 实现一个短信推送的系统的例子
// 检查短信字数是否超过限制
// 检查手机号是否正确
// 发送短信
// 返回状态

type ISMS interface{
	send(content string, phone int) error
}

// 短信基类
type sms struct {
	ISMS
}

func (s sms) Valid(content string) error {
	if len(content)>63{
		return errors.New("短信超长")
	}
	return nil
}

func (s *sms)Send(content string, phone int) error {
	if err:=s.Valid(content); err != nil {
		return err
	}

	return s.send(content,phone)
}

// TelecomSMS 走电信通道
type TelecomSMS struct {
	*sms
}

func NewTelecomSMS() *TelecomSMS  {
	 tel := &TelecomSMS{}
	// 这里有点绕，是因为 go 没有继承，用嵌套结构体的方法进行模拟
	// 这里将子类作为接口嵌入父类，就可以让父类的模板方法 Send 调用到子类的函数
	// 实际使用中，我们并不会这么写，都是采用组合+接口的方式完成类似的功能
	 tel.sms = &sms{ISMS:tel}
	 return tel
}

func (tel *TelecomSMS) send(content string, phone int) error{
	fmt.Println("send by telecom success")
	return nil
}

