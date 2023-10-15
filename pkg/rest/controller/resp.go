package controller

import "interview_test/pkg/model/records"

type Response struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Data    *records.Record `json:"data,omitempty"`
}
