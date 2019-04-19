package query_test

import (
	"github.com/azbyluthfan/hello-banner/modules/banner/model"
	"github.com/azbyluthfan/hello-banner/modules/banner/query/mocks"
	"github.com/golang/mock/gomock"
	"time"
	"testing"
	"errors"
)

func TestBannerQueryStub(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	bannerQuery := mocks.NewMockBannerQuery(mockCtrl)
	testBanner := &model.Banner{
		ID: 3,
		PromotionID: 3,
		DisplayStart: time.Date(2019, 04, 19,0,0,0, 0, time.UTC),
		DisplayEnd: time.Date(2019, 04, 20,0,0,0, 0, time.UTC),
	}
	dummyError := errors.New("dummy error")
	//dummyDate := time.Date(2019, 04, 19,10,0,0, 0, time.UTC)

	//bannerQuery.EXPECT().Fetch(dummyDate, "", "").Return(testBanner, dummyError)
	//bannerQuery.EXPECT().Displayed(testBanner, dummyDate, "", "").Return(true)
	bannerQuery.EXPECT().Fetch(gomock.Any(), gomock.Any(), gomock.Any()).Return(testBanner, dummyError)

}
