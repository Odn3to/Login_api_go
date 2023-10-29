package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"gorm.io/gorm"

	"Login-api/resources/jwt_func"
	"Login-api/database"
	"Login-api/models"
	"Login-api/resources/User"
)

// @Summary Logar usuário
// @Description Faz login com nome de usuário e senha
// @ID logar
// @Accept  json
// @Produce  json
// @Param   userCredentials     body    User.Login     true        "Credenciais do usuário"
// @Success 200 {object} jwt_func.Verfica
// @Router /login/token [post]
func Logar(c *fiber.Ctx) error {
	user := new(User.Login)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	data := new(models.User)
	db := database.GetDataBase()

	// Find user
	if err := db.Where("username = ? ", user.Username).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid credentials",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error retrieving user",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(user.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	token, err := jwt_func.GenerateJWT(user.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create JWT token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

// @Summary Verifica Jwt Token
// @Description Pega o token e verifica se é válido
// @ID verificadorJwt
// @Accept  json
// @Produce  json
// @Param   token     body    string     true        "Token"
// @Success 200 {object} jwt_func.RetornoVerificar
// @Router /login/validador [post]
func VerificaJwt(c *fiber.Ctx) error {
	request := new(jwt_func.Verfica)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	tokenString := request.Token
	if _, err := jwt_func.VerifyJWT(tokenString); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired JWT",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"ValidToken": true,
	})
}

// @Summary Cria Usuário
// @Description Pega o username e password e cria
// @ID createUser
// @Accept  json
// @Produce  json
// @Param   User.Login     body    User.Login     true        "Credenciais do usuário"
// @Success 200 {object} User.Login
// @Router /login/create [post]
func Create(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
	    fmt.Println("Error generating hash:", err)
	    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	        "error": "Internal server error",
	    })
	}
	
	// Convert byte slice to string
	user.Password = string(hashedPasswordBytes)

	db := database.GetDataBase()

	// Save the user to the database
	if err := db.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": user,
	})
}

// @Summary Deleta Usuário
// @Description Pega o username e senha e exclui
// @ID deleteUser
// @Accept  json
// @Produce  json
// @Param   User.Login     body    User.Login     true        "Credenciais do usuário"
// @Success 200 {object} User.Message
// @Router /login/delete [post]
func Delete(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	db := database.GetDataBase()

	// Retrieve user based on username
	var existingUser models.User
	if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Verify password
	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
	    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	        "error": "Invalid password",
	    })
	}

	// Delete the user from the database
	if err := db.Unscoped().Delete(&existingUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deleted successfully!",
	})
}
