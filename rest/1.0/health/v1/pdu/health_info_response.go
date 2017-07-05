package pdu

type HealthInfoResponse struct {
	Build BuildInfoResponse `json:"build,omitempty"`
}

type BuildInfoResponse struct {
	Name       string `json:"name"`
	Release    string `json:"release"`
	Build      string `json:"build"`
	Repository string `json:"repository"`
}
