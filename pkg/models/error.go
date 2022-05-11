package models

type ErrorResponse struct {
	Code 		int			`json:"errorCode"`
	Message 	string		`json:"errorMessage"`
}