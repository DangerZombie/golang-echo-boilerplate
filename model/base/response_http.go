package base

import (
	"reflect"
)

// swagger:model SuccessResponse
type responseHttp struct {
	// Meta is the API response information
	// in: MetaResponse
	Meta metaResponse `json:"meta"`
	// Data is our data
	// in: DataResponse
	Data data `json:"data"`
	// Errors is the response message
	// in: string
	Errors map[string]string `json:"errors,omitempty"`
}

// swagger:model MetaResponse
type metaResponse struct {
	// Code is the response code
	// in: int
	Code int `json:"code"`
	// Message is the response message
	// in: string
	Message string `json:"message"`
	// Pagination of the paginate respons
	// in: PaginationResponse
	Pagination *Pagination `json:"pagination,omitempty"`
}

// swagger:model DataResponse
type data struct {
	Records interface{} `json:"records,omitempty"`
	Record  interface{} `json:"record,omitempty"`
}

func SetHttpResponse(code int, message string, result interface{}, paging *Pagination, errMsg map[string]string) interface{} {
	dt := data{}

	isSlice := false
	reflectResult := reflect.ValueOf(result)
	if reflectResult.Kind() == reflect.Ptr {
		isSlice = reflectResult.Elem().Kind() == reflect.Slice
	} else {
		isSlice = reflectResult.Kind() == reflect.Slice
	}

	if isSlice {
		dt.Records = result
		dt.Record = nil
	} else {
		dt.Records = nil
		dt.Record = result
	}

	return responseHttp{
		Meta: metaResponse{
			Code:       code,
			Message:    message,
			Pagination: paging,
		},
		Data:   dt,
		Errors: errMsg,
	}
}

func GetHttpResponse(resp interface{}) *responseHttp {
	result, ok := resp.(responseHttp)

	if ok {
		return &result
	}

	return nil
}
