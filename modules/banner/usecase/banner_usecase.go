package usecase

import (
	"github.com/azbyluthfan/hello-banner/modules/banner/model"
	"github.com/azbyluthfan/hello-banner/modules/banner/query"
	"time"
)

type BannerUseCaseImpl struct {
	BannerQuery query.BannerQuery
}

func NewBannerUseCase(bannerQuery query.BannerQuery) BannerUseCase {
	return &BannerUseCaseImpl{
		BannerQuery: bannerQuery,
	}
}

func (impl *BannerUseCaseImpl) GetBanner(
	time time.Time,
	timezone string,
	ipAddress string) (*model.Banner, error) {

	banner, err := impl.BannerQuery.Fetch(time, timezone, ipAddress)
	if err != nil {
		return nil, err
	}

	return banner, nil
}