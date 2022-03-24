##elf-go
快速搭建web项目的脚手架框架，集成了gin、gorm、logrus等组件。

##@todo list
//0. 新增recover异常

//1.新增trace id，有并发问题，需要修改

//2.新增redis集群

// 3. printRequest中间件打印输出
//test tag

##已有功能
###一、sys配置
####1、配置调试模式
设置`debug: true`或者`false`可开启关闭调试模式。

调试模式暂时无实际功能。

###二、mysql配置
支持多db、多主从。
案例如下：
```
mysql:
  default:
    master: #主库
      - #可以配置多个主库
        ip: 127.0.0.1
        port: 3306
        user: root
        password: 123456
        db: test
    slaver: #从库
      - #可以配置多个从库
        ip: 127.0.0.1
        port: 3306
        user: root
        password: 123456
        db: test1
      - #可以配置多个从库
        ip: 127.0.0.1
        port: 3306
        user: root
        password: 123456
        db: test2
  backend:
    master: #主库
      - #可以配置多个主库
        ip: 127.0.0.1
        port: 3306
        user: root
        password: 123456
        db: test3
    slaver: #从库
      - #可以配置多个从库
        ip: 127.0.0.1
        port: 3306
        user: root
        password: 123456
        db: test4
```

支持mysql的hook，依赖gorm的hook，详见example/app/dao/plugin/plugin.go
###三、redis配置
```
redis:
  ip: 127.0.0.1
  port: 6379
```

###四、输出
统一输出json，统一输出字段格式，详见utils/helper/echo.go

###五、JWT
支持JWT标准使用，使用方法：

#### 生成JWT
生成JWT方法如下：
1. 调用`jwts.CreateJwtToken(map[string]interface{})`方法即可。@todo参数支持map[string]interface{}

#### 解析JWT
```golang
authRouter := router.Group("auth")//创建一个需要JWT认证的group
authRouter.Use(middleware2.ParseJwt)//引入JWT中间件
authRouter.GET("", handler.Auth)//将需要JWT认证的路由都挂载到该group下
```

`middleware2.ParseJwt`中间件做了以下操作来判断是否认证通过：
1. 读取请求header中的`Authorization`的值，去掉值中`Bearer `前缀，后面的值为待验证的JWT token。
2. 用公钥解析token，通过即可。

###六、log
log配置，可增加如下`yaml`配置来管理日志。
```yaml
#logrus日志
logrus:
  level: 1 #设定日志等级,详见logrus.Level，非正常值会默认level=logrus.InfoLevel
  WriteToFilePath: /Users/zhangxiong/work/logs/elf.log # 输出到文件，配置为空则不输出到文件
```
