# go-task

## 远程日志

### 上报ELK

以sendmessage服务为例
```go
func NewSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendLogic {
	return &SendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		//Logger: logx.WithContext(ctx),
		Logger: log.NewGoZeroELKLoggerWithContext(ctx, log.WithAppName("sendmessage"), log.WithFuncName("Send")),
	}
}
```

### 查询日志
http://127.0.0.1:5601/app/discover

通过KQL语句` @timestamp > now-5m`查询最近5分钟的日志


## 远程配置

新增/修改 Apollo配置

以添加发送名单为例

![设置apollo名单](README.assets/apollo.png)


## go-task模块介绍

![go-task](README.assets/go-task.png)


## 分布式任务队列

https://github.com/hibiken/asynq



docker run --rm --name asynqmon -d -p 8098:8080 -e REDIS_ADDR=172.17.0.1:6379 hibiken/asynqmon

Web页面: http://127.0.0.1:8098/



## http反向代理
在director中通过北极星（服务发现）完成目标节点信息获取

```go
getOneRequest := &api.GetOneInstanceRequest{}
getOneRequest.Namespace = namespace
getOneRequest.Service = serviceName
oneInstResp, err := gConsumer.GetOneInstance(getOneRequest)
if err != nil {
  errMsg = fmt.Sprintf("fail to getOneInstance, err is %v", err)
  return
}
instance := oneInstResp.GetInstance()
if instance == nil {
  errMsg = "no instance"
  return
}
```




## 服务发现与治理

### 心跳上报异常[TODO]



### 北极星服务发现

反向代理中获取实例



## 链路追踪 [TODO]



## 整合主流框架及组件 -- rk-boot

### 框架集成指标上报
rk-boot

https://github.com/rookie-ninja/rk-zero



### 手动上报重要指标
以manager/internal/metrics/metrics.go为例



### 在网关统一上报重要指标
以httproxy.go为例

https://github.com/1005281342/httproxy



## 服务监控

### grafana

![](README.assets/grafana.png)



### 导入prometheus数据

https://juejin.cn/post/6844903848230944776

![prometheus](README.assets/prometheus数据源导入.png)



## 服务告警
通过grafana配置告警时，出现访问接口异常（可能是哪个地方配置不对）【TODO】



使用alertmanager

![go协程数过多告警触发](README.assets/go协程数过多告警触发.png)

![go协程数过多告警邮件](README.assets/go协程数过多告警邮件.png)

### 参考

Grafana + prometheus搭建服务器监控系统（二）---使用alertmanger进行警报管理：https://blog.csdn.net/Alen_xiaoxin/article/details/123460647

prometheus热加载配置文件：https://blog.csdn.net/qq_21133131/article/details/117568214

Docker 容器部署的 Grafana 配置邮件告警遇到问题 https://blog.csdn.net/annghi/article/details/121374019



### QQ邮箱

授权码 https://service.mail.qq.com/cgi-bin/help?subtype=1&id=28&no=1001256



### 配置告警规则

grafana https://grafana.com/docs/grafana/latest/alerting/unified-alerting/




## 滚动升级

高度分布式且不断变化的！

1. 新版本服务启动一个节点，并主动观察其日志，主要关注错误日志，观察CPU、MEM等指标
2. 在北极星上踢掉一台旧版本节点流量，观察监控，可以查看groutine指标以及手动上报的重要接口的指标是否平稳
3. 将该节点的旧版本替换为新版本
4. 按照一定比例重复操作步骤1、2、3，直至全量，若中途发现异常（比如收到告警），（如果比较严重）立即踢掉新版本服务流量。