package model

type Author struct {
	Model
	Body struct {
		UserID     string `json:"userId"`
		Name       string `json:"name"`
		Image      string `json:"image"`
		ImageBig   string `json:"imageBig"`
		IsFollowed bool   `json:"isFollowed"`
		IsMypixiv  bool   `json:"isMypixiv"`
		IsBlocking bool   `json:"isBlocking"`
	} `json:"body"`
}
