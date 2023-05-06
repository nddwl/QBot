package configs

type Pixiv struct {
	Cookie    string `json:"cookie"`
	UserAgent string `json:"user_agent"`
	Version   string `json:"version"`
	Proxy     string `json:"proxy"`
}
