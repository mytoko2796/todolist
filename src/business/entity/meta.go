package entity

type Meta struct {
	Path       string           `json:"path"`
	StatusCode int              `json:"status_code"`
	Status     string           `json:"status"`
	Message    string           `json:"message"`
	Error      error `json:"error,omitempty" swaggertype:"primitive,object"`
	Timestamp  string           `json:"timestamp"`
}