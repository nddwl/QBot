package server

import (
	"QBot/mods/model"
	"QBot/utils"
	"QBot/utils/ecode"
	"QBot/utils/remember"
	"fmt"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"regexp"
	"strconv"
	"strings"
)

type ArtworkServer struct {
	*PixivGroup
	pid remember.Find
	tag remember.Find
}

func NewArtworkServer(p *PixivGroup) (a *ArtworkServer) {
	a = &ArtworkServer{PixivGroup: p}
	a.initGroup()
	return
}

// initGroup 集中注册
func (t *ArtworkServer) initGroup() {
	t.Add("pid", t.engine.OnMessage().Handle(t.Pid))
	t.pid = t.remember.Add("pid")
	t.Add("tag", t.engine.OnMessage().Handle(t.Tag))
	t.tag = t.remember.Add("tag")
}

func (t *ArtworkServer) Tag(ctx *zero.Ctx) {
	if kv, ok := utils.Find(ctx.MessageString(), "tag", 3, 1); ok {
		if t.tag.Forward(ctx) {
			return
		}
		var tag []string
		if strings.Contains(kv[1], ",") {
			tag = strings.Split(kv[1], ",")
			for _, v := range tag {
				if v == "" {
					ctx.SendChain(message.Text("标签用英文逗号(,)为间隔,逗号间不能为空"))
					return
				}
			}
		} else {
			if kv[1] == "" {
				ctx.SendChain(message.Text("标签不能为空"))
				return
			}
			tag = append(tag, kv[1])
		}
		messages := make([]int64, 2)
		var all bool
		if kv[2] != "1" {
			all = true
		}
		var page model.Pagination
		if regexp.MustCompile(`\d+`).MatchString(kv[2]) {
			page.Current, _ = strconv.Atoi(kv[3])
		}
		artworkTag, err := t.service.Pixiv.Artwork.FindArtworkTagByName(tag...)
		if err != nil {
			utils.SendErr(ctx, err)
			return
		}
		messages[0], err = t.SendTag(ctx, artworkTag)
		if err != nil {
			utils.SendErr(ctx, err)
		}
		tagName := make([]string, len(artworkTag))
		for k, v := range artworkTag {
			tagName[k] = v.Tag
		}
		artworkUrl, err := t.service.Pixiv.Artwork.FindArtworkUrlByArtworkTag(all, page, tagName...)
		if err != nil {
			utils.SendErr(ctx, err)
			return
		}
		messages[1], err = t.SendImageSmall(ctx, artworkUrl)
		if err != nil {
			utils.SendErr(ctx, err)
		}
		t.tag.Set(ctx.MessageString(), messages)
	}
}

func (t *ArtworkServer) Pid(ctx *zero.Ctx) {
	if kv, ok := utils.Find(ctx.MessageString(), "pid", 2, 1); ok {
		if t.pid.Forward(ctx) {
			return
		}
		if !regexp.MustCompile(`^[1-9]\d{5,10}$`).MatchString(kv[1]) {
			return
		}
		url, tag, err := t.FindArtwork(kv[1])
		if err != nil {
			utils.SendErr(ctx, err)
			return
		}
		messages := make([]int64, 2)
		messages[0], err = t.SendTag(ctx, tag)
		if err != nil {
			utils.SendErr(ctx, err)
		}
		switch {
		case kv[2] == "1":
			messages[1], err = t.SendImageOriginal(ctx, url)
			if err != nil {
				utils.SendErr(ctx, err)
				return
			}
		default:
			messages[1], err = t.SendImageSmall(ctx, url)
			if err != nil {
				utils.SendErr(ctx, err)
				return
			}
		}
		t.pid.Set(ctx.MessageString(), messages)
	}
}

func (t *ArtworkServer) SendImageSmall(ctx *zero.Ctx, artworkUrl []model.ArtworkUrl) (messagesId int64, err error) {
	images := make([]message.MessageSegment, len(artworkUrl))
	var add []int
	for k, v := range artworkUrl {
		if v.CqCodeSmall != "" {
			images[k] = message.Image(v.CqCodeSmall)
		} else {
			b, err := t.Request.GetImage(v.Url.Small)
			add = append(add, k)
			if err != nil {
				images[k] = message.Text(fmt.Sprintf("图片获取失败:%s", err))
				continue
			}
			images[k] = message.ImageBytes(b)
		}
	}
	messages := ctx.GetMessage(ctx.SendChain(images...))
	messagesId = messages.MessageId.ID()
	if messages.MessageId.ID() == 0 {
		err = ecode.SendMessageErr.ReSet("消息被封锁")
		return
	}
	if len(add) > 0 {
		Url := make([]model.ArtworkUrl, len(add))
		for k, v := range add {
			if s, ok := utils.CqImage(messages.Elements[v].String()); ok {
				Url[k].CqCodeSmall = s
			}
			Url[k].CqCodeOriginal = artworkUrl[v].CqCodeOriginal
			Url[k].ID = artworkUrl[v].ID
		}
		_, err = t.service.Pixiv.Artwork.UpdateArtworkUrlCqCode(Url...)
	}
	return
}

func (t *ArtworkServer) SendImageOriginal(ctx *zero.Ctx, artworkUrl []model.ArtworkUrl) (messagesId int64, err error) {
	images := make([]message.MessageSegment, len(artworkUrl))
	var add []int
	for k, v := range artworkUrl {
		if v.CqCodeOriginal != "" {
			images[k] = message.Image(v.CqCodeOriginal)
		} else {
			b, err := t.Request.GetImage(v.Url.Original)
			add = append(add, k)
			if err != nil {
				images[k] = message.Text(fmt.Sprintf("图片获取失败:%s", err))
				continue
			}
			images[k] = message.ImageBytes(b)
		}
	}
	messages := ctx.GetMessage(ctx.SendChain(images...))
	messagesId = messages.MessageId.ID()
	if messages.MessageId.ID() == 0 {
		err = ecode.SendMessageErr.ReSet("消息被封锁")
		return
	}
	if len(add) > 0 {
		Url := make([]model.ArtworkUrl, len(add))
		for k, v := range add {
			if s, ok := utils.CqImage(messages.Elements[v].String()); ok {
				Url[k].CqCodeOriginal = s
			}
			Url[k].CqCodeSmall = artworkUrl[v].CqCodeSmall
			Url[k].ID = artworkUrl[v].ID
		}
		_, err = t.service.Pixiv.Artwork.UpdateArtworkUrlCqCode(Url...)
	}
	return
}

func (t *ArtworkServer) SendTag(ctx *zero.Ctx, artworkTag []model.ArtworkTag) (messagesId int64, err error) {
	tag := "Tags:"
	for _, v := range artworkTag {
		if v.TagTranslation != "" {
			tag += "#" + v.TagTranslation + ","
		} else {
			tag += "#" + v.Tag + ","
		}
	}
	messages := ctx.GetMessage(ctx.SendChain(message.Text(tag[:len(tag)-1])))
	messagesId = messages.MessageId.ID()
	if messages.MessageId.ID() == 0 {
		err = ecode.SendMessageErr.ReSet("消息被封锁")
		return
	}
	return
}

func (t *ArtworkServer) FindArtwork(pid string) (m1 []model.ArtworkUrl, m2 []model.ArtworkTag, err error) {
	artworks, err := t.service.Pixiv.Artwork.FindArtwork(pid)
	artwork := &model.Artwork{}
	if err != nil {
		if !t.service.IsErrRecordNotFound(err) {
			return
		}
		artwork, err = t.CreateArtwork(pid)
		if err != nil {
			return
		}
		m1, err = t.CreateArtworkUrl(pid, artwork.ID)
		if err != nil {
			return
		}
		m2, err = t.CreateArtworkTag(pid, artwork.ID)
		return
	}
	artwork = &artworks[0]
	m1, err = t.service.Pixiv.Artwork.FindArtworkUrl(artwork.ID)
	if err != nil {
		if !t.service.IsErrRecordNotFound(err) {
			return
		}
		m1, err = t.CreateArtworkUrl(pid, artwork.ID)
		if err != nil {
			return
		}
	}
	m2, err = t.service.Pixiv.Artwork.FindArtworkTagByArtworkId(artwork.ID)
	if err != nil {
		if !t.service.IsErrRecordNotFound(err) {
			return
		}
		m2, err = t.CreateArtworkTag(pid, artwork.ID)
		if err != nil {
			return
		}
	}
	return
}

func (t *ArtworkServer) CreateArtwork(pid string) (m *model.Artwork, err error) {
	m, err = t.service.Pixiv.Artwork.CreateArtwork(pid)
	return
}

func (t *ArtworkServer) CreateArtworkUrl(pid string, artworkId uint) (m []model.ArtworkUrl, err error) {
	artworkUrl, err := t.Request.GetArtworkUrl(pid)
	if err != nil {
		return
	}
	for i := 0; i < len(artworkUrl); i++ {
		artworkUrl[i].ArtworkId = artworkId
	}
	m, err = t.service.Pixiv.Artwork.CreateArtworkUrl(artworkUrl...)
	return
}

func (t *ArtworkServer) CreateArtworkTag(pid string, artworkId uint) (m []model.ArtworkTag, err error) {
	artworkTag, err := t.Request.GetArtworkTag(pid)
	if err != nil {
		return
	}
	m, err = t.service.Pixiv.Artwork.FindOrCreateArtworkTag(artworkTag...)
	if err != nil {
		return
	}
	artworkTagId := make([]uint, len(m))
	for i := 0; i < len(m); i++ {
		artworkTagId[i] = m[i].ID
	}
	_, err = t.service.Pixiv.Artwork.CreateArtworkTagAssociation(artworkId, artworkTagId...)
	if err != nil {
		return
	}
	return
}
