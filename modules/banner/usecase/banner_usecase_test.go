package usecase

import (
	"errors"
	"github.com/azbyluthfan/hello-banner/modules/banner/model"
	"github.com/azbyluthfan/hello-banner/modules/banner/query/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBannerUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	bannerQueryMock := mocks.NewMockBannerQuery(ctrl)

	t.Run("should return a valid banner", func(t *testing.T) {
		clientDate := time.Date(2019, 04, 19, 10, 0, 0, 0, time.UTC)

		expectedBanner := model.Banner{
			ID:           1,
			PromotionID:  1,
		}
		bannerQueryMock.
			EXPECT().
			Fetch(clientDate,"UTC", "192.168.1.1").
			Return(&expectedBanner, nil)

		bannerImpl := &BannerUseCaseImpl{BannerQuery: bannerQueryMock}
		banner, err := bannerImpl.GetBanner(clientDate,"UTC", "192.168.1.1")
		assert.NoError(t, err)
		assert.Equal(t, expectedBanner.ID, banner.ID)
		assert.Equal(t, expectedBanner.PromotionID, banner.PromotionID)
	})

	t.Run("should return no banner error", func(t *testing.T) {
		clientDate := time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC)

		bannerQueryMock.
			EXPECT().
			Fetch(clientDate, "UTC", "192.168.1.1").
			Return(nil, errors.New("No Banner available"))
		bannerImpl := &BannerUseCaseImpl{BannerQuery: bannerQueryMock}
		_, err := bannerImpl.GetBanner(clientDate,"UTC", "192.168.1.1")

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "No Banner available")
	})
}