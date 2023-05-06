package server

type WorkServer struct {
	*PixivGroup
}

func NewWorkServer(p *PixivGroup) *WorkServer {
	return &WorkServer{p}
}
