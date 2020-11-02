package user

import (
	"api/config"
	"api/database"
	user2 "api/model/user"
	"api/services/gomail"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

type emailRequest struct {
	Email string `json:"email"`
	ID    string `json:"id"`
}

type responseIndex struct {
	Data *[]user2.User `json:"data"`
}

type responseUser struct {
	Data *user2.User `json:"data"`
}

type responseLocked struct {
	Data string `json:"data"`
}

func Login(context *gin.Context) {
	var request login
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't unmarshal json"})
		return
	}
	database.Open()
	userObject, err := user2.FindByEmail(request.Email)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Email or Password is wrong"})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(userObject.Password), []byte(request.Password))
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Email or Password is wrong"})
		return
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = userObject.Email
	claims["id"] = userObject.ID
	claims["role"] = userObject.Role
	claims["exp"] = time.Now().Add(time.Hour * 9).Unix()
	realToken, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't unmarshal json"})
		return
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
	context.JSON(http.StatusOK, endResponse)
}

func Index(context *gin.Context) {
	database.Open()
	users := user2.GetAllUsers()
	database.Close()
	context.JSON(http.StatusOK, responseIndex{Data: &users})
}
func User(context *gin.Context) {
	database.Open()
	userObject, err := user2.GetUserById(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Profile doesn't exist"})
		return
	}
	database.Close()
	context.JSON(http.StatusOK, responseUser{Data: userObject})
}

func LockUser(context *gin.Context) {
	database.Open()
	err := user2.SoftDeleteUser(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error"})
		return
	}
	database.Close()
	context.JSON(http.StatusOK, responseLocked{Data: "User is locked"})
}

func UnlockUser(context *gin.Context) {
	database.Open()
	err := user2.UnlockUser(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error"})
		return
	}
	database.Close()
	context.JSON(http.StatusOK, responseLocked{Data: "User is unlocked"})
}

func NewEmail(context *gin.Context) {
	var request emailRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't unmarshal json"})
		return
	}
	database.Open()
	newEmailCreated, response := user2.NewEmail(&request.ID, &request.Email)
	database.Close()
	if newEmailCreated {
		context.JSON(http.StatusOK, gin.H{"message": "New Email is updated"})
		return
	} else {
		context.JSON(http.StatusInternalServerError, gin.H{"error": response})
		return
	}
}

func ResetPassword(context *gin.Context) {
	var temp []string
	gomail.SendEmailSMTP(append(temp, "vitolinseriks@gmail.com"), "test", "test")
}
