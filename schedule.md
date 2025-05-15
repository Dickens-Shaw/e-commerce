# 开发计划

> 本项目采用 monorepo 结构，所有端和服务统一放在 packages 目录下，shared 目录用于共享类型和工具，依赖管理推荐 pnpm/yarn workspaces/turbo。

## ✅ 开发顺序建议（Go 后端优先）

### 第一阶段：后端服务优先（Go）

**为什么先做后端？**

- 数据模型需要先定义（用户、商品、订单等）
- 接口确定后前端才好联调
- 避免前端先做了 mock 接口，后端设计不一致导致返工
- BFF（Nest）和 Next.js 都依赖后端接口

---

## 🏗️ 开发步骤路线图（按阶段推进）

### 🔹 阶段一：Go 后端服务搭建

#### Go 服务目录结构示例

> packages/go-service：Go 微服务（如 user/order/product 等）

```text
/packages/go-service
  /internal
    /user
    /product
    /order
    /auth
  /config
  /pkg
  /cmd
  /api
  /model
  /service
  /repository
  /main.go
  /go.mod
  ...
```

#### 📁 技术栈（Go 后端）

- 框架：Gin 或 Fiber
- ORM：gorm
- 数据库：MySQL
- Redis：缓存 & 登录态支持
- 工具：Swagger、Viper、Zap、Wire、Docker

> 开发需遵循 cursorrules 文件中的统一代码风格、接口规范、测试和文档要求。

#### 📌 第一步：初始化项目结构（推荐 Monorepo 分模块）

```text
/packages/go-service
  /internal
    /user
    /product
    /order
    /auth
  /config
  /pkg
  /main.go
  /go.mod
```

#### 📌 第二步：基础服务搭建

- 路由注册：使用 Gin 路由组
- 配置加载：用 Viper 管理 YAML 配置
- 日志系统：用 zap 打通全局日志
- 中间件：日志、鉴权、限流、跨域
- Swagger 文档：基于 swaggo/gin-swagger

#### 📌 第三步：模块逐步实现

- 用户模块（注册登录、JWT 签发、鉴权中间件）
- 商品模块（分类/SPU/SKU 查询、库存等）
- 订单模块（下单、支付模拟、状态流转）

---

### 🔹 阶段二：Nest.js 中台（BFF 层）

#### BFF 目录结构示例

```text
/packages/bff
  /src
    /modules
    /common
    /config
    /dto
    /guards
    /main.ts
  /test
  /docs
  ...
```

#### 📁 技术栈（BFF）

- 框架：Nest.js + TypeScript
- 接口类型：REST / GraphQL 二选一（推荐 REST 初期）
- 安全：JWT + Passport.js
- 配置管理：@nestjs/config

> 开发需遵循 cursorrules 文件中的统一代码风格、接口规范、测试和文档要求。

#### 主要职责

- 聚合后端接口
- 权限控制
- 数据转换（DTO）
- 为前端做接口适配优化
- 接口安全
- 统一返回格式
- 异常处理
- API 文档自动化

---

### 🔹 阶段三：前端开发

- /packages/web：C端商城网站（Next.js 用户前台）
- /packages/admin：后台管理系统（Next.js + Ant Design/Chakra UI）

#### C端商城网站（/web）

- 首页、商品列表、商品详情
- 购物车、下单结算、订单页
- 个人中心、收货地址、售后入口
- 营销活动页（秒杀、拼团、优惠券等）
- SSR/SSG 优化 SEO，响应式设计，图片懒加载、骨架屏、预加载

#### 后台管理系统（/admin）

- 商品、订单、用户、活动管理
- 权限管理、数据大盘、内容管理
- 管理端专属功能（批量操作、权限分级、日志审计等）
- Ant Design/Chakra UI 管理端组件，RBAC 权限控制，数据可视化

## 🚀 推荐开发节奏（按周推进）

| 周数   | 内容                                 |
| ------ | ------------------------------------ |
| 第1周  | 搭建 Go 服务基础骨架（用户模块）     |
| 第2周  | 商品模块 & MySQL 表设计              |
| 第3周  | 订单模块 & Redis 接入                |
| 第4周  | 支付模拟 & BFF 层开发                |
| 第5周  | Next.js C端商城页面搭建（首页、商品、订单、购物车、个人中心等） |
| 第6周  | Next.js 后台管理系统页面搭建（商品、订单、用户、活动管理等）   |
| 第7周  | 联调 + 登录态接入 + 测试             |

> 各阶段开发需严格遵循 cursorrules 规范，确保代码质量和协作一致性。
