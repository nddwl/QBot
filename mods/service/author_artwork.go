package service

type AuthorArtworkService struct {
	*PixivGroup
}

func NewAuthorArtworkService(p *PixivGroup) *AuthorArtworkService {
	return &AuthorArtworkService{p}
}
