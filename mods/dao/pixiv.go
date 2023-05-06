package dao

type PixivGroup struct {
	*Dao
	Artwork          *ArtworkDao
	Author           *AuthorDao
	AuthorAllArtwork *AuthorAllArtworkDao
	AuthorArtwork    *AuthorArtworkDao
	AuthorPickup     *AuthorPickupDao
	Work             *WorkDao
}

func NewPixivGroup(d *Dao) (p *PixivGroup) {
	p = &PixivGroup{Dao: d}

	p.init()
	p.initGroup()
	return
}

func (t *PixivGroup) initGroup() {
	t.Artwork = NewArtworkDao(t)
	t.Author = NewAuthorDao(t)
	t.AuthorAllArtwork = NewAuthorAllArtworkDao(t)
	t.AuthorArtwork = NewAuthorArtwork(t)
	t.AuthorPickup = NewAuthorPickup(t)
	t.Work = NewWork(t)
}

func (t *PixivGroup) init() {

}
