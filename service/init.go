package service

import "go-Admin/define"

func NewQueryRequest() *QueryRequest {

	return &QueryRequest{
		Page:    1,
		Size:    define.DefaultSize,
		KeyWord: "",
	}
}
