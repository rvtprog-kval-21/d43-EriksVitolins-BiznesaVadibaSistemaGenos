package general

import (
	"api/model/blog"
	"api/model/user"
	"api/utlis/jwtParser"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"time"
)

type BlogsHomeResponse struct {
	Blog   blog.Blogs `json:"blog"`
	Viewed bool       `json:"viewed"`
}

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
	realID, _ := strconv.Atoi(context.Param("id"))
	newBlog.ID = realID
	newBlog.UserID = int(claims["id"].(float64))
	newBlog.Title = context.PostForm("title")
	newBlog.Topic = context.PostForm("topic")
	newBlog.Headtext = context.PostForm("headtext")
	newBlog.Content = context.PostForm("content")
	newBlog.PublishAt, _ = time.Parse("2006-01-02", context.PostForm("publish_at"))
	response := blog.UpdateBlog(&newBlog)
	if response != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": response})
		return
	}
	file, _ := context.FormFile("photo")
	if file != nil {
		path := "/blog/%s"
		path = fmt.Sprintf(path, fmt.Sprint(newBlog.ID))
		if _, err := os.Stat("storage" + path); os.IsNotExist(err) {
			os.MkdirAll("storage"+path, os.ModeDir)
		}
		path = path + "/banner.png"
		err := context.SaveUploadedFile(file, "storage"+path)
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
	}
	context.JSON(201, gin.H{"message": "Article is saved"})
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
	newBlog.PublishAt, _ = time.Parse("2006-01-02", context.PostForm("publish_at"))
	response := blog.AddBlog(&newBlog)
	if response != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error saving the blog"})
		return
	}
	file, _ := context.FormFile("photo")
	path := "/blog/%s"
	path = fmt.Sprintf(path, fmt.Sprint(newBlog.ID))
	if _, err := os.Stat("storage" + path); os.IsNotExist(err) {
		os.MkdirAll("storage"+path, os.ModeDir)
	}
	path = path + "/banner.png"
	err := context.SaveUploadedFile(file, "storage"+path)
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
	context.JSON(201, gin.H{"message": "Blog is created"})
}

func GetYourBlogs(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	blogs, err := blog.GetYourBlogs(claims["id"])
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(200, gin.H{"blogs": blogs})
}

func DeleteBlog(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	blogs := blog.GetBlogUnscoped(context.Param("id"))
	if blogs.UserID != int(claims["id"].(float64)) {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "You're not the owner of this blog"})
		return
	}
	results := blog.DeleteBlog(&blogs)
	if results != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": results})
		return
	}

	context.JSON(200, gin.H{"message": "Blog deleted"})
}

func UndeleteBlog(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	blogs := blog.GetBlogUnscoped(context.Param("id"))
	if blogs.UserID != int(claims["id"].(float64)) {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "You're not the owner of this blog"})
		return
	}
	results := blog.UnDeleteBlog(blogs.ID)
	if results != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": results})
		return
	}

	context.JSON(200, gin.H{"message": "Blog undeleted"})
}

func GetBlog(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	blogs := blog.GetBlogUnscoped(context.Param("id"))
	if blogs.UserID != int(claims["id"].(float64)) {
		if blogs.DeletedAt.Valid {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "This article doesn't exist"})
			return
		} else {
			blog.AddItemToLog(int(claims["id"].(float64)), blogs.ID)
		}
	} else {
		blog.AddItemToLog(int(claims["id"].(float64)), blogs.ID)
	}
	context.JSON(200, gin.H{"blog": blogs})
}

func GetBlogs(context *gin.Context) {
	blogs, err := blog.GetBlogs()
	if err != nil {
		context.JSON(200, gin.H{"error": err})
		return
	}
	context.JSON(200, gin.H{"blogs": blogs})
}

func GetBlogCount(context *gin.Context) {
	count, err := blog.GetCount(context.Param("id"))
	if err != nil {
		context.JSON(200, gin.H{"error": err})
		return
	}
	context.JSON(200, gin.H{"count": count})
}

func GetHomeBlogs(context *gin.Context) {
	blogs, err := blog.GetBlogsLimitedToFour()
	if err != nil {
		context.JSON(200, gin.H{"error": err})
		return
	}
	var response []BlogsHomeResponse
	for _, iter := range blogs {
		response = append(response, BlogsHomeResponse{Blog: iter, Viewed: blog.GetIsViewed(iter.UserID, iter.ID)})
	}

	context.JSON(200, gin.H{"blogs": response})
}

func BlogIsMemeber(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	blogger := blog.Profile(claims["id"])
	if blogger.ID == 0 {
		context.JSON(200, gin.H{"isCreator": false})
		return
	}
	context.JSON(200, gin.H{"isCreator": true})
}
