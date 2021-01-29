package general

import (
	"api/config"
	"api/model/projects"
	user2 "api/model/user"
	"api/utlis/jwtParser"
	"api/utlis/password"
	"crypto/tls"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/o1egl/govatar"
	"golang.org/x/crypto/bcrypt"
	gomail "gopkg.in/mail.v2"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

type login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type userLogin struct {
	ID          *int       `json:"id"`
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
	context.JSON(http.StatusOK, endResponse)
}

func UserList(context *gin.Context) {
	users := user2.GetAllUsers()
	context.JSON(http.StatusOK, gin.H{"users": users})
}
func User(context *gin.Context) {
	userObject, err := user2.GetUserById(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Profile doesn't exist"})
		return
	}
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	projects := projects.GetProjectsThatUserIsPartOF(claims["id"])

	context.JSON(http.StatusOK, gin.H{"data": userObject, "projects": projects})
}

func LockUser(context *gin.Context) {
	err := user2.SoftDeleteUser(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error"})
		return
	}
	context.JSON(http.StatusOK, responseLocked{Data: "User is locked"})
}

func UnlockUser(context *gin.Context) {
	err := user2.UnlockUser(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error"})
		return
	}
	context.JSON(http.StatusOK, responseLocked{Data: "User is unlocked"})
}

func NewEmail(context *gin.Context) {
	var request emailRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't unmarshal json"})
		return
	}
	newEmailCreated, response := user2.NewEmail(&request.ID, &request.Email)
	if newEmailCreated {
		context.JSON(http.StatusOK, gin.H{"message": "New Email is updated"})
		return
	} else {
		context.JSON(http.StatusInternalServerError, gin.H{"error": response})
		return
	}
}

func UserSignUp(context *gin.Context) {
	var request user2.User
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(422, gin.H{"error": "Couldn't unmarshal json"})
		return
	}
	if request.Email == "" || request.LastName == "" || request.Name == "" || request.Role == "" {
		context.JSON(422, gin.H{"error": "Fields were not filled"})
		return
	}

	pass := RandStringRunes(12)

	request.Password = password.HashAndSalt([]byte(pass))
	request.Birthday = time.Now()
	request.NameDay = time.Now()
	response, request := user2.CreateUsers(request)
	if response != nil {
		context.JSON(422, gin.H{"error": "User " + request.Email + "wasn't available"})
		return
	}
	id := strconv.Itoa(request.ID)
	os.MkdirAll("storage/avatar/" + id, os.ModeDir)

	err = govatar.GenerateFile(govatar.MALE, "storage/avatar/" + id + "/avatar.jpg")
	if err == nil {
		request.Avatar = "/avatar/" + id + "/avatar.jpg"
		user2.UpdateAvatar(request)
	}
	sendEmail("Hello, we are glad that you have joined our company. Here is your password: " + pass, request.Email)
	context.JSON(200, gin.H{"users": "User created successfully"})
}

func ResetPassword(context *gin.Context) {
	//var temp []string
	//gomail.SendEmailSMTP(append(temp, "vitolinseriks@gmail.com"), "test", "test")
}

func SearchUsers(context *gin.Context) {
	users := user2.SearchUsers(context.Query("condition"))
	context.JSON(200, gin.H{"users": users})
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func sendEmail(content string, email string)  {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "Weeping.Willow177@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", email)

	// Set E-Mail subject
	m.SetHeader("Subject", "News from X")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", content)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "Weeping.Willow177@gmail.com", "LLAOcI0plH2s")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return
}