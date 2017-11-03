package mongoModels

import (
	"gopkg.in/mgo.v2/bson"
)

// MongoID is the interface of all mongo items
type MongoID struct {
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
}
