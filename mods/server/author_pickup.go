package server

type AuthorPickupServer struct {
	*PixivGroup
}

func NewAuthorPickupServer(p *PixivGroup) *AuthorPickupServer {
	return &AuthorPickupServer{p}
}
