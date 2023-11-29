package models

import "gorm.io/gorm"

type SysRole struct {
	gorm.Model
	Name    string `gorm:"column:name;type:varchar(100);" json:"name"`         //角色名称
	IsAdmin int8   `gorm:"column:is_admin;type:varchar(100);" json:"is_admin"` //是否超管
	Sort    int64  `gorm:"column:sort;type:varchar(100);" json:"sort"`         //排序，序号越小越靠前
	Remarks string `gorm:"column:remarks;type:varchar(20);" json:"remarks"`    //备注
}

// TableName 设置表名称
func (table *SysRole) TableName() string {
	return "sys_role"
}

// GetRoleList 获取角色列表数据
func GetRoleList(keyword string) *gorm.DB {

	tx := DB.Model(new(SysRole)).Select("id,name,is_admin,sort,created_at,updated_at")

	if keyword != "" {
		tx.Where("name LIKE ?", "%"+keyword+"%")

	}
	tx.Order("sort ASC")
	return tx
}

func GetRoleDetail(id uint) (*SysRole, error) {

	sr := new(SysRole)
	err := DB.Model(new(SysRole)).Where("id = ?", id).First(sr).Error
	return sr, err
}
