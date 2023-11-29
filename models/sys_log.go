package models

import "gorm.io/gorm"

// SysLog 日志管理
type SysLog struct {
	gorm.Model
	Username     string `gorm:"column:username;type:varchar(255);" json:"username"`
	Name         string `gorm:"column:name;type:varchar(255);" json:"name"`
	IP           string `gorm:"column:ip;type:varchar(200);" json:"ip"`
	State        string `gorm:"column:state;type:varchar(200);" json:"state"`
	Request      string `gorm:"column:request;type:varchar(200);" json:"request"`
	Path         string `gorm:"column:path;type:varchar(200);" json:"path"`
	ReqParameter string `gorm:"column:req_parameter;type:varchar(200);" json:"req_parameter"`
	Reply        string `gorm:"column:reply;type:varchar(200);" json:"reply"`
	Take         int    `gorm:"column:take;type:int(200);" json:"take"`
	Browser      string `gorm:"column:browser;type:varchar(200);" json:"browser"`
}

// TableName 设置数据表名称
func (table *SysLog) TableName() string {
	return "sys_log"
}
