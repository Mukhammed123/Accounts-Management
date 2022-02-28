package scheduler

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/labstack/echo/v4"

	"apulse.ai/tzuchi-upmp/server/service"
)

type Scheduler struct {
	scheduler *gocron.Scheduler
	service   *service.Service
	logger    echo.Logger
}

func NewScheduler(service *service.Service, logger echo.Logger) *Scheduler {
	return &Scheduler{
		scheduler: gocron.NewScheduler(time.Local),
		service:   service,
		logger:    logger,
	}
}

func (s *Scheduler) SetupJobs() {
	s.scheduler.Every(time.Hour).StartAt(time.Now().Truncate(time.Hour)).Do(func() {
		if err := s.service.HDSS.SetAccessToken(); err != nil {
			s.logger.Error(err)
		}
		if err := s.service.HDSS.SyncDataOfUser(); err != nil {
			s.logger.Error(err)
		}
	})
}

func (s *Scheduler) Start() {
	s.scheduler.StartAsync()
}
