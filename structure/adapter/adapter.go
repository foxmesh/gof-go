// Package adapter 适配器模式
// 定义：用来做适配，将不兼容的接口转换成可兼容的接口，让原本由于接口中不兼容而不能在一起工作的类可以一起工作
// 实现方式：1）类适配器 - 继承 ： 需要修改的接口少
//         2) 对象适配器 - 组合 ：  需要修改的接口很多
// 应用场景：1）封装有缺陷的接口 2）统一多个类的接口设计 3）替换依赖的外部关系  4）兼容老版本接口 5）适配不同格式的数据
//
// 和其它几种模式的区别：
// 代理模式 - 在不改变原类接口的条件下，为原类定义一个代理类，主要目的是控制访问，而非加强功能
// 桥接模式 - 将接口和实现分离，从而让它们可以容易且相对独立地加以改变
// 装饰器模式 - 在不改变原类接口的情况下，对原类功能进行增强，并且支持多个装饰器的嵌套使用
// 适配器模式 - 是一种事后补救策略。适配器提供原类不同的接口，而代理模式、装饰器模式提供的都是原类相同的接口
package adapter

import (
	"fmt"
	"testing"
)

type ICreateServer interface {
	CreateServer(cpu, mem float64) error
}

type AWSClient struct {
}

func (c AWSClient) RunInstance(cpu, mem float64) error {

	fmt.Printf("aws client run success, cpu： %f, mem: %f", cpu, mem)
	return nil
}

type AWSClientAdapter struct {
	client AWSClient
}

func (a *AWSClientAdapter) CreateServer(cpu, mem float64) error {
	a.client.RunInstance(cpu, mem)
	return nil
}

type AliClient struct {
}

func (c AliClient) RunInstance(cpu, mem float64) error {
	fmt.Printf("aliyun client run success, cpu： %f, mem: %f", cpu, mem)
	return nil
}

type AliClientAdapter struct {
	client AliClient
}

func (a *AliClientAdapter) CreateServer(cpu, mem float64) error {
	a.client.RunInstance(cpu, mem)
	return nil
}

func TestCreateServer(t *testing.T) {
	var a ICreateServer = &AliClientAdapter{client: AliClient{}}

	a.CreateServer(1.0, 2.0)

	var b ICreateServer = &AWSClientAdapter{client: AWSClient{}}
	b.CreateServer(1.0, 2.0)
}
