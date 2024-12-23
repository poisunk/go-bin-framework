# Go-Gin-Framework

一个基于 Gin 框架的 Web 应用骨架，使用 Wire 进行依赖注入。

## 项目结构

├──── cmd
│ └──── app
│   ├──── main.go # 主程序入口
│   ├──── cmd.go # 命令行工具
│   └──── wire.go # 依赖注入配置
├──── configs # 配置文件目录
│ └──── config.yaml # 应用配置文件
└──── internal # 内部代码
  ├──── base # 基础设施
  │ ├──── conf # 配置管理
  │ └──── data # 数据库连接
  ├──── controller # 控制器层
  ├──── model # 数据模型
  ├──── repository # 数据访问层
  ├──── router # 路由管理
  ├──── server # 服务器配置
  └──── service # 业务逻辑层




