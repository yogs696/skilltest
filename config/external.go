package config

// External API/Microservices configuration key value
type external struct {
	// API/Microservices host
	Host map[string]string `json:"host"`

	// API/Microservices token
	Token map[string]string `json:"token"`
}
