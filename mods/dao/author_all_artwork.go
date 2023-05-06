package dao

type AuthorAllArtworkDao struct {
	*PixivGroup
}

func NewAuthorAllArtworkDao(p *PixivGroup) *AuthorAllArtworkDao {
	return &AuthorAllArtworkDao{p}
}
