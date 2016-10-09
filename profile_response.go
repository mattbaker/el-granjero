package main

// ProfileResponse is a response to a profile request from bungie
type ProfileResponse struct {
	*BaseResponse
	Response []Profile
}
