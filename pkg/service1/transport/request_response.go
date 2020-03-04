package transport

/*
NEW_HANDLER_STEP4: add request/response for transport layer. They describe what parameters needed to call
a new function and how response from it will be presented.
*/
type ConfigRequest struct {
}

type ConfigResponse struct {
	Config map[string]interface{} `json:"config"`
	Err    error                  `json:"error,omitempty"`
}

type InfoRequest struct {
}

type InfoResponse struct {
	Uptime float64 `json:"uptime"`
	Err    error   `json:"error,omitempty"`
}
