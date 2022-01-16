// Package chain 职责链模式
// 定义：将请求的发送和接收解耦，让多个接收对象都有机会处理这个请求。将这些接收对象串成一条链
// 		 并沿着这条链传递这个请求，直到链上的某个接收对象能够处理它为止（也可以每个对象都处理）。
// 场景：复用的扩展。在实际的项目开发中比较常见，特别是框架的开发中，我们可以利用它们来提供框架的
//       扩展点。能够让架构使用者在不修改框架源码的情况下，基于扩展点定制化框架的功能
// 实现方式：链表或数组

package chain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 例子： 假设我们现在有个校园论坛，由于社区规章制度、广告、法律法规的原因需要对用户的发言进行敏感词过滤
// 如果被判定为敏感词，那么这篇帖子将会被封禁

type SensitiveWordFilter interface {
	Filter(content string) bool
}

type SensitiveWordFilterChain struct {
	filters []SensitiveWordFilter
}

func (c *SensitiveWordFilterChain) AddFilter(filter SensitiveWordFilter)  {
	c.filters = append(c.filters, filter)
}

func (c *SensitiveWordFilterChain) Filter(content string) bool  {
	for _,filter:=range c.filters{
		if filter.Filter(content){
			return true
		}
	}
	return false
}

type AdSensitiveWordFilter struct {
}

func (f *AdSensitiveWordFilter) Filter(content string) bool  {
	// 过滤算法
	return false
}


// PoliticalWordFilter 政治敏感
type PoliticalWordFilter struct{}

// Filter 实现过滤算法
func (f *PoliticalWordFilter) Filter(content string) bool {
	// TODO: 实现算法
	return true
}

func TestSensitiveWordFilterChain_Filter(t *testing.T) {
	chain := &SensitiveWordFilterChain{}
	chain.AddFilter(&AdSensitiveWordFilter{})
	assert.Equal(t, false, chain.Filter("test"))

	chain.AddFilter(&PoliticalWordFilter{})
	assert.Equal(t, true, chain.Filter("test"))
}