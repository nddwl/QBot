package service

type PixivGroup struct {
	*Service
	Artwork          *ArtworkService
	Author           *AuthorService
	AuthorAllArtwork *AuthorAllArtworkService
	AuthorArtwork    *AuthorArtworkService
	AuthorPickup     *AuthorPickupService
	Work             *WorkService
}

func NewPixivGroup(s *Service) (p *PixivGroup) {
	p = &PixivGroup{
		Service: s,
	}
	p.init()
	p.initGroup()
	return
}

func (t *PixivGroup) init() {

}

func (t *PixivGroup) initGroup() {
	t.Artwork = NewArtworkService(t)
	t.Author = NewAuthorService(t)
	t.AuthorAllArtwork = NewAuthorAllArtworkService(t)
	t.AuthorArtwork = NewAuthorArtworkService(t)
	t.AuthorPickup = NewAuthorPickupService(t)
	t.Work = NewWorkService(t)
}
