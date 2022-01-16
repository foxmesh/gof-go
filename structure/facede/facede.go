// Package facade 外观模式（门面模式）
// 定义： 为子系统提供一组统一的接口，定义一组高层接口让子系统更易用
// 应用场景 ：1）解决易用性问题 - 用来封装系统的底层实现，隐藏系统的复杂性，提供一组更加简单易用、更高层的接口
// 			2) 解决性能问题 - 减少网络请求
// 			3) 解决分布式事务问题 -  不用多次调用了，一次调用就可以在同一个进程的事务中解决
//
// 例子：假设现在我有一个网站，以前有登录和注册的流程，登录的时候调用用户的查询接口，注册时调用用户的创建接口。
// 为了简化用户的使用流程，我们现在提供直接验证码登录/注册的功能，如果该手机号已注册那么我们就走登录流程，如果该手机号未注册，那么我们就创建一个新的用户。
package facade

import (
	"testing"
)

type User struct {
	Name string
}

// IUserFacade 提供外观
type IUserFacade interface {
	LoginOrRegister(phone int, code int) (*User, error)
}

type UserService struct {
}

func (s UserService) Login(phone int, code int) (*User, error) {
	// 校验 ...
	return &User{Name: "test login"}, nil
}

func (s UserService) Register(phone int, code int) (*User, error) {
	// 校验 ...
	// 创建用户
	return &User{Name: "test register"}, nil
}

func (s UserService) LoginOrRegister(phone int, code int) (*User, error) {
	user, err := s.Login(phone, code)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return user, nil
	}
	return s.Register(phone, code)
}

func TestUserService_Login(t *testing.T) {
	service := UserService{}
	user, err := service.Login(13001010101, 1234)
	t.Log(user, err)
}

func TestUserService_LoginOrRegister(t *testing.T) {
	service := UserService{}
	user, err := service.LoginOrRegister(13001010101, 1234)
	t.Log(user, err)
}

func TestUserService_Register(t *testing.T) {
	service := UserService{}
	user, err := service.Register(13001010101, 1234)
	t.Log(user, err)
}
