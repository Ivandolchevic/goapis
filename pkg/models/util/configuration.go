package utilModels

type KeyValue struct {
	Key   string
	Value string
}

// Configuration is the configuration of the service
type Configuration struct {
	HostName     string
	DatabaseName string
	Port         string
	LogFolder    string
	OptionsHeaders      []KeyValue
}

// GetServerURL return the host server complete url
func (c *Configuration) GetServerURL() string {
	return c.HostName + ":" + c.Port
}
