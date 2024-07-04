# merchant

## 初始化项目工具
* ***请使用linux环境***
* 安装mysqldump
* 将./scripts/exec/sql2struct 执行文件拷贝至环境变量目录下
* 安装swag: 
  ```
  go get -u github.com/swaggo/swag/cmd/swag
  // 引入依赖
  go get -u github.com/swaggo/gin-swagger
  go get -u github.com/swaggo/files
  ```

## 项目结构
```markdown
.
├── bin         //执行文件目录
│   └── configs
├── cmd                    //项目启动目录
│   └── api
├── configs               //配置文件目录，打包命令会拷贝至bin目录
├── internal              //逻辑层目录
│   ├── config             //逻辑配置目录
│   ├── data               //数据层目录
│   │   ├── conn     //数据层链接目录
│   │   └── po       //持久化结构目录，PO:用于持久化数据到数据库
│   ├── logic              //业务逻辑目录
│   │   ├── dto      //传输数据结构目录，DTO: 用于在不同的应用程序组件之间传输数据
│   │   └── nsq      //NSQ消息队列方法实现
│   ├── router             //路由目录
│   │   └── middleware   //中间件目录
│   └── server                    //服务目录
├── log
├── pkg                                 //内部包目录
│   ├── auth
│   ├── bootstrap
│   ├── ginx
│   ├── statusx
│   └── util
└── scripts                           //执行脚本相关目录
├── exec
├── sql
└── template
```
> 参考 : 领域驱动系列-浅析VO、DTO、DO、PO


## 使用说明
> 项目集成内容

* Gin
* Gorm
* Redis
* NSQ
* ClickHouse
* ./stop.sh 支持优雅关闭http服务
* 参数校验: [validator] 直接在参数中配置TAG binding:"required"
* swag `swag init`可以生成swagger文档，http://localhost:8099/swagger/index.html#/ 访问接口文档


> 中间件集成内容

* Oauth2
* 跨域，与nginx配置的跨域冲突
* Casbin鉴权

## Casbin鉴权相关介绍
* ./configs/rbac_model.conf 存放Casbin校验规则
* ./configs/perm_config.json 存放当前目录下所有用户组path及method列表


## 相关命令
> 
> 注 : make执行需要的执行程序详见scripts/exec目录

* 通过传参导出.sql文件: `make sql2file DATABASE=log TABLES=user TABLES=sys_roles`     //TABLES需要传入数据库对应表明，数据库地址及密码在根目录makefile文件中修改 **WARNING : 重复生成会覆盖**
* 生成po目录文件: `make sql2po TABLES=user TABLES=sys_roles`
* 生成dto目录文件: `make sql2dto TABLES=user TABLES=sys_roles`
* 生成增删改查文件: 具体查看logic/template逻辑,请求接口传入对应参数

> 请求参数：`{
   "fun_name": "RoleMenu",  //生成的函数名
   "db_struct_name": "SysRoleMenu", //生成PO目录文件对应的函数名
   "param_name": "rm",  //传参使用的参数名
   "file_name": "role_menu"  //最终生成的文件名
   }`

* 编译文件: `make build`
* 打包文件: `make package`
* 启动服务: 拷贝bin目录至任何位置，或直接进入bin目录。测试环境执行 `./start.sh config.test.yml`