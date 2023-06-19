#### 项目启动和停止命令

```shell
#启动门户网站服务
go run main.go start app --host=0.0.0.0 --port=8080
#停止门户网站服务
go run main.go stop app 

#启动后台管理服务
go run main.go start admin --host=0.0.0.0 --port=9090
#停止后台管理服务
go run main.go stop admin 
```
