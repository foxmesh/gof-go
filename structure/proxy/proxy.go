// Package proxy 代理模式
// 效果：在不改变原始类代码的情况下，通过引入代理类来给原始类附加功能
// 应用场景：如业务系统的非功能性需求开发（监控、统计、鉴权、限流、事务...) 、RPC、缓存
// 代理模式分为静态代理和动态代理
// 静态代理：1）代理类实现和源类相同的接口，每个类都单独编写一个代理类
//         2) 一方面，我们需要在代理类中，将原始类中的所有的方法，都重新实现一遍，并且为每个方法都附加相似的代码逻辑，
//            另一方面，如果要添加的附加功能的类有不止一个，我们需要针对每个类都创建一个代理类
// 动态代理：就是我们不事先为每个原始类编写代理类，而是在运行的时候，动态地创建原始类对应的代理类，然后在系统中用代理类替换掉原始类。一般采用反射实现
package proxy

import (
	"fmt"
	"time"
)

// 静态代理
type IUser interface {
	Login(username, password string) error
}

type User struct {
}

func (u *User) Login(username, password string) error {
	return nil
}

type UserProxy struct {
	user *User
}

func NewUserProxy(user *User) *UserProxy {
	return &UserProxy{user: user}
}

func (p *UserProxy) Login(username, password string) error {
	// 在原来的业务上增加新的逻辑，并表新逻辑与原来的业务没有关系
	start := time.Now()

	if err := p.user.Login(username, password); err != nil {
		return err
	}
	fmt.Printf("user login span time: %s", time.Now().Sub(start))
	return nil
}

// 动态代理 https://lailin.xyz/post/proxy.html
