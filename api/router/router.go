package router

import (
	"api/config"
	"api/controllers/tags"
	"api/controllers/user"
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
	router.POST("/auth/login", user.Login)
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
	admin.POST("/users", user.Index)
	admin.POST("/user/:id/lock", user.LockUser)
	admin.POST("/user/:id/unlock", user.UnlockUser)
	admin.POST("/user/:id/newEmail", user.NewEmail)
	admin.GET("/user/:id/passwordReset", user.ResetPassword)
}

func apiRoutes(api *gin.RouterGroup) {
	api.GET("/user/:id/profile", user.User)
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
