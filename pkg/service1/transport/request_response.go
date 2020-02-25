package transport

type ConfigRequest struct {
}

type ConfigResponse struct {
	Config map[string]interface{} `json:"config"`
	Err    error                  `json:"error,omitempty"`
}
