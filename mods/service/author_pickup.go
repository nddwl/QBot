package service

type AuthorPickupService struct {
	*PixivGroup
}

func NewAuthorPickupService(p *PixivGroup) *AuthorPickupService {
	return &AuthorPickupService{p}
}
