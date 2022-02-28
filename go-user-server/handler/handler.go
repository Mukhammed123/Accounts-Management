package handler

import (
	"apulse.ai/tzuchi-upmp/server/handler/user"
	"apulse.ai/tzuchi-upmp/server/service"
	"apulse.ai/tzuchi-upmp/server/store"
)

type Handler struct {
	user *user.Handler
}

func NewHandler(store *store.Store, service *service.Service) *Handler {
	return &Handler{
		user: user.NewHandler(store, service),
	}
}
