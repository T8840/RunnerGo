

## 基于go语言的一体化接口性能测试平台

RunnerGo致力于打造成一款全栈式测试平台，采用了较为宽松的Apache-2.0 license开源协议，方便志同道合的朋友一起为开源贡献力量，目前实现了接口测试、场景自动化测试、性能测试等测试能力。随着不断的迭代，我们将会推出更多的测试功能。我们的目的是为研发赋能，让测试更简单。

## 工具特性：
- go语言运行：基于go语言开发，运行速度快、更节省资源
- 智能调度算法：自研的调度算法，合理利用服务器资源，降低资源消耗
- 实时生成测试报告：运行任务后，可实时查看执行结果，快速诊断服务病症
- 丰富的报告图表： 全方位展示各个指标运行曲线图
- 实时修改： 可根据压测模式实时修改并发数、持续时长等
- 实时日志： 可在压测过程中开启日志模式，查看请求响应信息
- 可编辑报告：可在任务运行结束后，针对测试结果进行测试分析，实时编写报告
- Flow场景流：可视化的业务流，通过连线就可快速搭建起来自己的业务流，还可直接调试运行场景，电流般的业务流转
- 多种压测模式：支持并发模式、阶梯模式、错误率模式、响应时间模式、每秒应答数模式、轮次模式等多种压测模式，支持根据机器自定义分布配置，满足所有业务需求
- 自持接口自动化，采用用例集概念，生成丰富的自动化报告
- Mock服务：支持自定义请求校验与响应期望
- 企业管理后台：支持多团队管理，通过权限设置来管理员工，保护公司数据安全和流量资源

### 首页展示
![interface](https://apipost.oss-cn-beijing.aliyuncs.com/kunpeng/images/home.jpg)

### 性能测试报告
![report](https://apipost.oss-cn-beijing.aliyuncs.com/kunpeng/images/stress_report.jpg)

### 性能测试报告对比
![报告对比.jpg](https://apipost.oss-cn-beijing.aliyuncs.com/kunpeng/images/contrast.jpg)

### 自动化测试报告
![report](https://apipost.oss-cn-beijing.aliyuncs.com/kunpeng/images/auto_report.jpg)

### Mock服务

![report](https://apipost.oss-cn-beijing.aliyuncs.com/kunpeng/images/mock.png)

### 企业管理后台

![report](https://apipost.oss-cn-beijing.aliyuncs.com/kunpeng/images/2.3.0-9.png)


## 快速开始

开源版安装教程请见 deploy.md

默认超管账号**runnergo**  密码**runnergo**

里面有非常详细的图文教程，如需远程指导，也可划到当前页面最下方添加我们的微信，我们会为您提供安装帮助。


## 技术栈
- 后端: GoLang
- 前端: React.js
- 中间件: MySQL, MongoDB, Kafka, ZooKeeper, Redis
- 基础设施: Docker
- 测试引擎: GoLang

## 技术架构
![struct](https://apipost.oss-cn-beijing.aliyuncs.com/kunpeng/images/struct.png)

## 业务流转图
![flow](https://apipost.oss-cn-beijing.aliyuncs.com/kunpeng/images/flow.png)

## 相关源码清单：

### runnerGo-management-websocket-open

### runnergo-management-open

### runnergo-collector-open

### runnergo-engine-open

### file-server

### runnergo-fe-open

### runnergo-fe-admin-open

### runnergo-permission-open