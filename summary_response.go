package main

// SummaryResponse is a response to an account summary query
type SummaryResponse struct {
	*BaseResponse
	Response struct {
		Data struct {
			Characters []Character
		}
	}
}
