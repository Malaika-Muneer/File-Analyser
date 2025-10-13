package service

import (
	"github.com/malaika-muneer/File-Analyser/db"
	"github.com/malaika-muneer/File-Analyser/models"
)

type UserService interface {
	SignupUser(user models.User) error
	AuthenticateUser(usernameOrEmail, password string) (*models.User, error)
	UploadFile(fileContent []byte, Username string, id int) (models.FileAnalysis, error)
}
type UserServiceImpl struct {
	Dao db.DaoLayer
}

func NewUserService(dao db.DaoLayer) UserService {
	return &UserServiceImpl{
		Dao: dao,
	}
}
