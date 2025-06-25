# xdi

本项目的 README 文档。

## 项目简介

请在此处填写本项目的基本介绍，包括项目的用途、主要功能和技术栈等。

## 安装方法

```bash
git clone https://github.com/winfa/xdi.git
cd xdi
# 安装依赖
# 运行项目
```

## 使用方法

本项目以依赖注入（Dependency Injection）为核心，主要包含 container、provide、invoke、resolve 等核心 API。以下为 Go 语言的基本用法示例：

### 1. 创建和使用容器（Container）

```go
import "github.com/winfa/xdi"

container := xdi.New()
```

### 2. 注册依赖（Provide）

```go
container.Provide("serviceName", func() any {
    return &MyService{}
})
```
或者注册已实例化的对象：
```go
config := &Config{Port: 8080}
container.ProvideValue("config", config)
```

### 3. 注入并运行（Invoke）

```go
err := container.Invoke(func(service *MyService, config *Config) {
    // 在这里可以使用 service 和 config
    fmt.Println(service, config)
})
if err != nil {
    panic(err)
}
```

### 4. 解析获取依赖（Resolve）

```go
svc, err := container.Resolve("serviceName")
if err != nil {
    panic(err)
}
service := svc.(*MyService)
fmt.Println(service)
```

## 贡献指南

欢迎贡献！请阅读 [CONTRIBUTING.md](CONTRIBUTING.md) 了解详情。

## 许可证

本项目采用 MIT 许可证，详见 [LICENSE](LICENSE) 文件。