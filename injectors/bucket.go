package injectors

import (
	"github.com/user0608/expertos/handlers"
	"github.com/user0608/expertos/repository"
	"github.com/user0608/expertos/services"
	"gorm.io/gorm"
)

var ( //db connextion
	db *gorm.DB
	//storage
	loginRepository      *repository.LoginRepository
	estudianteRepository *repository.EstudianteRepository
	casmRepository       *repository.CASMRepository
	bergerRepository     *repository.BergerRepository
	heaRepository        *repository.HEARepository
	testRepository       *repository.TestRepository
	//services
	loginService      *services.LoginService
	estudianteService *services.EstudianteService
	questionService   *services.QuestionService
	testService       *services.TestService
	//handlers
	loginHandler      *handlers.LogginHandler
	estudianteHandler *handlers.EstudianteHandler
	questionHandler   *handlers.QuestionHandler
	testHandler       *handlers.TestHandler
)
