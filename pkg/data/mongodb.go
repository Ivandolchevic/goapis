package data

import (
	"encoding/json"
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	. "github.com/Ivandolchevic/goapis/pkg/models/util"
)

// the current database
var hostName string
var port string
var name string
var login string
var password string

// InitDatabase initialise the database configuration
func InitDatabase(configuration Configuration) {
	hostName = configuration.Database.HostName
	port = configuration.Database.Port
	name = configuration.Database.Name
	login = configuration.Database.Login
	password = configuration.Database.Password

}

// Connect open a new session to the server and connect to the database
func Connect() *mgo.Database {
	// Open a session
	session, err := mgo.Dial(hostName + ":" + port)

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	if err != nil {
		panic(err)
	}

	fmt.Println("login: " + login + ", password: " + password)

	database := session.DB(name)

	// login
	if err := database.Login(login, password); err != nil {
		panic(err)
	}

	return database
}

// Upsert update or insert an element from the database
func Upsert(collectionName string, element interface{}) (interface{}, error) {
	// define the target collection
	collection := Connect().C(collectionName)

	// insert the element
	err := collection.Insert(element)

	// Check if there is an error
	if err != nil {
		log.Fatal(err)
	}

	return element, err
}

type Container []interface{}

// Find return all the elements of a collection using the given filter
func Find(collectionName string, filter string) (Container, error) {
	// define the target collection
	collection := Connect().C(collectionName)
	log.Println("filter:" + filter)

	var raw map[string]interface{}

	json.Unmarshal([]byte(filter), &raw)

	var result Container

	err := collection.Find(raw).All(&result)

	return result, err
}

// FindByID return the element that matches with the given identifier
func FindByID(collectionName string, id string) (interface{}, error) {
	var result interface{}

	// define the target collection
	collection := Connect().C(collectionName)

	err := collection.FindId(bson.ObjectIdHex(id)).One(&result)

	return result, err
}

// FindAll return all the elements of a collection
func FindAll(collectionName string) Container {
	// define the target collection
	collection := Connect().C(collectionName)

	var result Container

	collection.Find(bson.M{}).All(&result)

	return result
}
