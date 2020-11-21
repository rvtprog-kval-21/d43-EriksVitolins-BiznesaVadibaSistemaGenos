package general

import (
	"api/model/blog"
	"api/model/user"
	"api/utlis/jwtParser"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddUserToBlogRole(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	response := blog.JoinRole(context.Param("id"))
	if response != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": response})
		return
	}
	context.JSON(201, gin.H{"message": "Congrats user is now a blogger"})
}

func DeleteUserFromBlogRole(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	response := blog.DeleteRole(context.Param("id"))
	if response != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": response})
		return
	}
	context.JSON(201, gin.H{"message": "User isn't a blogger anymore"})
}

func GetAllBloggers(context *gin.Context) {
	bloggers := blog.All()
	var users []user.User
	for _, iter := range bloggers {
		users = append(users, iter.User)
	}

	context.JSON(201, gin.H{"users": users})
}

func UpdateBlog(context *gin.Context) {

}

func AddBlog(context *gin.Context) {

}
