# legend-admin
基于beego可快速构建的后台系统


1. 基于beego和layui. 用于快速构建后台的页面. 使用命令行命令快速生成对应的文件. 快速实现了增删改查的功能
2. 快速构建方法:
    1. 创建model文件, 例如: user. 创建相应模型struct结构, init里注册orm和builder.
    2. 模型结构添加tag form:"表单名,类型,提示名,Options[|OptionsUrl]", 类型支持: text, textarea, select, password, Options是select参数, 表示options来源js常量, 可参考static/js/macro_option_func.js. OptionsUrl表示options来源ajax OPTIONS. 以路由方式请求, 对应的模型需要实现options方法
    3. 使用命令 go run main.go --op=gen -m={model}.  具体的模型名称填写注册builder的名称, 将会在json目录下生成对应的JSON配置文件. 请修改title字段.
    4. 使用命令 go run main.go --op=make -m={model}. 该命令将生成对应的model_function, controller, router, view文件.

```golang
type User struct{
    Base
    Id            int
    Username      string    `orm:"unique;size(32)" form:"Username,text,用户名" valid:"Required;MaxSize(32);MinSize(3)"`
    Password      string    `orm:"size(32)" form:"Password,password,密码" json:"-" valid:"Required;MaxSize(32);MinSize(3)"`
    Nickname      string    `orm:"unique;size(32)" form:"Nickname,text,昵称" valid:"Required;MaxSize(32);MinSize(3)"`
    State 		  int8		`form:"State,select,状态,Options:StateSelector"`
}

// 需要一个init方法注册相应的结构体 会自动变成小写的表名 如user
func init()  {
    // 需要注册model, 在gen的时候使用name来索引相应的结构体
    registerBuildModel("user", new(User))
    orm.RegisterModel(new(User))
}
```