package service

type WorkService struct {
	*PixivGroup
}

func NewWorkService(p *PixivGroup) *WorkService {
	return &WorkService{p}
}
