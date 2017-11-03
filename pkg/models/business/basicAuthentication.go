package businessModels

import (
	"encoding/json"
	"errors"

	mongoModels "github.com/Ivandolchevic/goapis/pkg/models/mongo"
	cryptoUtil "github.com/Ivandolchevic/goapis/pkg/utils/cryptoUtil"
)

// BasicAuthentication the type of a basic authentication object
type BasicAuthentication struct {
	ID          mongoModels.MongoID `json:"id" bson:"_id,omitempty"`
	Login       string              `json:"login"`
	Password    string              `json:"password"`
	Application string              `json:"application"`
	Origin      string              `json:"origin"`
}

// BasicAuthentications the type of a colllection of Basic objects
type BasicAuthentications []BasicAuthentication

// Mapper tries to map a json object to the current object type
func (b *BasicAuthentication) Mapper(byt []byte) error {
	var ojson map[string]interface{}

	// instanciate a structured object
	if err := json.Unmarshal(byt, &ojson); err != nil {
		return err
	}

	// Checking mandatory fields
	if ojson["login"] == nil || ojson["password"] == nil || ojson["application"] == nil || ojson["origin"] == nil {
		return errors.New("One or more mandatory fields could not be found")
	}

	// Map the values
	b.Login = ojson["login"].(string)
	b.Password = ojson["password"].(string)
	b.Application = ojson["application"].(string)
	b.Origin = ojson["origin"].(string)

	return nil
}

// Secure securizes the object
func (b *BasicAuthentication) Secure() error {
	cryptedpassword, err := cryptoUtil.Encrypt(b.Password)

	if err != nil {
		return err
	}

	b.Password = cryptedpassword

	return nil
}
