package model

import "time"

type (
	Banner struct {
		ID				uint64
		PromotionID		uint64
		DisplayStart	time.Time
		DisplayEnd		time.Time
	}
)