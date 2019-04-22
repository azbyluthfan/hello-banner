package query

import (
	"github.com/azbyluthfan/hello-banner/modules/banner/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBannerQueryStub(t *testing.T) {

	t.Run("should return banner with earlier display period end time", func(t *testing.T) {
		expectedBanner := model.Banner{
			ID:           2,
			PromotionID:  2,
		}

		bannerQuery, _ := NewBannerStub()
		banner, err := bannerQuery.Fetch(time.Date(2019, 04, 19, 10, 0, 0, 0, time.UTC),
			"UTC", "192.168.1.1")

		assert.NoError(t, err)
		assert.Equal(t, banner.ID, expectedBanner.ID)
	})

	t.Run("should return a", func(t *testing.T) {
		expectedBanner := model.Banner{
			ID:           1,
			PromotionID:  1,
		}

		bannerQuery, _ := NewBannerStub()
		banner, err := bannerQuery.Fetch(time.Date(2019, 04, 29, 10, 0, 0, 0, time.UTC),
			"UTC", "192.168.1.1")

		assert.NoError(t, err)
		assert.Equal(t, banner.ID, expectedBanner.ID)
	})

	t.Run("should return a banner when testing with local IP", func(t *testing.T) {
		expectedBanner := model.Banner{
			ID:           2,
			PromotionID:  2,
		}

		bannerQuery, _ := NewBannerStub()
		banner, err := bannerQuery.Fetch(time.Date(2019, 01, 01, 10, 0, 0, 0, time.UTC),
			"UTC", "10.0.0.1")

		assert.NoError(t, err)
		assert.Equal(t, banner.ID, expectedBanner.ID)
	})
}
