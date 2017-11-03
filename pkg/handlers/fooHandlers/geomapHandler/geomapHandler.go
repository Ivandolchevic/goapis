package geomapHandler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"

	database "github.com/Ivandolchevic/goapis/pkg/data"
	. "github.com/Ivandolchevic/goapis/pkg/models/business"
	. "github.com/Ivandolchevic/goapis/pkg/models/util"
	responseUtil "github.com/Ivandolchevic/goapis/pkg/utils/responseUtil"
)

// GetAll return all the elements of a the geomap collection
func GetAll(w http.ResponseWriter, r *http.Request) *APIError {
	// read the filter
	filter := r.Header.Get("filter")

	// get the datas
	basicAuthentications, err := database.Find("geomaps", filter)

	// send back the response
	responseUtil.Answer(w, basicAuthentications, err)

	return nil
}

// Get return an element of a collection using its identifier
func Get(w http.ResponseWriter, r *http.Request) *APIError {
	return nil
}

// Put create a new basic authentication element into the database
func Put(w http.ResponseWriter, r *http.Request) *APIError {
	/*
		var geomap Geomap

		// Read the inputed data
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

		// Handle error on data reading
		if err != nil {
			responseUtil.Answer(w, body, err)
		}

		// Handler error on body closing
		if err := r.Body.Close(); err != nil {
			responseUtil.Answer(w, body, err)
		}

		// copy the body into an object
		err = json.Unmarshal(body, &geomap)
	*/
	/*
		var geomap Geomap
		reflected, err := putReflect(r)

		geomap = reflected

		// Handle the errors on data reflection
		if err != nil {
			responseUtil.Answer(w, geomap, err)
		} else {
			// Insert the data into the database
			_, err = database.Upsert("geomaps", geomap)

			// Handle errors on data insertion
			if err != nil {
				responseUtil.Answer(w, geomap, err)
			}

			responseUtil.Answer(w, geomap, nil)
		}
	*/
	var geomap Geomap

	oJson, _ := extractJSON(r)

	geomap.Mapper(oJson)

	// copy the body into an object
	//_ = json.Unmarshal(oJson, &geomap)

	fmt.Printf("%v", geomap)

	return nil
}

/*
var typeRegistry = make(map[Geomap]reflect.Type)

func init() {
    typeRegistry["MyString"] = reflect.TypeOf(MyString{})
*/
func extractJSON(r *http.Request) ([]byte, error) {
	// Read the inputed data
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	// Handler error on body closing
	r.Body.Close()

	return body, err
}

func fieldsOfInterface(i interface{}) {
	val := reflect.ValueOf(i).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", typeField.Name, valueField.Interface(), tag.Get("tag_name"))
	}
}
