package model

import "github.com/truanter/top-go/pkg/errors"

type BaseResponse struct {
	ErrorResponse errors.ErrorResponse `json:"error_response" xml:"error_response"`
	// ResultList Simplify json array type result
	ResultList []map[string]interface{} `json:"result_list,omitempty" xml:"result_list"`
	// Data Simplify json map type result
	Data map[string]interface{} `json:"data,omitempty" xml:"data"`
}
