#fms

#gin + gorm

1.用户登陆验证管理、群组管理

2.故障记录,dashboard功能

3.调取Prometheus接口,动态获取应用状态信息,根据逻辑判断是否为应用故障,自动生成故障信息

4.config为配置文件模块,修改部署地址以及端口、mysql地址、redis地址、监控实例;
  database为mysql和redis连接实例的初始化;
  errorLog为logrus的初始化,错误日志的输出;
  ini调用所有模块的初始化;
  prometheusAPI为获取应用实例状态的模块;
  response为gin请求响应的封装;
  security为前端调用数据的一个处理;
  util为数据的处理、算法实现等;

5.直接部署:
    go build;
    config文件与生成的二进制可执行文件放在同一目录下;
    修改config中的配置项;
    nohup或者./;
    errorLog.txt为错误日志;

6.air热重载编译:
    go mod tidy;
    按需修改.air.conf文件,然后在根目录下执行指令air,修改源码可以实时的重载编译;

#前端部分功能还未开发完全