package usecase

import (
	"github.com/azbyluthfan/hello-banner/modules/banner/model"
	"time"
)

type BannerUseCase interface {
	GetBanner(time time.Time, timezone string, ipAddress string) (*model.Banner, error)
}