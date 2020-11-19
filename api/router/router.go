package router

import (
	"api/config"
	"api/controllers/general"
	"api/controllers/tags"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
)

func Init() {
	gin.ForceConsoleColor()
	router := gin.Default()
	setupCors(router)
	setupRoutes(router)
	err := router.Run(":8000")
	if err != nil {
		panic(err)
	}
}

func setupCors(router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("authorization", "x-requested-with")
	router.Use(cors.New(config))
}

func setupRoutes(router *gin.Engine) {
	router.POST("/auth/login", general.Login)
	router.Static("/static", "./storage")

	api := router.Group("/api")
	api.Use(jwt.Auth(config.Secret))

	admin := api.Group("/admin", IsAdmin)

	tags := api.Group("/tags")

	adminRoutes(admin)
	apiRoutes(api)
	tagRoutes(tags)
}

func adminRoutes(admin *gin.RouterGroup) {
	admin.POST("/users", general.UserList)

	admin.POST("/user/:id/lock", general.LockUser)
	admin.POST("/user/:id/unlock", general.UnlockUser)
	admin.POST("/user/:id/newEmail", general.NewEmail)
	admin.GET("/user/:id/passwordReset", general.ResetPassword)

	admin.GET("/blog/:id/add", general.AddUserToBlogRole)
	admin.GET("/blog/:id/delete", general.DeleteUserFromBlogRole)
	admin.GET("/users/blog", general.GetAllBloggers)
}

func apiRoutes(api *gin.RouterGroup) {
	api.GET("/user/:id/profile", general.User)
	api.GET("/ping", general.Ping)
	api.GET("/usersonline", general.UsersOnline)
	api.GET("/users/search", general.SearchUsers)
}

func tagRoutes(tag *gin.RouterGroup) {
	tag.GET("/private", tags.IndexPrivate)
	tag.GET("/public", tags.IndexPublic)
	tag.GET("/tag/:id/profile", tags.Profile)
	tag.POST("/tag/:id/join", tags.JoinTag)
	tag.POST("/tag/:id/delete", tags.DeleteTag)
	tag.POST("/tag/:id/newname", tags.NewName)
	tag.POST("/tag/:id/newabout", tags.NewAbout)
	tag.GET("/tag/:id/makePublic", tags.MakePublic)
	tag.GET("/tag/:id/makePrivate", tags.MakePrivate)
	tag.POST("/tag/:id/leave", tags.LeaveTag)
	tag.POST("/tag/:id/setavatar", tags.SetAvatar)
}
