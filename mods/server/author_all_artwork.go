package server

type AuthorAllArtworkServer struct {
	*PixivGroup
}

func NewAuthorAllArtworkServer(p *PixivGroup) *AuthorAllArtworkServer {
	return &AuthorAllArtworkServer{p}
}
