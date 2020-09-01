package reactions

import "net/http"

// Service is the service struct defined for the comment package for rpc service "reaction.*"
type Service struct{}

// List returns a list of reactions for the comments requested
func (s Service) List(r *http.Request, args *ListArgs, reply *ListResponse) error {
	return list(r, args, reply)
}

// React creates reactions for comments.
func (s Service) React(r *http.Request, args *ReactArgs, reply *ReactResponse) error {
	return react(r, args, reply)
}