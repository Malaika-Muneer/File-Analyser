package routes

import (
	"github.com/malaika-muneer/File-Analyser/service"
)

type Router struct {
	UserService   service.UserService
	UploadService service.UploadServiceInterface
}

func NewRouter(userService service.UserService, uploadService service.UploadServiceInterface) *Router {
	return &Router{
		UserService:   userService,
		UploadService: uploadService,
	}
}
