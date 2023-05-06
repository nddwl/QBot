package server

import (
	"QBot/mods/server/http"
	"QBot/mods/service"
	"QBot/utils/remember"
	"fmt"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"strings"
)

type PixivGroup struct {
	*Server
	service          *service.Service
	remember         *remember.Remember
	Request          *http.Pixiv
	Artwork          *ArtworkServer
	Author           *AuthorServer
	AuthorAllArtwork *AuthorAllArtworkServer
	AuthorArtwork    *AuthorArtworkServer
	AuthorPickup     *AuthorPickupServer
	Work             *WorkServer
}

func NewPixivGroup(s *Server) (p *PixivGroup) {
	p = &PixivGroup{
		Server:   s,
		service:  service.New(),
		remember: remember.New(),
		Request:  http.NewPixiv(),
	}
	p.init()
	p.initGroup()
	return
}

func (t *PixivGroup) initGroup() {
	t.Artwork = NewArtworkServer(t)
	t.Author = NewAuthorServer(t)
	t.AuthorAllArtwork = NewAuthorAllArtworkServer(t)
	t.AuthorArtwork = NewAuthorArtworkServer(t)
	t.AuthorPickup = NewAuthorPickupServer(t)
	t.Work = NewWorkServer(t)
}

func (t *PixivGroup) init() {
	t.Add("pixiv", t.engine.OnFullMatch("pixiv").Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Text(fmt.Sprintf("pixiv is ready!\n已经注册的匹配:\n%s", strings.Join(t.FindGroupAll(), "\n"))))
	}))
}

func (t *PixivGroup) Add(prefix string, matcher *zero.Matcher) (m *zero.Matcher) {
	m = t.matcher.Add(prefix, matcher)
	return
}

func (t *PixivGroup) Get(name string) (m *zero.Matcher) {
	return t.matcher.Get(name)
}

func (t *PixivGroup) FindGroupAll() (name []string) {
	return t.matcher.FindAll()
}
