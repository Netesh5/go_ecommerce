package responsehandler

import "github.com/netesh5/go_ecommerce/internal/models"

type SuccessResponse struct {
	Success        bool               `json:"success"`
	Message        string             `json:"message,omitempty"`
	Data           interface{}        `json:"data,omitempty"`
	PaginationData *models.Pagination `json:"pagination,omitempty"`
}

func SuccessWithData(data interface{}, message string) SuccessResponse {
	return SuccessResponse{
		Success: true,
		Data:    data,
	}
}

func SuccessMessage(message string) SuccessResponse {
	return SuccessResponse{
		Success: true,
		Message: message,
	}
}

func SuccessWithPaginatedData(data interface{}, pagination models.Pagination, message string) SuccessResponse {
	return SuccessResponse{
		Success:        true,
		Data:           data,
		PaginationData: &pagination,
	}
}
