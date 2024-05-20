package apiResponse

import ()

type ApiResponseDTO struct {
	Code    int32  `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}
