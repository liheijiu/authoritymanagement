package models

import "gorm.io/gorm"

type SysUser struct {
	gorm.Model
	UserName  string `gorm:"column:username;type:varchar(50);" json:"userName"`
	PassWord  string `gorm:"column:password;type:varchar(36);" json:"passWord"`
	Phone     string `gorm:"column:phone;type:varchar(20);" json:"phone"`
	WxUnionId string `gorm:"column:wx_union_id;type:varchar(255);" json:"wxUnionId"`
	WxOpenId  string `gorm:"column:wx_open_id;type:varchar(255);"  json:"wxOpenId"`
	Avatar    string `gorm:"column:avatar;type:varchar(255);" json:"avatar"`
	Sex       string `gorm:"column:sex;type:varchar(20);" json:"sex"`
	Email     string `gorm:"column:email;type:varchar(20);" json:"email"`
	Remarks   string `gorm:"column:remarks;type:varchar(20);" json:"remarks"`
}

// TableName 设置表名称
func (table *SysUser) TableName() string {
	return "sys_user"
}

func GetUserByNamePassWord(userName, passWord string) (*SysUser, error) {
	data := new(SysUser)
	err := DB.Where("username = ? AND password = ?", userName, passWord).First(data).Error
	return data, err
}

// GetUserList 获取管理员数据列表
func GetUserList(keyWord string) *gorm.DB {
	tx := DB.Model(new(SysUser)).Select("id,username,phone,avatar,created_at,updated_at")

	if keyWord != "" {
		tx.Where("username Like ?", "%"+keyWord+"%")
	}
	return tx
}

// GetUserDetail 根据ID获取管理员信息
func GetUserDetail(id uint) (*SysUser, error) {
	su := new(SysUser)
	err := DB.Model(new(SysUser)).Where("id = ?", id).First(su).Error
	return su, err
}
