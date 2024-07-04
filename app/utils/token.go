package utils

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var env_data = GetConfig()

var secretKey = []byte(env_data.SecretKey)

func CreateJWTToken(userID int) (string, error) {
	// Criar token com payload contendo o ID do usuário e expiração
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(), // Expira em 1 hora, ajuste conforme necessário
	})

	// Assinar o token com a chave secreta
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyTokenMiddleware(c *gin.Context) {
	// Extrair o token do cabeçalho Authorization
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token ausente"})
		c.Abort()
		return
	}

	// Remover o prefixo "Bearer " do token (se presente)
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	// Validar o token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verificar o método de assinatura
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Método de assinatura inválido: %v", token.Header["alg"])
		}

		return secretKey, nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		c.Abort()
		return
	}

	// Verificar se o token está válido
	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		c.Abort()
		return
	}

	// Adicionar os detalhes do token ao contexto (opcional)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		c.Abort()
		return
	}

	c.Set("userID", claims["user_id"])

	// Continuar para a próxima rota
	c.Next()
}

func RefreshToken(c *gin.Context) {
	// Extrair o token do corpo da requisição
	var request struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de requisição inválido"})
		return
	}

	// Remover o prefixo "Bearer " do token (se presente)
	refreshToken := strings.Replace(request.RefreshToken, "Bearer ", "", 1)

	// Validar o token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		// Verificar o método de assinatura
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Método de assinatura inválido: %v", token.Header["alg"])
		}

		return secretKey, nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token inválido"})
		return
	}

	// Extrair os claims do token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token inválido"})
		return
	}

	// Criar novo token
	newToken, err := CreateJWTToken(int(claims["user_id"].(float64)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar novo token"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar novo token"})
		return
	}

	// Retornar o novo token
	c.JSON(http.StatusOK, gin.H{"token": newToken})
}
