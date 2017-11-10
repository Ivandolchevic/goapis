package utilModels

type KeyValue struct {
	Key   string
	Value string
}

type APIConfiguration struct {
	HostName string
	Port     string
}

type DatabaseConfiguration struct {
	HostName string
	Port     string
	Name     string
	Login    string
	Password string
}

// Configuration is the configuration of the service
type Configuration struct {
	API            APIConfiguration
	Database       DatabaseConfiguration
	LogFolder      string
	OptionsHeaders []KeyValue
}

// GetServerURL return the host server complete url
func (c *Configuration) GetServerURL() string {
	return c.API.HostName + ":" + c.API.Port
}
