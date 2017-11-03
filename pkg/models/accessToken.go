package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// AccessToken the type of an access token  object
type AccessToken struct {
	ID    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Value string        `json:"value"`
	Limit time.Time     `json:"limit"`
}

// AccessTokens the type of a colllection of Basic objects
type AccessTokens []AccessToken

// AccessTokenRef is an object to reference an access token from another mongo table
type AccessTokenRef struct {
	ID string `json:"$id"`
}
