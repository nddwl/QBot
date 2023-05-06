package bot

import (
	zero "github.com/wdvxdr1123/ZeroBot"
)

var Default = New()

func New() *zero.Engine {
	engine := zero.New()
	return engine
}
