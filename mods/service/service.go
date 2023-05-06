package service

import (
	"QBot/mods/dao"
	"QBot/utils/config"
	"errors"
	"gorm.io/gorm"
)

type Service struct {
	dao   *dao.Dao
	Pixiv *PixivGroup
}

func New() (s *Service) {
	s = &Service{
		dao: dao.New(),
	}

	s.initGroup()
	s.init()
	return
}

func (t *Service) init() {
}

func (t *Service) initGroup() {
	t.Pixiv = NewPixivGroup(t)
}

func (t *Service) IsErrRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func newTest() *Service {
	config.Init()
	return New()
}
