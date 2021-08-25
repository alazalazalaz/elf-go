##elf-go
快速搭建web项目的脚手架框架，集成了gin、gorm、logrus等组件。

##@todo list
//0. 新增recover异常

//1.新增trace id

//2.新增redis集群

//3. 新增jwt

//test tag

##已有功能
###一、sys配置
####1、配置调试模式
设置`debug: true`或者`false`可开启关闭调试模式。

调试模式开启后会打印所有日志(包括debug/info/warning/error等)。

调试模式关闭后会只会打印info/warning/error日志。

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

###三、redis配置
```
redis:
  ip: 127.0.0.1
  port: 6379
```

