package routes

import "github.com/malaika-muneer/File-Analyser/service"

type Router struct {
	userService service.UserService
}

// NewRouter initializes a Router instance with the provided UserService
func NewRouter(userService service.UserService) *Router {
	return &Router{
		userService: userService,
	}
}
