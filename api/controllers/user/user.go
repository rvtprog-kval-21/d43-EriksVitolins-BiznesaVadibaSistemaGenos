package user_controller

import (
	"api/config"
	"api/database"
	user "api/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
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

type userLogin struct {
	ID          *uint      `json:"id"`
	Email       *string    `json:"email"`
	Role        *string    `json:"role"`
	Avatar      *string    `json:"avatar"`
	Name        *string    `json:"name"`
	Lastname    *string    `json:"lastname"`
	AccessToken *string    `json:"access_token"`
	About       *string    `json:"about"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type responseLogin struct {
	Data userLogin `json:"data"`
}

type responseIndex struct {
	Data *[]user.User `json:"data"`
}

type responseUser struct {
	Data *user.User `json:"data"`
}

type responseLocked struct {
	Data string `json:"data"`
}

func Login(context *fiber.Ctx) error {
	request := new(login)
	if err := context.BodyParser(request); err != nil {
		return context.Status(500).JSON(loginError{Error: "Coudn't unmarshal json"})
	}
	database.Open()
	userObject, err := user.FindByEmail(request.Email)
	if err != nil {
		return context.Status(401).JSON(loginError{Error: "Email or Password is wrong"})
	}
	err = bcrypt.CompareHashAndPassword([]byte(userObject.Password), []byte(request.Password))
	if err != nil {
		return context.Status(401).JSON(loginError{Error: "Email or Password is wrong"})
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = userObject.Email
	claims["id"] = userObject.ID
	claims["role"] = userObject.Role
	claims["exp"] = time.Now().Add(time.Hour * 9).Unix()
	realToken, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return context.SendStatus(fiber.StatusInternalServerError)
	}
	endResponse := userLogin{
		ID:          &userObject.ID,
		Email:       &userObject.Email,
		Role:        &userObject.Role,
		Avatar:      &userObject.Avatar,
		Name:        &userObject.Name,
		Lastname:    &userObject.LastName,
		AccessToken: &realToken,
	}
	database.Close()
	return context.JSON(responseLogin{Data: endResponse})
}

func Index(context *fiber.Ctx) error {
	database.Open()
	users := user.GetAllUsers()
	database.Close()
	return context.JSON(responseIndex{Data: &users})
}
func User(context *fiber.Ctx) error {
	database.Open()
	userObject, err := user.GetUserById(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(loginError{Error: "Profile doesn't exist"})
	}
	database.Close()
	return context.JSON(responseUser{Data: userObject})
}

func LockUser(context *fiber.Ctx) error {
	database.Open()
	err := user.SoftDeleteUser(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(loginError{Error: "There was an error"})
	}
	database.Close()
	return context.JSON(responseLocked{Data: "User is locked"})
}

func UnlockUser(context *fiber.Ctx) error {
	database.Open()
	err := user.UnlockUser(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(loginError{Error: "There was an error"})
	}
	database.Close()
	return context.JSON(responseLocked{Data: "User is unlocked"})
}
