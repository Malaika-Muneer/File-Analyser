package service

import (
	"github.com/malaika-muneer/File-Analyser/db"
	"github.com/malaika-muneer/File-Analyser/db/mongodb"
	"github.com/malaika-muneer/File-Analyser/models"
)

// -----------------------------
//  USER SERVICE INTERFACE & IMPL
// -----------------------------

// Defines what the user-related service should provide
type UserService interface {
	SignupUser(user models.User) error
	AuthenticateUser(usernameOrEmail, password string) (*models.User, error)
	UploadFile(fileContent []byte, username string, id int, numChunks int) (map[string]interface{}, error)
}

// Implementation struct for UserService
type UserServiceImpl struct {
	Dao db.DaoLayer
}

// Constructor for UserServiceImpl
func NewUserService(dao db.DaoLayer) UserService {
	return &UserServiceImpl{
		Dao: dao,
	}
}

// Temporary stub for UploadFile (not used directly here, just to satisfy the interface)
func (s *UserServiceImpl) UploadFile(fileContent []byte, username string, id int, numChunks int) (map[string]interface{}, error) {
	return nil, nil
}

// -----------------------------
//  UPLOAD SERVICE INTERFACE & IMPL
// -----------------------------

// Defines the required behavior for UploadService
type UploadServiceInterface interface {
	UploadFile(fileContent []byte, username string, id int, numChunks int) (map[string]interface{}, error)
}

// Struct implementation that interacts with MongoDB
type UploadService struct {
	MongoDAO *mongodb.MongoDAO
}

// Constructor for UploadService
func NewUploadService(mongoDAO *mongodb.MongoDAO) UploadServiceInterface {
	return &UploadService{
		MongoDAO: mongoDAO,
	}
}
