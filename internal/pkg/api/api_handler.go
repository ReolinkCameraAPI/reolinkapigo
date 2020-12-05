package api

type ApiHandler struct {
	Device  DeviceMixin
	Display DisplayMixin
	Image   ImageMixin
}

func NewApiHandler() *ApiHandler {
	return &ApiHandler{}
}
