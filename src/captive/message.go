package captive

// Enum of request data keys
const (
	URL  string = "url"
	Code string = "code"
)

// Request payload
type Request struct {
	Op   string            `json:"op"`
	Data map[string]string `json:"data"`
}

// Response payload
type Response struct {
	Message string `json:"message"`
}
