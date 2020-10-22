package user_controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang-api/config"
	"golang-api/database"
	user_model "golang-api/model"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginError struct {
	Error string `json:"error"`
}

type responseLogin struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	Avatar      string `json:"avatar"`
	Name        string `json:"name"`
	Lastname    string `json:"lastname"`
	AccessToken string `json:"access_token"`
}

func Login(context *fiber.Ctx) error {
	request := new(login)
	if err := context.BodyParser(request); err != nil {
		return context.Status(500).JSON(loginError{Error: "Coudn't unmarshal json"})
	}
	database.Open()
	user, err := user_model.FindByEmail(request.Email)
	if err != nil {
		return context.Status(401).JSON(loginError{Error: "Email or Password is wrong"})
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return context.Status(401).JSON(loginError{Error: "Email or Password is wrong"})
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["id"] = user.ID
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 9).Unix()
	realToken, err := token.SignedString([]byte(config.SECRET))
	if err != nil {
		return context.SendStatus(fiber.StatusInternalServerError)
	}
	endResponse := responseLogin{
		ID:          user.ID,
		Email:       user.Email,
		Role:        user.Role,
		Avatar:      user.Avatar,
		Name:        user.Name,
		Lastname:    user.LastName,
		AccessToken: realToken,
	}
	database.Close()
	return context.JSON(endResponse)
}

func Index(c *fiber.Ctx) error {
	return c.SendString("Hello, World 1!")
}
func User(c *fiber.Ctx) error {
	return c.SendString("Hello, World 2!")
}
