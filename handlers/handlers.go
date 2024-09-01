package handlers

import (
	database "fiber-api/db"
	"fiber-api/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// DTOs
// @Description DTO para registro de usuário
// @Accept json
// @Produce json
type RegisterDTO struct {
	UsuarioNome  string `json:"usuarioNome"`
	UsuarioLogin string `json:"usuarioLogin"`
	Senha        string `json:"senha"`
}

// @Description DTO para login do usuário
// @Accept json
// @Produce json
type LoginDTO struct {
	UsuarioLogin string `json:"usuarioLogin"`
	Senha        string `json:"senha"`
}

// @Description DTO para resposta de autenticação
// @Accept json
// @Produce json
type AuthResponse struct {
	UsuarioNome       string `json:"usuarioNome"`
	UsuarioLogin      string `json:"usuarioLogin"`
	DataHoraExpiracao string `json:"dataHoraExpiracao"`
	Token             string `json:"token"`
	Autenticado       bool   `json:"autenticado"`
}

const SecretKey = "your_secret_key"

// @Summary Criar um novo usuário
// @Description Crie um novo usuário fornecendo o nome de usuário, login e senha
// @Accept json
// @Produce json
// @Param user body RegisterDTO true "Registro de usuário"
// @Success 201 {object} models.User
// @Router /register [post]
func Register(c *fiber.Ctx) error {
	var dto RegisterDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "não é possível analisar JSON"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Senha), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "não é possível fazer hash da senha"})
	}

	user := models.User{
		UsuarioNome:  dto.UsuarioNome,
		UsuarioLogin: dto.UsuarioLogin,
		Senha:        string(hashedPassword),
	}

	database.DB.Create(&user)

	return c.Status(fiber.StatusCreated).JSON(nil)
}

// @Summary Autenticar um usuário
// @Description Autenticar um usuário com login e senha
// @Accept json
// @Produce json
// @Param credentials body LoginDTO true "Credenciais do usuário"
// @Success 200 {object} AuthResponse
// @Router /login [post]
func Login(c *fiber.Ctx) error {
	var dto LoginDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "não é possível analisar JSON"})
	}

	var user models.User
	database.DB.Where("usuario_login = ?", dto.UsuarioLogin).First(&user)
	if user.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "usuário não encontrado"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Senha), []byte(dto.Senha)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "senha inválida"})
	}

	// Criar o token JWT
	expirationTime := time.Now().Add(time.Hour * 24)
	claims := jwt.MapClaims{
		"usuarioNome":  user.UsuarioNome,
		"usuarioLogin": user.UsuarioLogin,
		"exp":          expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "não é possível gerar token"})
	}

	response := AuthResponse{
		UsuarioNome:       user.UsuarioNome,
		UsuarioLogin:      user.UsuarioLogin,
		DataHoraExpiracao: expirationTime.Format(time.RFC3339),
		Token:             tokenString,
		Autenticado:       true,
	}

	return c.JSON(response)
}
