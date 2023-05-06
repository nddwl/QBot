package matcher

import (
	"fmt"
	zero "github.com/wdvxdr1123/ZeroBot"
)

type Matcher map[string]*zero.Matcher

func New() *Matcher {
	return &Matcher{}
}

func (t *Matcher) Add(prefix string, matcher *zero.Matcher) *zero.Matcher {
	if _, ok := (*t)[prefix]; ok {
		panic(fmt.Sprintf("已经存在相同名字的方法"))
	}
	(*t)[prefix] = matcher
	return matcher
}

// Get 获取已存在的匹配
func (t *Matcher) Get(name string) *zero.Matcher {
	if m, ok := (*t)[name]; !ok {
		panic(fmt.Sprintf("不存在此方法"))
	} else {
		return m
	}
}

// FindAll 查询所有已经注册的匹配
func (t *Matcher) FindAll() (name []string) {
	name = make([]string, len(*t))
	var index int
	for k := range *t {
		name[index] = k
		index++
	}
	return
}
