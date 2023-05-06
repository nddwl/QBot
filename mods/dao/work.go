package dao

type WorkDao struct {
	*PixivGroup
}

func NewWork(p *PixivGroup) *WorkDao {
	return &WorkDao{p}
}
