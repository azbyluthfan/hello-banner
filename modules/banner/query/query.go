package query

import (
	"github.com/azbyluthfan/hello-banner/modules/banner/model"
	"time"
)

type BannerQuery interface {
	Fetch(time time.Time, timezone string, ipAddress string) (*model.Banner, error)
	Displayed(banner model.Banner, time time.Time, timezone string, ipAddress string) bool
}
