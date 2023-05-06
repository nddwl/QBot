package dao

type AuthorDao struct {
	*PixivGroup
}

func NewAuthorDao(p *PixivGroup) *AuthorDao {
	return &AuthorDao{p}
}
