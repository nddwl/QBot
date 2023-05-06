package model

type AuthorArtworks struct {
	Model
	Uid  string
	Body struct {
		Works     map[string]Work `json:"works"`
		ExtraData struct {
			Meta struct {
				Title       string `json:"title"`
				Description string `json:"description"`
				Canonical   string `json:"canonical"`
				Ogp         struct {
					Description string `json:"description"`
					Image       string `json:"image"`
					Title       string `json:"title"`
					Type        string `json:"type"`
				} `json:"ogp"`
				Twitter struct {
					Description string `json:"description"`
					Image       string `json:"image"`
					Title       string `json:"title"`
					Card        string `json:"card"`
				} `json:"twitter"`
			} `json:"meta"`
		} `json:"extraData"`
	} `json:"body"`
}
