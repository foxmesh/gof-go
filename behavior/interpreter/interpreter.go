// Package interpreter 解析器模式
// 定义：解释器模式为某个语言定义它的语法（或叫文法）表示，并定义一个解释器用来处理这个语法
// 实现：将语法解析的工作拆分到各个小类中，以此来避免大而全的解析类
//      将语法规则分成一些小的独立单元，然后对每个单元进行解析，最终合并为对整个语法规则的解析
package interpreter

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

// 例子：现在需要实现一个告警模块，可以根据输入的告警规则来决定是否触发告警
// 		告警规则支持 &&、>、< 3种运算符
// 		其中 >、< 优先级比  && 更高

// IExpression 表达式接口
type IExpression interface {
	Interpret(stats map[string]float64) bool
}

// AlertRule 告警规则
type AlertRule struct {
	expression IExpression
}

func NewAlertRule(rule string) (*AlertRule, error) {
	exp, err := NewAndExpression(rule)
	return &AlertRule{expression: exp}, err
}

func (r AlertRule) Interpret(stats map[string]float64) bool {
	return r.expression.Interpret(stats)
}

// GreaterExpression > 表达式
type GreaterExpression struct {
	key   string
	value float64
}

func (e *GreaterExpression) Interpret(stats map[string]float64) bool {
	v, ok := stats[e.key]
	if !ok {
		return false
	}
	return v > e.value
}

func NewGreaterExpression(exp string) (*GreaterExpression, error) {
	data := regexp.MustCompile(`\s+`).Split(strings.TrimSpace(exp), -1)
	if len(data) != 3 || data[1] != ">" {
		return nil, fmt.Errorf("exp is invalid: %s", exp)
	}
	val, err := strconv.ParseFloat(data[2], 10)
	if err != nil {
		return nil, fmt.Errorf("exp is invalid: %s", exp)
	}

	return &GreaterExpression{
		key:   data[0],
		value: val,
	}, nil
}

type LessExpression struct {
	key   string
	value float64
}

func (e LessExpression) Interpret(stats map[string]float64) bool {
	v, ok := stats[e.key]
	if !ok {
		return false
	}
	return v < e.value
}

func NewLessExpression(exp string) (*LessExpression, error) {
	data := regexp.MustCompile(`\s+`).Split(strings.TrimSpace(exp), -1)
	if len(data) != 3 || data[1] != "<" {
		return nil, fmt.Errorf("exp is invalid: %s", exp)
	}
	val, err := strconv.ParseFloat(data[2], 10)
	if err != nil {
		return nil, fmt.Errorf("exp is invalid: %s", exp)
	}
	return &LessExpression{
		key:   data[0],
		value: val,
	}, nil
}

type AndExpression struct {
	expressions []IExpression
}

func (e AndExpression) Interpret(stats map[string]float64) bool {
	for _, exp := range e.expressions {
		if !exp.Interpret(stats) {
			return false
		}
	}
	return true
}

func NewAndExpression(exp string) (*AndExpression, error) {
	exps := strings.Split(exp, "&&")
	expressions := make([]IExpression, len(exps))

	for i, e := range exps {
		var expression IExpression
		var err error

		switch {
		case strings.Contains(e, ">"):
			expression, err = NewGreaterExpression(e)
		case strings.Contains(e, "<"):
			expression, err = NewLessExpression(e)
		default:
			err = fmt.Errorf("exp is invalid:%s", exp)
		}

		if err != nil {
			return nil, err
		}
		expressions[i] = expression
	}
	return &AndExpression{expressions: expressions}, nil
}

func TestAlertRule_Interpret(t *testing.T) {
	stats := map[string]float64{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	tests := []struct {
		name  string
		stats map[string]float64
		rule  string
		want  bool
	}{
		{
			name:  "case1",
			stats: stats,
			rule:  "a > 1 && b > 10 && c < 5",
			want:  false,
		},
		{
			name:  "case2",
			stats: stats,
			rule:  "a < 2 && b > 10 && c < 5",
			want:  false,
		},
		{
			name:  "case3",
			stats: stats,
			rule:  "a < 5 && b > 1 && c < 10",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := NewAlertRule(tt.rule)
			require.NoError(t, err)
			assert.Equal(t, tt.want, r.Interpret(tt.stats))
		})
	}
}
