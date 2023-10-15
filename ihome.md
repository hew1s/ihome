# ihome租房项目实战

- **go-micro+gin+jquery**

## 第一天

搭建基本项目结构框架

MVC (一种代码项目规范)

- M model 用户存放数据文件模型
- V view 前端显示视图
- C controller 控制器 主要为逻辑业务处理函数

ihome项目结构：

ihome

service-- 服务端

web-- 客户端（gin框架实现）

- conf-- 配置文件

- controller-- 控制

- dao--数据库操作（mysql+redis）
	- mysql-- mysql数据库文件
	- redis--redis数据库文件
- logger--日志层（可选）
- logic-- 业务处理层
- model--数据文件模型
- utils--工具文件
- router--定义路由层
- settings-- 配置文件处理层