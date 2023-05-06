package dao

type AuthorArtworkDao struct {
	*PixivGroup
}

func NewAuthorArtwork(p *PixivGroup) *AuthorArtworkDao {
	return &AuthorArtworkDao{p}
}
