package transport

// Request interface
type Transport interface {
	// Create request
	Call(method string, out interface{}, params ...interface{}) error
	// Close transport connection
	Close() error
}

// Create a new transport object
func NewTransport(url string, headers map[string]string) (Transport, error) {
	return newHTTP(url, headers), nil
}
