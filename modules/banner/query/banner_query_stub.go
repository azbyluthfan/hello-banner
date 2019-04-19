package query

import (
	"errors"
	"github.com/azbyluthfan/hello-banner/modules/banner/model"
	"sort"
	"time"
)

type BannerStub struct {
	data []model.Banner
}

var (
	whitelistIP = []string{"10.0.0.1", "10.0.0.2"}
)

func NewBannerStub() (*BannerStub, error) {

	return &BannerStub{
		data: []model.Banner{
			{
				ID: 1,
				PromotionID: 1,
				DisplayStart: time.Date(2019, 04, 19, 0, 0, 0, 0, time.UTC),
				DisplayEnd: time.Date(2019, 04, 30, 0, 0, 0, 0, time.UTC),
			},
			{
				ID: 2,
				PromotionID: 2,
				DisplayStart: time.Date(2019, 04, 19, 0, 0, 0, 0, time.UTC),
				DisplayEnd: time.Date(2019, 04, 20, 0, 0, 0, 0, time.UTC),
			},
		},
	}, nil
}


func (stub *BannerStub) Fetch(time time.Time, timezone string, ipAddress string) (*model.Banner, error) {

	sort.Slice(stub.data[:], func(i, j int) bool {
		return stub.data[i].DisplayEnd.Before(stub.data[j].DisplayEnd)
	})

	for _, banner := range stub.data {
		if stub.Displayed(banner, time, timezone, ipAddress) {
			return &banner, nil
		}
	}

	return nil, errors.New("No Banner available")
}

func (stub *BannerStub) Displayed(banner model.Banner, clientTime time.Time, timezone string, ipAddress string) bool {

	clientLoc, err := time.LoadLocation(timezone)
	if err != nil {
		clientLoc, _ = time.LoadLocation("UTC")
	}

	clientTime = clientTime.In(clientLoc)

	whitelisted := false
	for _, wl := range whitelistIP {
		if wl == ipAddress {
			whitelisted = true
		}
	}

	if (banner.DisplayStart.Before(clientTime) || whitelisted) &&
		banner.DisplayEnd.After(clientTime) {
		return true
	}
	return false
}
