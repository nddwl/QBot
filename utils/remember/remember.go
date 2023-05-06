package remember

import (
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

type Remember map[string]Find

type Find map[string]interface{}

func New() *Remember {
	return &Remember{}
}

func (t *Remember) Get(name string) Find {
	if v, ok := (*t)[name]; ok {
		return v
	}
	panic("不存在")
}

func (t *Remember) Add(name string) Find {
	if _, ok := (*t)[name]; ok {
		panic("已存在")
	}
	(*t)[name] = Find{}
	return (*t)[name]
}

func (t *Find) Set(name string, value interface{}) {
	(*t)[name] = value
}

func (t *Find) Get(name string) (interface{}, bool) {
	if v, ok := (*t)[name]; ok {
		return v, true
	}
	return nil, false
}

func (t *Find) Forward(ctx *zero.Ctx) bool {
	messages, ok := t.Get(ctx.MessageString())
	if !ok {
		return false
	}
	mid := messages.([]int64)
	if len(mid) > 1 {
		node := make([]message.MessageSegment, len(mid))
		for k, v := range mid {
			if v != 0 {
				node[k] = message.Node(v)
			}
		}
		if mid := ctx.GetMessage(ctx.SendChain(node...)).MessageId.ID(); mid == 0 {
			ctx.SendChain(message.Text("消息被封锁"))
		} else {
			t.Set(ctx.MessageString(), []int64{mid})
		}
		return true
	}
	if ctx.GetMessage(ctx.SendChain(message.Node(mid[0]))).MessageId.ID() == 0 {
		ctx.SendChain(message.Text("消息被封锁"))
	}
	return true
}
