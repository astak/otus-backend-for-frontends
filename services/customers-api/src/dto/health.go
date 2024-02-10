package dto

const (
	statusOk = "OK"
)

type healthResponse struct {
	Status string `json:"status"`
}

func NewHealthOK() *healthResponse {
	return &healthResponse{Status: statusOk}
}
