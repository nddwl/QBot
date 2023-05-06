package service

type AuthorAllArtworkService struct {
	*PixivGroup
}

func NewAuthorAllArtworkService(p *PixivGroup) *AuthorAllArtworkService {
	return &AuthorAllArtworkService{p}
}
