package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"iris-demo/service"
)

type AdminController struct {
	Ctx     iris.Context
	Service service.AdminService
}

func (c *AdminController) Get() mvc.Result {
	id := c.Ctx.URLParam("id")
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

func (c *AdminController) GetEdit() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	var data *models.StarInfo
	if err == nil {
		data = c.Service.Get(id)
	} else {
		data = &models.StarInfo{
			Id: 0,
		}
	}
	//fmt.Println(id, data)
	// set the model and render the view template.
	return mvc.View{
		Name: "admin/edit.html",
		Data: iris.Map{
			"Title": "管理后台",
			"info":  data,
		},
		Layout: "admin/layout.html", // 不要跟前端的layout混用
	}
}
