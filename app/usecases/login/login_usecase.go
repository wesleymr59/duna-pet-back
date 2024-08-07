package usecases

import (
	"duna-pet-back/app/entities"
	interfaces "duna-pet-back/app/interfaces/login"
	"duna-pet-back/app/utils"
	"errors"
	"fmt"
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

	userFromDB, err := s.userRepository.FindByEmailAndPassword(userLogin.Email, userLogin.Passwd)
	if err != nil {
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
	fmt.Println("ta aqui carai")
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return entities.User{}, errors.New("binding JSON failed")
	}

	hashPasswd := utils.CreateHash(user.Passwd)
	user.Passwd = hashPasswd

	if err := s.userRepository.CreateUser(&user); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "email already registered"})
		return entities.User{}, errors.New("email already registered")
	}
	return user, nil
}
