package services

import "time"

type ErrorStruct struct {
	TimeStamp int64       `json:"time_stamp"`
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

func (e ErrorStruct) AddTimeStamp() ErrorStruct {
	e.TimeStamp = time.Now().UnixMilli()
	return e
}
