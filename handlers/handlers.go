package handlers

import (
	"fiber-api/database"
	"fiber-api/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// DTOs
// @Description DTO for user registration
// @Accept json
// @Produce json
type RegisterDTO struct {
	UsuarioNome  string `json:"usuarioNome"`
	UsuarioLogin string `json:"usuarioLogin"`
	Senha        string `json:"senha"`
}

// @Description DTO for user login
// @Accept json
// @Produce json
type LoginDTO struct {
	UsuarioLogin string `json:"usuarioLogin"`
	Senha        string `json:"senha"`
}

// @Description DTO for authentication response
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

// @Summary Create a new user
// @Description Create a new user by providing the username, login, and password
// @Accept json
// @Produce json
// @Param user body RegisterDTO true "User Registration"
// @Success 201 {object} models.User
// @Router /register [post]
func Register(c *fiber.Ctx) error {
	var dto RegisterDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Senha), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot hash password"})
	}

	user := models.User{
		UsuarioNome:  dto.UsuarioNome,
		UsuarioLogin: dto.UsuarioLogin,
		Senha:        string(hashedPassword),
	}

	database.DB.Create(&user)

	return c.Status(fiber.StatusCreated).JSON(nil)
}

// @Summary Authenticate a user
// @Description Authenticate a user with login and password
// @Accept json
// @Produce json
// @Param credentials body LoginDTO true "User Credentials"
// @Success 200 {object} AuthResponse
// @Router /login [post]
func Login(c *fiber.Ctx) error {
	var dto LoginDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	var user models.User
	database.DB.Where("usuario_login = ?", dto.UsuarioLogin).First(&user)
	if user.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Senha), []byte(dto.Senha)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid password"})
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot generate token"})
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
