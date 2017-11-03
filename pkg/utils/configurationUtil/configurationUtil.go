// Package configuration provides an access to all configuration parameters
package configurationUtil

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	. "github.com/Ivandolchevic/goapis/pkg/models/util"
)

// Get return the configuration object from the configuration file
func Get() Configuration {
	path := filepath.Base(".")
	fmt.Print(path)
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}
