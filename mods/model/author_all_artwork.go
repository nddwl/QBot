package model

type AuthorAllArtwork struct {
	Model
	Uid  string
	Body struct {
		Illusts map[string]string `json:"illusts"`
		Manga   map[string]string `json:"manga"`
		Pickup  []AuthorPickup    `json:"pickup"`
	} `json:"body"`
}
