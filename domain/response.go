package domain

type BaseResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}
