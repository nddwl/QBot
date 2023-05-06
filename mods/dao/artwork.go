package dao

import (
	"QBot/mods/model"
	"gorm.io/gorm/clause"
)

type ArtworkDao struct {
	*PixivGroup
}

func NewArtworkDao(p *PixivGroup) *ArtworkDao {
	return &ArtworkDao{p}
}

func (t *ArtworkDao) CreateArtwork(pid string) (m *model.Artwork, err error) {
	m = &model.Artwork{
		Model: model.Model{},
		Pid:   pid,
	}
	err = t.db.Model(&model.Artwork{}).Create(&m).Error
	return
}

func (t *ArtworkDao) DeleteArtwork(artworkId uint) (m *model.Artwork, err error) {
	err = t.db.Model(&model.Artwork{}).Where("id", artworkId).Delete(&m).Error
	return
}

func (t *ArtworkDao) FindArtwork(pid ...string) (m []model.Artwork, err error) {
	err = t.db.Model(&model.Artwork{}).Where("pid IN (?)", pid).First(&m).Error
	return
}

func (t *ArtworkDao) CreateArtworkUrl(artworkUrl ...model.ArtworkUrl) (m []model.ArtworkUrl, err error) {
	m = make([]model.ArtworkUrl, len(artworkUrl))
	for k, v := range artworkUrl {
		m[k] = model.ArtworkUrl{
			ArtworkId: v.ArtworkId,
			Url:       v.Url,
		}
	}
	err = t.db.Model(&model.ArtworkUrl{}).Create(&m).Error
	return
}

func (t *ArtworkDao) DeleteArtworkUrl(artworkUrlId ...uint) (m []model.ArtworkUrl, err error) {
	err = t.db.Model(&model.ArtworkUrl{}).Where("id IN (?)", artworkUrlId).Delete(&m).Error
	return
}

func (t *ArtworkDao) UpdateArtworkUrl(artworkUrl ...model.ArtworkUrl) (m []model.ArtworkUrl, err error) {
	m = make([]model.ArtworkUrl, len(artworkUrl))
	for k, v := range artworkUrl {
		m[k] = model.ArtworkUrl{
			Model:          model.Model{ID: v.ID},
			Url:            v.Url,
			CqCodeSmall:    "",
			CqCodeOriginal: "",
		}
	}
	err = t.db.Model(&model.ArtworkUrl{}).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"id", "updated_at", "thumb_mini", "small", "regular", "original", "cq_code_small", "cq_code_original"}),
	}).Create(&m).Error
	return
}

func (t *ArtworkDao) UpdateArtworkUrlCqCode(artworkUrl ...model.ArtworkUrl) (m []model.ArtworkUrl, err error) {
	m = make([]model.ArtworkUrl, len(artworkUrl))
	for k, v := range artworkUrl {
		m[k] = model.ArtworkUrl{
			Model:          model.Model{ID: v.ID},
			CqCodeSmall:    v.CqCodeSmall,
			CqCodeOriginal: v.CqCodeOriginal,
		}
	}
	err = t.db.Model(&model.ArtworkUrl{}).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"id", "updated_at", "cq_code_small", "cq_code_original"}),
	}).Create(&m).Error
	return
}

func (t *ArtworkDao) FindArtworkUrl(artworkId ...uint) (m []model.ArtworkUrl, err error) {
	err = t.db.Model(&model.ArtworkUrl{}).Where("artwork_id IN (?)", artworkId).Find(&m).Error
	return
}

func (t *ArtworkDao) CreateArtworkTag(artworkTag ...model.ArtworkTag) (m []model.ArtworkTag, err error) {
	m = make([]model.ArtworkTag, len(artworkTag))
	for k, v := range artworkTag {
		m[k] = model.ArtworkTag{
			Tag:            v.Tag,
			TagTranslation: v.TagTranslation,
		}
	}
	err = t.db.Model(&model.ArtworkTag{}).Create(&m).Error
	return
}

func (t *ArtworkDao) FindArtworkTag(artworkTagId ...uint) (m []model.ArtworkTag, err error) {
	err = t.db.Model(&model.ArtworkTag{}).Where("id IN (?)", artworkTagId).Find(&m).Error
	return
}

func (t *ArtworkDao) FindArtworkTagByName(tag ...string) (m []model.ArtworkTag, err error) {
	err = t.db.Model(&model.ArtworkTag{}).Where("tag IN (?)", tag).Find(&m).Error
	return
}

func (t *ArtworkDao) FindArtworkTagByArtworkId(artworkId ...uint) (m []model.ArtworkTag, err error) {
	err = t.db.Model(&model.ArtworkTag{}).Joins("JOIN artwork_tag_association ata ON ata.artwork_tag_id = artwork_tag.id").
		Distinct("ata.artwork_tag_id").Where("artwork_id IN (?)", artworkId).Select("artwork_tag.*").Find(&m).Error
	return
}

func (t *ArtworkDao) CreateArtworkTagAssociation(artworkId uint, artworkTagId ...uint) (m []model.ArtworkTagAssociation, err error) {
	m = make([]model.ArtworkTagAssociation, len(artworkTagId))
	for k, v := range artworkTagId {
		m[k] = model.ArtworkTagAssociation{
			ArtworkId:    artworkId,
			ArtworkTagId: v,
		}
	}
	err = t.db.Model(&model.ArtworkTagAssociation{}).Create(&m).Error
	return
}

func (t *ArtworkDao) DeleteArtworkTagAssociation(artworkId uint, artworkTagId ...uint) (m []model.ArtworkTagAssociation, err error) {
	m = make([]model.ArtworkTagAssociation, len(artworkTagId))
	for k, v := range artworkTagId {
		m[k] = model.ArtworkTagAssociation{
			ArtworkId:    artworkId,
			ArtworkTagId: v,
		}
	}
	err = t.db.Model(&model.ArtworkTagAssociation{}).Where("artworkId", artworkId).Where("artwork_tag_id IN (?)", artworkTagId).Delete(&m).Error
	return
}

func (t *ArtworkDao) DeleteArtworkTagAssociationAll(artworkId uint) (m []model.ArtworkTagAssociation, err error) {
	err = t.db.Model(&model.ArtworkTagAssociation{}).Where(artworkId).Delete(&m).Error
	return
}

func (t *ArtworkDao) FindArtworkTagAssociation(artworkId uint) (m []model.ArtworkTagAssociation, err error) {
	err = t.db.Model(&model.ArtworkTagAssociation{}).Where("artwork_id", artworkId).Find(&m).Error
	return
}

func (t *ArtworkDao) FindArtworkUrlByArtworkTag(all bool, page model.Pagination, tag ...string) (m []model.ArtworkUrl, err error) {
	query := t.db.Model(&model.ArtworkUrl{}).Joins("JOIN artwork_tag_association ata ON ata.artwork_id = artwork_url.artwork_id").
		Joins("JOIN artwork_tag ON artwork_tag.id = ata.artwork_tag_id").Where("artwork_tag.tag IN (?)", tag)
	if all {
		query.Group("artwork_url.id").Having("COUNT(1) = (?)", len(tag))
	}
	err = query.Select("artwork_url.*").Scopes(page.Sql()).Find(&m).Error
	return
}
