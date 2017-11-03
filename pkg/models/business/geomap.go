package businessModels

import (
	"encoding/json"
	"fmt"

	mongoModels "github.com/Ivandolchevic/goapis/pkg/models/mongo"
)

// Geomap the type of a geographic map object
type Geomap struct {
	ID    mongoModels.MongoID `json:"id" bson:"_id,omitempty"`
	Value string              `json:"value"`
}

// Geomaps is the type of a colllection of Geomap objects
type Geomaps []Geomap

func (g *Geomap) Mapper(byt []byte) {
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat["value"].(string))
}
