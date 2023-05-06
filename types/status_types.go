package types

type ServerStatus struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ServerStatusArray []ServerStatus
