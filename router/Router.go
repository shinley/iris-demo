package routes

import (
	"github.com/kataras/iris/mvc"
	"iris-demo/bootstrap"
	"iris-demo/controller"
	"iris-demo/service"
	//"github.com/yz124/superstar/web/middleware"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {
	//superstarService := service.NewAdminService()
	//
	//index := mvc.New(b.Party("/"))
	//index.Register(superstarService)
	//index.Handle(new(controller.IndexController))

	adminService := service.NewAdminService()
	admin := mvc.New(b.Party("/admin"))
	//admin.Router.Use(middleware.BasicAuth)
	admin.Register(adminService)
	admin.Handle(new(controller.AdminController))

	//b.Get("/follower/{id:long}", GetFollowerHandler)
	//b.Get("/following/{id:long}", GetFollowingHandler)
	//b.Get("/like/{id:long}", GetLikeHandler)
}
