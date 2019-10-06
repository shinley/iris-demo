package service

import (
	"iris-demo/dao"
	"iris-demo/datasource"
	"iris-demo/model"
)

// 定定接口
type AdminService interface {
	// 通过管理员用户史+密码 获取管理员实体 如果查询到, 控台回管理员实体,并控台回true
	GetByAdminAndPassword(username, password string) (*model.Admin, bool)

	// 获取管理员总数
	GetAdminById(id int) *model.Admin
}

func NewAdminService() AdminService {
	return &adminService{
		adminDao: dao.NewAdminDao(datasource.InstanceMaster()),
	}
}

// 管理员的服务实现结构体
type adminService struct {
	adminDao *dao.AdminDao
}

// 查询管理员总数
func (ac *adminService)GetAdminById(id int) *model.Admin {
	am := ac.adminDao.Get(id)
	return am
}

func (ac *adminService)GetByAdminAndPassword(username, password string)(*model.Admin, bool)  {
	var admin *model.Admin
	admin = ac.adminDao.GetByNaneAndPassword(username, password)
	return admin, true
}