# go-task

## 远程配置

上报ELK

查询日志

## 远程日志
新增/修改 Apollo配置


## 分布式任务队列
https://github.com/hibiken/asynq

docker run --rm --name asynqmon -d -p 8098:8080 -e REDIS_ADDR=172.17.0.1:6379 hibiken/asynqmon

Web页面: http://127.0.0.1:8098/

## 服务发现与服务
心跳上报异常[TODO]

北极星

## 链路追踪 [TODO]

## 整合主流框架及主件 -- rk-boot

### 框架集成指标上报
rk-boot

### 手动上报重要指标
以manager/internal/metrics/metrics.go为例

### 在网关统一上报重要指标
以httproxy.go为例

## 服务监控

配置grafana

### 告警


## 滚动升级

高度分布式且不断变化的！