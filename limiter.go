package ratelimiterlibrary

type Limter interface {
	AllowRequest(userId string) bool
	AllowNRequests(userId string, requests int)
}
