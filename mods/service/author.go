package service

type AuthorService struct {
	*PixivGroup
}

func NewAuthorService(p *PixivGroup) *AuthorService {
	return &AuthorService{p}
}
