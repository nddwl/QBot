package model

import "encoding/json"

type Artwork struct {
	Model
	Pid  string          `json:"-"`
	Body json.RawMessage `json:"body" gorm:"-"`
}

type ArtworkUrl struct {
	Model
	ArtworkId uint `json:"-"`
	Url       struct {
		ThumbMini string `json:"thumb_mini"`
		Small     string `json:"small"`
		Regular   string `json:"regular"`
		Original  string `json:"original"`
	} `json:"urls" gorm:"embedded"`
	CqCodeSmall    string `json:"-"`
	CqCodeOriginal string `json:"-"`
}

type ArtworkTag struct {
	Model
	Tag            string `json:"tag" gorm:"unique"`
	TagTranslation string `json:"tag_translation"`
}

type ArtworkTagAssociation struct {
	Model
	Pagination   `json:"-" gorm:"-"`
	ArtworkId    uint
	ArtworkTagId uint
}
