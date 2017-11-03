package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// RefreshToken refresh token are used to renew the access token
type RefreshToken struct {
	ID            bson.ObjectId  `json:"id" bson:"_id,omitempty"`
	Value         string         `json:"value"`
	AccessTokenID AccessTokenRef `json:"accesstoken"`
	Limit         time.Time      `json:"limit"`
}
