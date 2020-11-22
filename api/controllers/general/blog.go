package general

import (
	"api/model/blog"
	"api/model/user"
	"api/utlis/jwtParser"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
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
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	blogger := blog.Profile(claims["id"])
	if blogger.ID == 0 {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "You don't have access to make this change"})
		return
	}
	var newBlog blog.Blogs
	newBlog.UserID = int(claims["id"].(float64))
	newBlog.Title = context.PostForm("title")
	newBlog.Topic = context.PostForm("topic")
	newBlog.Headtext = context.PostForm("headtext")
	newBlog.Content = context.PostForm("content")
	newBlog.PublishAt, _ = time.Parse("2006-01-02",context.PostForm("publish_at"))
	response := blog.AddBlog(&newBlog)
	if response != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error saving the blog"})
		return
	}
	file, _ := context.FormFile("photo")
	path := "storage/blog/%s"
	path = fmt.Sprintf(path,  fmt.Sprint(newBlog.ID))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModeDir)
	}
	path = path + "/banner.png"
	err := context.SaveUploadedFile(file,path)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error saving the banner"})
		return
	}
	newBlog.Photo = path
	response = blog.UpdateBannerBlog(&newBlog)
	if response != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error saving the blog banner"})
		return
	}
	context.JSON(201, gin.H{"message": "Article is created"})
}
