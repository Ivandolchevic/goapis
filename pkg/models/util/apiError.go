package utilModels

// APIError is the configuration of the service
type APIError struct {
	Error   error
	Message string
	Code    int
}
