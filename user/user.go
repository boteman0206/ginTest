package user

type User struct {

	//todo json定义的是json的格式数据
	// todo form表单绑定的字段
	UserName string `form:"name" json:"user_name" uri:"name"`
	UserAge  int    `form:"age" json:"user_age" uri:"age"`
	Address  string `form:"address" json:"address"`
}
