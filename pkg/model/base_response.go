package model

import "github.com/truanter/top-go/pkg/errors"

type BaseResponse struct {
	ErrorResponse errors.ErrorResponse `json:"error_response" xml:"error_response"`
	// ResultList Simplify json result
	ResultList []map[string]interface{} `json:"result_list" xml:"result_list"`
}
