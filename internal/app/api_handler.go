package app

import (
	"github.com/ReolinkCameraAPI/reolinkapi/internal/pkg/api"
)

type ApiHandler struct {
	*api.DeviceMixin
	*api.DisplayMixin
	*api.ImageMixin
	*api.AuthMixin
	*api.RtspMixin
	*api.NetworkMixin
	*api.PtzMixin
	*api.RecordingMixin
	*api.SystemMixin
	*api.UserMixin
	*api.ZoomFocusMixin
}

func NewApiHandler() (*ApiHandler, error) {
	return &ApiHandler{
		&api.DeviceMixin{},
		&api.DisplayMixin{},
		&api.ImageMixin{},
		&api.AuthMixin{},
		&api.RtspMixin{},
		&api.NetworkMixin{},
		&api.PtzMixin{},
		&api.RecordingMixin{},
		&api.SystemMixin{},
		&api.UserMixin{},
		&api.ZoomFocusMixin{},
	}, nil
}
