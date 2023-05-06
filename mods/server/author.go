package server

type AuthorServer struct {
	*PixivGroup
}

func NewAuthorServer(p *PixivGroup) *AuthorServer {
	return &AuthorServer{p}
}
