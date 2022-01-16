// Package composite 组合模式
// 定义：将一组对象组织成树形结构，以表示一种“部分-整体”的层次结构。组合让使用方可以统一单个对象和组合对象的处理逻辑
// 应用场景：1）业务场景必须能够表示成树形结构
//         2) 组合模式，将一组对象组织成树形结构，将单个对象和组合对象都看做树中的节点，以统一处理逻辑，并且它利用树形结构的特点，
//            递归地处理每个子树，依次简化代码实现
// 例子：部门和员工关系； 文件夹和文件
package composite

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 例子：公司的人员组织就是一个典型的树状结构，现在假设我们现在有部门和员工两种角色，一个部门下面可以存在子部门和员工，员工下面不能再包含其他节点
//      例子要实现一个统计一个部门下员工数量的功能

type IOrganization interface {
	Count() int
}

type Employee struct {
	Name string
}

func (e Employee) Count() int {
	return 1
}

type Department struct {
	Name             string
	SubOrganizations []IOrganization
}

func (d Department) Count() int {
	c := 0
	for _, org := range d.SubOrganizations {
		c += org.Count()
	}
	return c
}

func (d *Department) AddSub(org IOrganization) {
	d.SubOrganizations = append(d.SubOrganizations, org)
}

func NewOrganization() IOrganization {
	root := &Department{Name: "root"}
	for i := 0; i < 10; i++ {
		root.AddSub(&Employee{})
		root.AddSub(&Department{
			Name:             "sub",
			SubOrganizations: []IOrganization{&Employee{}},
		})
	}
	return root
}

func TestNewOrganization(t *testing.T) {
	got := NewOrganization().Count()
	assert.Equal(t, 20, got)
}
