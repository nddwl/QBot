package dao

type AuthorPickupDao struct {
	*PixivGroup
}

func NewAuthorPickup(p *PixivGroup) *AuthorPickupDao {
	return &AuthorPickupDao{p}
}
