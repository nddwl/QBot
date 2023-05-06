package model

import (
	"time"
)

type Work struct {
	Model
	Pid             string    `json:"id"`
	Title           string    `json:"title"`
	IllustType      int       `json:"illustType"`
	URL             string    `json:"url"`
	Description     string    `json:"description"`
	Tags            []string  `json:"tags"`
	UserID          string    `json:"userId"`
	UserName        string    `json:"userName"`
	PageCount       int       `json:"pageCount"`
	Alt             string    `json:"alt"`
	CreateDate      time.Time `json:"createDate"`
	UpdateDate      time.Time `json:"updateDate"`
	AiType          int       `json:"aiType"`
	ProfileImageURL string    `json:"profileImageUrl"`
}
