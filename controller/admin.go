package controller

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"iris-demo/service"
)

type AdminController struct {
	Ctx     iris.Context
	Service service.AdminService
}

func (c *AdminController) Get() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("id:", id)

	datalist := c.Service.GetAdminById(id)
	// set the model and render the view template.
	return mvc.View{
		Name: "admin/index.html",
		Data: iris.Map{
			"Title":    "管理后台",
			"Datalist": datalist,
		},
		Layout: "admin/layout.html", // 不要跟前端的layout混用
	}
}
