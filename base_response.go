package main

// BaseResponse is a general API response from bungie
type BaseResponse struct {
	ErrorCode       int
	ThrottleSeconds int
	ErrorStatus     string
	Message         string
}
