package common

import "time"

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type TBaseResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type HttpArgs struct {
	Method         string
	Url            string
	Data           []byte
	Headers        map[string]string
	ResponseStruct interface{}
	Proxy          string
	TimeoutSecond  time.Duration
}
