package router

import (
	"api/config"
	"api/controllers/general"
	"api/controllers/projects"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	gin.ForceConsoleColor()
	router := gin.Default()
	setupCors(router)
	setupRoutes(router)
	return router
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

	adminRoutes(admin)
	apiRoutes(api)
	projectRoutes(api)
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

	admin.GET("/users/managers", general.GetAllManagers)
	admin.GET("/manager/:id/add", general.AddUserToManagerRole)
	admin.GET("/manager/:id/delete", general.DeleteUserFromManagerRole)
}

func apiRoutes(api *gin.RouterGroup) {
	api.GET("/user/:id/profile", general.User)
	api.GET("/ping", general.Ping)
	api.GET("/usersonline", general.UsersOnline)
	api.GET("/users/search", general.SearchUsers)

	blogRoutes(api)
	managerRoutes(api)
	trackingRoutes(api)

}

func blogRoutes(api *gin.RouterGroup) {
	blog := api.Group("/blog")
	blog.POST("/add", general.AddBlog)
	blog.POST("/update/:id", general.UpdateBlog)
	blog.GET("/owner", general.GetYourBlogs)
	blog.GET("/owner/get/:id/delete", general.DeleteBlog)
	blog.GET("/owner/get/:id/undelete", general.UndeleteBlog)
	blog.GET("/get/:id/", general.GetBlog)
	blog.GET("/all/", general.GetBlogs)
	blog.GET("/get/:id/count", general.GetBlogCount)
	blog.GET("/home/limited", general.GetHomeBlogs)
	blog.GET("/check/isowner/", general.BlogIsMemeber)
}

func trackingRoutes(api *gin.RouterGroup) {
	tracking := api.Group("/tracking")
	tracking.POST("/add/", general.AddSubmission)
	tracking.POST("/add/attachment/", general.AddAttachments)
	tracking.GET("/user/list", general.SeePersonalSubmissions)
	tracking.GET("/manager/check/ismanager/", general.ManagerIsMemeber)
	tracking.GET("/manager/list/", general.SeeSubmissions)
	tracking.GET("/manager/item/:id/open", general.OpenSubmissions)
	tracking.GET("/manager/item/:id/close", general.CloseSubmissions)
}

func managerRoutes(api *gin.RouterGroup) {
	manager := api.Group("/manager")
	manager.GET("/check/ismanager/", general.ManagerIsMemeber)
	manager.GET("/list/", general.SeeSubmissions)
}

func projectRoutes(api *gin.RouterGroup) {
	project := api.Group("/projects")
	project.POST("/add/new", projects.CreateProject)
	project.GET("/all/", projects.GetAll)
	project.GET("/get/:id/item/", projects.GetProject)
	project.GET("/add/:id/admin/", projects.MakeAdmin)
	project.GET("/remove/:id/admin/", projects.UnMakeAdmin)
	project.GET("/kick/:id/member/", projects.KickUser)
	project.GET("/leave/:id/member/", projects.LeaveProject)
	project.GET("/remove/:id/project/", projects.ArchiveProject)
	project.POST("/change/:id/about/", projects.ChangeAbout)
	project.POST("/change/:id/name/", projects.ChangeName)
	project.POST("/change/:id/avatar/", projects.ChangeAvatar)
	project.GET("/get/:id/nonmembers/", projects.GetNonMembers)
	project.POST("/invite/:id/users/", projects.AddUsers)
}
