package router

import (
	"api/config"
	"api/controllers/calendar"
	"api/controllers/general"
	"api/controllers/notifications"
	"api/controllers/projects"
	"api/controllers/timetable"
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
	notifications := api.Group("/notifications")
	adminRoutes(admin)
	apiRoutes(api)
	projectRoutes(api)
	notificationRoutes(notifications)
	userSettingsRoutes(api)
	userAnnouncementsRoutes(api)
	timetableRoutes(api)
}

func notificationRoutes(notifi *gin.RouterGroup) {
	notifi.GET("/get/count/current", notifications.NotifsAvailable)
	notifi.GET("/get/existing/current", notifications.Notifs)
	notifi.GET("/update/:id/notification/seen", notifications.UpdateSeenToTrue)
}

func adminRoutes(admin *gin.RouterGroup) {
	admin.POST("/users", general.UserList)
	admin.POST("/signup", general.UserSignUp)

	admin.POST("/user/settings/:id/lock", general.LockUser)
	admin.POST("/user/settings/:id/unlock", general.UnlockUser)
	admin.POST("/user/settings/:id/newEmail", general.NewEmail)
	admin.POST("/user/reset/password", general.ResetPassword)

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
	api.POST("/users", general.UserList)
	api.POST("/users/search", general.SearchUser)
	api.GET("/account/check/", general.SeeIfUserInitiated)
	api.GET("/account/init/", general.InitUser)
	api.POST("/users/get/multiple", general.GetUsersMultiple)
	blogRoutes(api)
	managerRoutes(api)
	trackingRoutes(api)
	calendarRoutes(api)
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

func calendarRoutes(api *gin.RouterGroup) {
	calander := api.Group("/calendar")
	calander.POST("/create/new/event/", calendar.CreateEvent)
	calander.POST("/get/current/events/", calendar.GetEvents)
	calander.GET("/delete/:id/event", calendar.DeleteEvent)
	calander.GET("/leave/:id/event", calendar.LeaveEvent)
	calander.POST("/update/:id/event", calendar.UpdateEvent)
	calander.POST("/get/home/events/", calendar.GetEventsHome)
}

func managerRoutes(api *gin.RouterGroup) {
	manager := api.Group("/manager")
	manager.GET("/check/ismanager/", general.ManagerIsMemeber)
	manager.GET("/list/", general.SeeSubmissions)
}

func timetableRoutes(api *gin.RouterGroup) {
	manager := api.Group("/timetable")
	manager.POST("/save/schedule", timetable.Save)
	manager.POST("/get/personal/schedule", timetable.Get)
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
	project.POST("/change/:id/tags/", projects.ChangeTags)
	project.POST("/update/:id/tags/", projects.UpdateTags)
	project.POST("/create/:id/new/announcement", projects.SaveAnnouncement)
	project.GET("/see/:id/current/announcement", projects.SeeAnnouncements)
	project.GET("/gather/members/tags/list", projects.GetTagsOfMemberProject)
}

func userSettingsRoutes(api *gin.RouterGroup) {
	user := api.Group("/user/settings/new")
	user.POST("/name", general.UpdateUserName)
	user.POST("/birthday", general.UpdateUserBirthday)
	user.POST("/about", general.UpdateUserAbout)
	user.POST("/avatar", general.UpdateUserAvatar)
	user.POST("/background", general.UpdateUserBackground)
	user.POST("/number", general.UpdateUserNumber)
	user.POST("/password", general.UpdateUserPassword)
	user.POST("/title", general.UpdateUserTitle)
}

func userAnnouncementsRoutes(api *gin.RouterGroup) {
	user := api.Group("/user")
	user.POST("/announcements/new/announcements", general.NewUserAnnc)
	user.POST("/announcements/get/user", general.UsersAnnc)
	user.POST("/announcements/gets/follow", general.FollowedAnnc)
	user.POST("/announcements/delete/announcement", general.DeleteAnnc)
	user.POST("follower/start", general.AddFollower)
	user.POST("follower/delete", general.DeleteFollower)
	user.POST("follower/check", general.SeeIfFollowing)
}