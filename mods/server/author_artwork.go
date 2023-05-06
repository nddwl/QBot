package server

type AuthorArtworkServer struct {
	*PixivGroup
}

func NewAuthorArtworkServer(p *PixivGroup) *AuthorArtworkServer {
	return &AuthorArtworkServer{p}
}
