package repositories

import "goHexBoilerplate/src/modules/analytics/domain/contracts/entities"

type Create struct {
	ShortenerId string
	UserIp      *string
	ReferrerUrl *string
	UserAgent   *string
}

type ClickRepository interface {
	Create(params Create) (*entities.Click, error)
}
