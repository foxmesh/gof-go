// Package command 命令模式
// 定义：将请求（命令）封装为一个对象，这样可以使用不同的请求参数化其他对象（将不同请求依赖注入到其他对象）
// 		并且能够支持请求（命令）的排除执行、记录日志、撤销等（附加控制）功能
// 实现：主要是把函数封装成对象，实现相同接口，使其可以通过参数传递。因为go函数可以直接当作参数传递，所以不必使用对象来实现
// 场景：用来控制命令的执行，比如，异步、延迟、排除执行命令、撤销重做命令、存储命令、给命令记录日志等
//
// 比较：命令模式 - 不同的命令包含了不同的目的和功能，不能够相互替换
//      策略模式 - 不同的策略是为了达到相同目的的不同实现，他们之间可以互相替换

package command

import (
	"fmt"
	"testing"
	"time"
)

// 实现一，使用对象

type ICommand interface {
	Execute() error
}

type StartCommand struct {
}

func NewStartCommand() *StartCommand {
	return &StartCommand{}
}

func (c *StartCommand) Execute() error {
	fmt.Println("game start")
	return nil
}

type ArchiveCommand struct {
}

func NewArchiveCommand() *ArchiveCommand {
	return &ArchiveCommand{}
}

func (c *ArchiveCommand) Execute() error {
	fmt.Println("game archive")
	return nil
}

func TestDemo1(t *testing.T) {
	// 用于测试，模拟来自客户端的事件
	eventChan := make(chan string)
	go func() {
		events := []string{"start", "archive", "start", "archive", "start", "start"}
		for _, e := range events {
			eventChan <- e
		}
	}()
	defer close(eventChan)

	// 使用命令队列缓存命令
	commands := make(chan ICommand, 1000)
	defer close(commands)

	go func() {
		for {
			// 从请求或者其他地方获取相关事件参数
			event, ok := <-eventChan
			if !ok {
				return
			}

			var command ICommand
			switch event {
			case "start":
				command = NewStartCommand()
			case "archive":
				command = NewArchiveCommand()
			}

			// 将命令入队
			commands <- command
		}
	}()

	for {
		select {
		case c := <-commands:
			c.Execute()
		case <-time.After(1 * time.Second):
			fmt.Println("timeout 1s")
			return
		}
	}
}
