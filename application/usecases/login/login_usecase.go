package usecases

import (
	database "duna-pet-back/adapters/infrastructure/mySql/config"
	interfaces "duna-pet-back/application/interfaces/login"
	"duna-pet-back/application/utils"
	"duna-pet-back/domain/entities"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginUsecase struct {
	userRepository interfaces.UserRepository
}

func NewLoginService(repo interfaces.UserRepository) *LoginUsecase {
	return &LoginUsecase{userRepository: repo}
}

func (s *LoginUsecase) LoginUser(c *gin.Context) (string, error) {
	var userLogin entities.Login
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return "", errors.New("binding JSON failed")
	}

	hashPasswd := utils.CreateHash(userLogin.Passwd)
	userLogin.Passwd = hashPasswd
	log.Println(userLogin)
	var userFromDB entities.User
	if result := database.DB.Where("email = ? AND passwd = ? AND active = ?", userLogin.Email, userLogin.Passwd, 1).First(&userFromDB); result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return "", errors.New("binding GET USER failed")
	}
	// Criar token JWT
	token, err := utils.CreateJWTToken(userFromDB.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar token JWT"})
		return "", errors.New("binding TOKEN failed")
	}
	return token, nil

}

func (s *LoginUsecase) RegisterUser(c *gin.Context) (entities.User, error) {
	var user entities.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return entities.User{}, errors.New("binding JSON failed")
	}

	hashPasswd := utils.CreateHash(user.Passwd)
	user.Passwd = hashPasswd
	log.Println(hashPasswd)
	if result := database.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusConflict, "email already registered")
		return entities.User{}, errors.New("email already registered")
	}
	return user, nil

}
