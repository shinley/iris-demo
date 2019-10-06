package model

import "time"

type Admin struct {
	// 如果fieldId名称为Id, 而且类型为int64, 并没有定义tag, 则会被xorm视为主键,并且拥有自增属性
	Id int					`xorm:"pk autoincr" json:"id"`
	AdminName string 		`xorm:"varchar(32)" json:"admin_name"`
	CreateTime time.Time 	`xorm:"DateTime" json:"create_time"`
	Status int				`xorm:"default 0" json:"status"`
	Avatar string			`xorm:"varchar(255)" json:"avatar"`
	Pwd		string			`xorm:"varchar(255)" json:"pwd"`
	CityName	string		`xorm:"varchar(12)" json:"city_name"`
	CityId	string 			`xorm:"index" json:"city_id"`
	//City	*City			`xorm:"- <- ->"`
}