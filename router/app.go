package router

import (
	"github.com/gin-gonic/gin"
	"go-Admin/middleware"
	"go-Admin/service"
)

func App() *gin.Engine {
	r := gin.Default()

	//添加跨域中间件
	r.Use(middleware.Cors(), middleware.Log(), gin.Logger())

	//根据用户名和密码登陆路由
	r.POST("/login/password", service.LoginPassWord)

	//管理 开始
	//管理员列表
	r.GET("/user", service.GetUserList)
	//新增管理员信息
	r.POST("/user", service.AddUser)
	//获取管理员信息
	r.GET("/user/detail/:id", service.GetUserDetail)
	//修改管理员信息
	r.PUT("/user", service.UpdateUser)
	//删除管理员
	r.DELETE("/user/:id", service.DeleteUser)
	//管理 结束

	//角色管理开始
	//角色列表
	r.GET("/role", service.GetRoleList)
	//添加角色
	r.POST("/role", service.AddRole)
	//获取角色详情
	r.GET("/role/detail/:id", service.GetRoleDetail)
	//更新角色信息
	r.PUT("/role", service.UpdateRole)
	//删除角色信息
	r.GET("/role/:id", service.DeleteRole)
	//修改角色管理信息
	r.PATCH("/role/:id/:isAdmin", service.PatchRoleAdmin)
	//角色管理结束

	//菜单管理 开始
	//获取菜单列表
	r.GET("/menu", service.GetMenuList)
	//新增菜单
	r.POST("/menu", service.AddMenu)
	//菜单管理结束

	return r
}
