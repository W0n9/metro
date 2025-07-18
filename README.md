# 北京地铁查询工具

基于北京地铁数据，查询从指定站点出发，在给定票价预算下可达的所有站点及路径。

本项目采用 Go 语言实现，具有高性能、跨平台编译、单文件部署等特点。

## 功能特点

- 根据出发站和预算查询可达站点
- 使用 Dijkstra 算法计算最短路径
- 显示距离、票价、路径和所属线路信息
- 数据来源于远程 JSON 接口，实时更新
- 高性能 Go 语言实现
- 跨平台编译支持
- 单文件部署，无需额外依赖

## 项目结构

```
metro/
├── main.go                 # 程序主入口
├── go.mod                  # Go 模块定义
├── Makefile                # 构建和开发工具
├── src/                    # 源码目录
│   ├── data/
│   │   └── types.go        # 数据结构定义
│   └── utils/
│       ├── utils.go        # 工具函数
│       ├── utils_test.go   # 单元测试
│       └── dijkstra.go     # Dijkstra 算法实现
├── build/                  # 构建输出目录
└── README.md
```

## 安装依赖

```bash
# 下载并整理依赖
go mod download
go mod tidy
```

## 使用方法

```bash
# 直接运行
go run main.go <出发站名> <车费预算(元)>

# 示例
go run main.go 西直门 4

# 使用 Makefile 运行示例
make run
```

### 输出示例

```
正在加载地铁数据...
以 西直门 为起点，预算 4 元，可达站点如下：
积水潭 | 距离: 3000m | 票价: 4元 | 路径: 西直门 -> 积水潭 | 线路: 二号线
建国门 | 距离: 8500m | 票价: 4元 | 路径: 西直门 -> 复兴门 -> 建国门 | 线路: 一号线, 二号线
...
共 15 个站点可达。
```

## 构建可执行文件

```bash
# 构建当前平台
make build                    # 或 go build -o build/metro main.go

# 使用 Makefile 构建
make build-all               # 构建所有平台版本
make build-common            # 构建常用平台 (AMD64)
make build-arm64             # 构建所有 ARM64 版本

# 单独构建特定平台
make build-linux             # Linux AMD64
make build-linux-arm64       # Linux ARM64
make build-windows           # Windows AMD64  
make build-windows-arm64     # Windows ARM64
make build-macos             # macOS AMD64 (Intel Mac)
make build-macos-arm64       # macOS ARM64 (Apple Silicon)

# 构建特定操作系统的所有架构
make build-linux-all         # Linux 所有架构
make build-windows-all       # Windows 所有架构
make build-macos-all         # macOS 所有架构
```

## 开发和测试

```bash
# 运行测试
make test                    # 或 go test -v ./...

# 代码格式化
make fmt                     # 或 go fmt ./...

# 代码静态检查
make vet                     # 或 go vet ./...

# 完整检查 (格式化 + 检查 + 测试)
make check

# 清理构建文件
make clean

# 显示帮助信息
make help

# 显示项目信息
make info
```

## 核心算法

### 票价计算规则

根据北京地铁实际票价规则：
- 6公里以内：3元
- 6-12公里：4元
- 12-22公里：5元
- 22-32公里：6元
- 32公里以上：每增加20公里加1元

### 路径搜索

使用 Dijkstra 算法进行最短路径搜索，确保找到从起点到所有可达站点的最优路径。

## 技术特点

- **高性能**：Go 语言原生并发支持，处理速度快
- **内存安全**：类型安全，避免常见的内存错误
- **跨平台**：支持多种操作系统和架构
  - Linux (AMD64/ARM64)
  - Windows (AMD64/ARM64) 
  - macOS (Intel/Apple Silicon)
- **单文件部署**：编译后为单个可执行文件，部署简单
- **零依赖**：无需额外安装运行时或依赖库
- **完整测试**：包含单元测试，确保代码质量

## 依赖说明

项目使用 Go 标准库实现，主要依赖：
- `net/http`: HTTP 客户端，用于获取地铁数据
- `encoding/json`: JSON 解析
- `container/heap`: 堆数据结构，用于 Dijkstra 算法
- `strconv`: 字符串转换
- `fmt`: 格式化输出

## 开发说明

本项目采用标准的 Go 项目结构：
- `main.go`: 程序入口点
- `src/data/`: 数据结构定义
- `src/utils/`: 工具函数和核心算法
- `go.mod`: 依赖管理

代码遵循 Go 的代码规范，使用 `gofmt` 格式化，支持 `go vet` 静态检查。
