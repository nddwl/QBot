package service

import (
	"QBot/mods/model"
	"gorm.io/gorm"
)

type ArtworkService struct {
	*PixivGroup
}

func NewArtworkService(p *PixivGroup) *ArtworkService {
	return &ArtworkService{p}
}

func (t *ArtworkService) CreateArtwork(pid string) (m *model.Artwork, err error) {
	m, err = t.dao.Pixiv.Artwork.CreateArtwork(pid)
	return
}

func (t *ArtworkService) DeleteArtwork(artworkId uint) (m *model.Artwork, err error) {
	m, err = t.dao.Pixiv.Artwork.DeleteArtwork(artworkId)
	return
}

func (t *ArtworkService) DeleteArtworkAll(artworkId uint) (err error) {
	_, err = t.DeleteArtwork(artworkId)
	if err != nil {
		return
	}
	_, err = t.DeleteArtworkUrl(artworkId)
	if err != nil {
		return
	}
	_, err = t.DeleteArtworkTagAssociationAll(artworkId)
	return
}

func (t *ArtworkService) FindArtwork(pid ...string) (m []model.Artwork, err error) {
	m, err = t.dao.Pixiv.Artwork.FindArtwork(pid...)
	if len(m) == 0 {
		err = gorm.ErrRecordNotFound
	}
	return
}

func (t *ArtworkService) CreateArtworkUrl(artworkUrl ...model.ArtworkUrl) (m []model.ArtworkUrl, err error) {
	m, err = t.dao.Pixiv.Artwork.CreateArtworkUrl(artworkUrl...)
	return
}

func (t *ArtworkService) DeleteArtworkUrl(artworkUrlId ...uint) (m []model.ArtworkUrl, err error) {
	m, err = t.dao.Pixiv.Artwork.DeleteArtworkUrl(artworkUrlId...)
	return
}

func (t *ArtworkService) UpdateArtworkUrl(artworkUrl ...model.ArtworkUrl) (m []model.ArtworkUrl, err error) {
	m, err = t.dao.Pixiv.Artwork.UpdateArtworkUrl(artworkUrl...)
	return
}

func (t *ArtworkService) UpdateArtworkUrlCqCode(artworkUrl ...model.ArtworkUrl) (m []model.ArtworkUrl, err error) {
	m, err = t.dao.Pixiv.Artwork.UpdateArtworkUrlCqCode(artworkUrl...)
	return
}

func (t *ArtworkService) FindArtworkUrl(artworkId ...uint) (m []model.ArtworkUrl, err error) {
	m, err = t.dao.Pixiv.Artwork.FindArtworkUrl(artworkId...)
	if len(m) == 0 {
		err = gorm.ErrRecordNotFound
	}
	return
}

func (t *ArtworkService) CreateArtworkTag(artworkTag ...model.ArtworkTag) (m []model.ArtworkTag, err error) {
	m, err = t.dao.Pixiv.Artwork.CreateArtworkTag(artworkTag...)
	return
}

func (t *ArtworkService) FindArtworkTag(artworkTagId ...uint) (m []model.ArtworkTag, err error) {
	m, err = t.dao.Pixiv.Artwork.FindArtworkTag(artworkTagId...)
	if len(m) == 0 {
		err = gorm.ErrRecordNotFound
	}
	return
}

func (t *ArtworkService) FindArtworkTagByName(tag ...string) (m []model.ArtworkTag, err error) {
	m, err = t.dao.Pixiv.Artwork.FindArtworkTagByName(tag...)
	if len(m) == 0 {
		err = gorm.ErrRecordNotFound
	}
	return
}

func (t *ArtworkService) FindOrCreateArtworkTag(artworkTag ...model.ArtworkTag) (m []model.ArtworkTag, err error) {
	if len(artworkTag) == 0 {
		artworkTag = append(artworkTag, model.ArtworkTag{Tag: "can't find tag"})
	}
	tag := make([]string, len(artworkTag))
	for k, v := range artworkTag {
		tag[k] = v.Tag
	}
	m1, err := t.Pixiv.Artwork.FindArtworkTagByName(tag...)
	if err != nil {
		if !t.IsErrRecordNotFound(err) {
			return
		}
	}
	var m2 []model.ArtworkTag
	if len(m1) != len(artworkTag) {
		m2 = make([]model.ArtworkTag, len(artworkTag)-len(m1))
		var index int
		var artworkTagMap = make(map[string]struct{}, len(m1))
		for _, v := range m1 {
			artworkTagMap[v.Tag] = struct{}{}
		}
		for i := 0; i < len(artworkTag); i++ {
			if _, ok := artworkTagMap[artworkTag[i].Tag]; !ok {
				m2[index] = artworkTag[i]
				index++
			}
		}
		m2, err = t.Pixiv.Artwork.CreateArtworkTag(m2...)
		if err != nil {
			return
		}
	}
	m = append(m1, m2...)
	return
}

func (t *ArtworkService) FindArtworkTagByArtworkId(artworkId ...uint) (m []model.ArtworkTag, err error) {
	m, err = t.dao.Pixiv.Artwork.FindArtworkTagByArtworkId(artworkId...)
	if len(m) == 0 {
		err = gorm.ErrRecordNotFound
	}
	return
}

func (t *ArtworkService) CreateArtworkTagAssociation(artworkId uint, artworkTagId ...uint) (m []model.ArtworkTagAssociation, err error) {
	m, err = t.dao.Pixiv.Artwork.CreateArtworkTagAssociation(artworkId, artworkTagId...)
	return
}

func (t *ArtworkService) FindArtworkTagAssociation(artworkId uint) (m []model.ArtworkTagAssociation, err error) {
	m, err = t.dao.Pixiv.Artwork.FindArtworkTagAssociation(artworkId)
	if len(m) == 0 {
		err = gorm.ErrRecordNotFound
	}
	return
}

func (t *ArtworkService) DeleteArtworkTagAssociation(artworkId uint, artworkTagId ...uint) (m []model.ArtworkTagAssociation, err error) {
	m, err = t.dao.Pixiv.Artwork.DeleteArtworkTagAssociation(artworkId, artworkTagId...)
	return
}

func (t *ArtworkService) DeleteArtworkTagAssociationAll(artworkId uint) (m []model.ArtworkTagAssociation, err error) {
	m, err = t.dao.Pixiv.Artwork.DeleteArtworkTagAssociationAll(artworkId)
	return
}

func (t *ArtworkService) FindArtworkUrlByArtworkTag(all bool, page model.Pagination, tag ...string) (m []model.ArtworkUrl, err error) {
	m, err = t.dao.Pixiv.Artwork.FindArtworkUrlByArtworkTag(all, page, tag...)
	if len(m) == 0 {
		err = gorm.ErrRecordNotFound
	}
	return
}
