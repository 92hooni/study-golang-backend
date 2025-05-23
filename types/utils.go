package types

// network, repository, service 등에서 사용 할 다양한 타입을 정의한다.

type ApiResponse struct {
	Result      int64       `json:"result"`
	Description string      `json:"description"`
	ErrCode     interface{} `json:"errCode"`
}

func NewApiResponse(result int64, description string, errCode interface{}) *ApiResponse {
	return &ApiResponse{
		Result:      result,
		Description: description,
		ErrCode:     errCode,
	}
}
