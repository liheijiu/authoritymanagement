package service

// LoginPassWordRequest 接收登陆参数结构体
type LoginPassWordRequest struct {
	UserName string `json:"userName"` //登陆名称
	PassWord string `json:"passWord"` //登陆密码
}

// LoginPassWordReply 登陆成功后的Token结构体
type LoginPassWordReply struct {
	Token        string `json:"token"`
	ReFreshToken string `json:"reFreshToken"`
}

// GetUserListRequest 获取管理员列表参数的结构体
type GetUserListRequest struct {
	*QueryRequest
}

// QueryRequest 关键字和页面信息结构体
type QueryRequest struct {
	Page    int    `json:"page" form:"page"`
	Size    int    `json:"size" form:"size"`
	KeyWord string `json:"keyWord" form:"keyWord"`
}

// GetUserListReply 返回管理员信息结构体
type GetUserListReply struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// AddUserRequest  接收添加管理员表单数据结构体
type AddUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Remarks  string `json:"remarks"`
}

// GetUserDetailReply 获取管理员信息结构体
type GetUserDetailReply struct {
	ID uint `json:"id"`
	AddUserRequest
}

// UpdateUserRequest 接收更新管理员信息结构体
type UpdateUserRequest struct {
	ID uint `json:"id"`
	AddUserRequest
}

// GetRoleListRequest 获取角色列表查询数据结构体
type GetRoleListRequest struct {
	*QueryRequest
}

// GetRoleListReply 返回角色列表数据
type GetRoleListReply struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Sort      int    `json:"sort"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// AddRoleRequest 新增角色数据结构体
type AddRoleRequest struct {
	Name    string `json:"name"`
	Sort    int64  `json:"sort"`
	IsAdmin int8   `json:"isAdmin"`
	Remarks string `json:"remarks"`
}

// GetRoleDetailReply 返回角色详情信息
type GetRoleDetailReply struct {
	ID uint `json:"id"`
	AddRoleRequest
}

// UpdateRoleRequest 更新角色信息结构体
type UpdateRoleRequest struct {
	ID uint `json:"id"`
	AddRoleRequest
}

// MenuReply 菜单列表数据返回结构体
type MenuReply struct {
	ID            uint         `json:"id"`
	ParentId      uint         `json:"parent_id"`
	Name          string       `json:"name"`
	WebIcon       string       `json:"web_icon"`
	Path          string       `json:"path"`
	Sort          int          `json:"sort"`
	Level         int          `json:"level"`
	ComponentName string       `json:"component_name"`
	SubMenus      []*MenuReply `json:"sub_menus"`
}

// AllMenu 所有菜单的数据结构体
type AllMenu struct {
	ID            uint   `json:"id"`
	ParentId      uint   `json:"parent_id"`
	Name          string `json:"name"`
	WebIcon       string `json:"web_icon"`
	Path          string `json:"path"`
	Sort          int    `json:"sort"`
	Level         int    `json:"level"`
	ComponentName string `json:"component_name"`
}

// 新增菜单结构
type AddMenuRequest struct {
	ParentId      uint   `json:"parent_id"`      //父级菜单唯一标志，不填默认为顶级菜单
	Name          string `json:"name"`           //菜单名称
	WebIcon       string `json:"web_icon"`       //网页图标
	Path          string `json:"path"`           //路径
	Sort          int    `json:"sort"`           //排序
	Level         int    `json:"level"`          //菜单等级 {0:目录，1:菜单，2:按钮}
	ComponentName string `json:"component_name"` //组件路径
}
