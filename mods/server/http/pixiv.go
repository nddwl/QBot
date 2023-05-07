package http

import (
	"QBot/mods/model"
	"QBot/utils"
	"QBot/utils/config"
	"QBot/utils/ecode"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

const (
	PIXIV = "https://www.pixiv.net/"
)

type Pixiv struct {
	client *http.Client
}

type PixivData struct {
	Error   bool            `json:"error"`
	Message string          `json:"message"`
	Body    json.RawMessage `json:"body"`
}

func (p *PixivData) IsOk() bool {
	return p.Error == false
}

func NewPixivErr(data PixivData) ecode.Codes {
	return ecode.PixivErr.ReSet(data.Message)
}

func NewPixiv() (p *Pixiv) {
	p = &Pixiv{
		client: http.DefaultClient,
	}
	p.setProxy()
	return
}

func (p *Pixiv) setProxy() {
	u, err := url.Parse(config.Pixiv.Proxy)
	if err != nil {
		log.Fatalf("parse proxy err")
	}
	p.client.Transport = &http.Transport{
		Proxy: http.ProxyURL(u),
	}
}

func (p *Pixiv) HandleStatusCode(code int, b []byte) (err error) {
	if code == 200 {
		return nil
	}
	var pixivErr PixivData
	err = utils.Json(&pixivErr, b)
	if err == nil && !pixivErr.IsOk() {
		err = NewPixivErr(pixivErr)
		return
	}
	if code == 403 {
		err = NewPixivErr(PixivData{Message: "需要登录"})
		return
	}
	if code == 404 {
		err = NewPixivErr(PixivData{Message: "页面不存在"})
		return
	}
	return NewPixivErr(PixivData{Message: fmt.Sprintf("错误的状态码:%d", code)})
}

// GET 请求
func (p *Pixiv) GET(url string) (b []byte, err error) {
	fmt.Printf("GET:%s\n", url)
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Http->Pixiv->GET->err:%s\n", err)
	}

	r.Header.Add("Cookie", config.Pixiv.Cookie)
	r.Header.Add("User-Agent", config.Pixiv.UserAgent)
	r.Header.Add("Referer", PIXIV)

	re, err := p.client.Do(r)
	if err != nil {
		log.Fatalf("Http->Pixiv->GET->err:%s\n", err)
	}
	defer re.Body.Close()
	fmt.Println(re.Request.Method, re.Request.URL, ":", re.Status)
	b, err = io.ReadAll(re.Body)
	if err != nil {
		log.Fatalf("Http->Pixiv->GET->err:%s\n", err)
	}
	err = p.HandleStatusCode(re.StatusCode, b)
	return
}

// POST 请求
func (p *Pixiv) POST(url string) (b []byte, err error) {
	fmt.Printf("POST:%s\n", url)
	r, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatalf("Http->Pixiv->POST->err:%s\n", err)
	}

	r.Header.Add("Cookie", config.Pixiv.Cookie)
	r.Header.Add("User-Agent", config.Pixiv.UserAgent)
	r.Header.Add("Referer", PIXIV)

	re, err := p.client.Do(r)
	if err != nil {
		log.Fatalf("Http->Pixiv->POST->err:%s\n", err)
	}
	defer re.Body.Close()
	fmt.Println(re.Request.Method, re.Request.URL, ":", re.Status)
	b, err = io.ReadAll(re.Body)
	if err != nil {
		log.Fatalf("Http->Pixiv-POST->err:%s\n", err)
	}
	err = p.HandleStatusCode(re.StatusCode, b)
	return
}

// GetJson 请求
func (p *Pixiv) GetJson(url string) (body json.RawMessage, err error) {
	b, err := p.GET(url)
	if err != nil {
		return
	}
	var data PixivData
	err = utils.Json(&data, b)
	if err != nil {
		return
	}
	if !data.IsOk() {
		err = NewPixivErr(data)
		return
	}
	body = data.Body
	return
}

func (p *Pixiv) GetArtworkUrl(pid string) (m []model.ArtworkUrl, err error) {

	u := PIXIV + "ajax/illust/" + pid + "/pages?lang=zn&version=" + config.Pixiv.Version

	b, err := p.GetJson(u)
	if err != nil {
		return
	}
	err = utils.JsonRaw(&m, b)
	return
}

func (p *Pixiv) GetArtworkTag(pid string) (m []model.ArtworkTag, err error) {
	u := PIXIV + "ajax/tags/frequent/illust?ids[]=" + pid + "&lang=zn&version=" + config.Pixiv.Version

	b, err := p.GetJson(u)
	if err != nil {
		return
	}
	err = utils.JsonRaw(&m, b)
	return
}

func (p *Pixiv) GetImage(url string) (b []byte, err error) {
	b, err = p.GET(url)
	return
}

func (p *Pixiv) PostImage(url string) (b []byte, err error) {
	b, err = p.POST(url)
	return
}
