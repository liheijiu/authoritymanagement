package main

import (
	"go-Admin/models"
	"go-Admin/router"
)

func main() {
	//mysql初始化
	models.NewGormDB()
	r := router.App()
	r.Run(":8081")
}

/*

创建web：
1.npm create vite@latest
2.项目名名字
3. 回车
4. vue
5. TypeScript
6. npm install  依赖包
7. 运行 npm  run dev

8.npm  install vue-router@4
9. npm  install pinia
10.npm  install pinia-plugin-persistedstate
11. npm install element-plus
12. npm install axios
13. npm  install vite-plugin-svg-icons -D  若报错，则 npm  install fast-glob  -D



*/
