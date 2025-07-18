# Copilot Instructions for Metro Codebase

## 项目概览

本项目用于根据北京地铁数据，查询从指定站点出发，在给定票价预算下可达的所有站点及路径。项目采用 Go 语言实现，具有高性能、跨平台编译、单文件部署等特点。

## 主要文件

- `main.go`：程序主入口
- `src/data/types.go`：数据结构定义
- `src/utils/utils.go`：工具函数
- `src/utils/dijkstra.go`：Dijkstra 算法实现
- `src/utils/utils_test.go`：单元测试
- `Makefile`：构建和开发工具链
- `go.mod`：Go 模块依赖管理

## 架构与数据流

- 启动时通过命令行参数指定出发站和预算。
- 自动下载地铁数据（站点、距离、票价规则）。
- 使用 Dijkstra 算法遍历所有可达站点，计算路径、距离和票价。
- 仅输出票价等于预算的所有可达站点及路径。

## 依赖管理

```bash
# 下载并整理依赖
go mod download
go mod tidy
```

## 运行方式

- 直接运行：
  ```bash
  go run main.go <出发站名> <车费预算(元)>
  ```
- 使用 Makefile：
  ```bash
  make run                    # 运行示例
  make build                  # 构建当前平台
  make build-all             # 构建所有平台版本
  make build-arm64           # 构建所有 ARM64 版本
  ```
- 构建后运行：
  ```bash
  make build
  ./build/metro 西直门 4
  ```

## 关键约定与模式

### 通用规则
- 票价计算逻辑在 `CalcPrice` 函数，与北京地铁实际规则保持一致。
- 路径搜索统一用 Dijkstra 算法，所有站点和路径均通过此算法生成。
- 地铁数据通过远程 JSON 获取，无需本地数据文件。
- 结果输出格式为：站点名、距离、路径、票价、所属线路。

### Go 版本特点
- 模块化设计，分为 `data`、`utils` 包。
- 支持跨平台编译，包括 AMD64 和 ARM64 架构。
- 单文件部署，无需运行时依赖。
- 包含完整的测试套件和开发工具链。

## 调试与扩展建议

### Go 版本
- 票价计算逻辑在 `src/utils/utils.go` 的 `CalcPrice` 函数。
- Dijkstra 算法实现在 `src/utils/dijkstra.go`。
- 数据结构定义在 `src/data/types.go`。
- 使用 `make test` 运行测试，`make check` 进行完整检查。
- 支持多架构编译：`make build-all` 构建所有平台版本。

## 示例输出

```
站点: 西直门
距离: 0.0km
票价: 0元
路径: 西直门
线路: 二号线, 四号线, 十三号线
---
站点: 车公庄
距离: 2.0km
票价: 3元
路径: 西直门 -> 车公庄
线路: 二号线, 六号线
```

---

如有不清楚或遗漏的部分，请反馈以便进一步完善。
