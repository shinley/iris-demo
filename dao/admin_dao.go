package dao

import (
	"github.com/go-xorm/xorm"
	"iris-demo/model"
)

type AdminDao struct {
	engine *xorm.Engine
}


func NewAdminDao(enine *xorm.Engine) *AdminDao {
	return &AdminDao{
		engine:enine,
	}
}

// 定义dao 方法
func (d *AdminDao) Get(id int) *model.Admin {
	data := &model.Admin{Id:id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}


func (d *AdminDao) GetByNaneAndPassword(name, password string) *model.Admin {
	var admin model.Admin
	d.engine.Where("name = ?", name).And("password=?", password).Get(&admin)
	return &admin
}