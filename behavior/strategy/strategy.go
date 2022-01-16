// Package strategy 策略模式
// 定义：定义一族算法类，将每个算法分别封装起来，让它们可以互相替换
//       策略模式可以使算法的变化独立于使用者
// 解耦：工厂模式 -  解耦对象的创建和使用
//       观察者模式 -  解耦观察者和被观察者
//       策略模式 - 解耦的是策略的定义、创建、使用这三部分
//
package strategy

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

// 例子：我们在保存文件的时候，由于政策或者其他的原因可能需要选择不同的存储方式，敏感数据我们需要加密存储，不敏感的数据我们可以直接明文保存

// 定义策略接口
type StorageStrategy interface {
	Save(name string, data []byte) error
}

var strategys = map[string]StorageStrategy{
	"file":&fileStorage{},
	"encrypt_file":&encryptFileStorage{},
}

func NewStorageStrategy(t string)(StorageStrategy,error){
	s, ok := strategys[t]
	if !ok{
		return nil, fmt.Errorf("not found StorageStrategy: %s", t)
	}
	return s, nil
}

// 保存到文件
type fileStorage struct {
}

func (s fileStorage) Save(name string, data []byte) error {
	return ioutil.WriteFile(name, data, os.ModeAppend)
}

// 加密保存
type encryptFileStorage struct {
}

func (s encryptFileStorage) Save(name string, data []byte)  error {
	date,err:=encrypt(data)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(name,date, os.ModeAppend)
}

func encrypt(data []byte)([]byte, error)  {
	// 加密
	return data, nil
}

func Test_demo(t *testing.T) {
	// 假设这里获取数据，以及数据是否敏感
	data, sensitive := getData()
	strategyType := "file"
	if sensitive {
		strategyType = "encrypt_file"
	}

	storage, err := NewStorageStrategy(strategyType)
	assert.NoError(t, err)
	assert.NoError(t, storage.Save("./test.txt", data))
}

// getData 获取数据的方法
// 返回数据，以及数据是否敏感
func getData() ([]byte, bool) {
	return []byte("test data"), false
}