package main

import (
	"fmt"
	bannerQuery   "github.com/azbyluthfan/hello-banner/modules/banner/query"
	bannerUseCase "github.com/azbyluthfan/hello-banner/modules/banner/usecase"
	"time"
)

func main() {

	bQuery, _ := bannerQuery.NewBannerStub()
	bannerImpl := bannerUseCase.NewBannerUseCase(bQuery)

	// two banners available, showing earlier deadline
	banner, err := bannerImpl.GetBanner(time.Date(2019, 04, 19, 10, 0, 0, 0, time.UTC),
		"UTC", "192.168.1.1")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Showing Banner ID:", banner.ID, "Promotion ID:", banner.PromotionID)
	}

	// one banner available
	banner, err = bannerImpl.GetBanner(time.Date(2019, 04, 20, 10, 0, 0, 0, time.UTC),
		"Asia/Tokyo", "192.168.1.1")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Showing Banner ID:", banner.ID, "Promotion ID:", banner.PromotionID)
	}

	// local IP Address, before banner display period starts
	banner, err = bannerImpl.GetBanner(time.Date(2019, 04, 01, 0, 0, 0, 0, time.UTC),
		"Asia/Tokyo", "10.0.0.1")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Showing Banner ID:", banner.ID, "Promotion ID:", banner.PromotionID)
	}
}