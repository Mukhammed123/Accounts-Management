package service

import (
	"apulse.ai/tzuchi-upmp/server/service/hdss"
	"apulse.ai/tzuchi-upmp/server/store"
)

type Service struct {
	HDSS *hdss.Service
}

func NewService(store *store.Store) *Service {
	return &Service{
		HDSS: hdss.NewService(store),
	}
}
