package server

import (
	"QBot/utils/bot"
	"QBot/utils/matcher"
	"fmt"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"strings"
)

type Server struct {
	engine  *zero.Engine
	matcher *matcher.Matcher
	Pixiv   *PixivGroup
}

func New() (s *Server) {
	s = &Server{
		engine:  bot.Default,
		matcher: matcher.New(),
	}
	s.initGroup()
	s.init()
	fmt.Println("加载项完成")
	return
}

func (t *Server) initGroup() {
	t.Pixiv = NewPixivGroup(t)
}

func (t *Server) init() {
	t.engine.OnFullMatch("all").Handle(func(ctx *zero.Ctx) {
		pixiv := fmt.Sprintf("pixiv is ready!\n已经注册的匹配:\n%s", strings.Join(t.Pixiv.FindGroupAll(), "\n"))
		ctx.SendChain(message.Text(strings.Join([]string{pixiv}, "\n")))
	})
}
