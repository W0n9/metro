# Copilot Instructions for Metro Codebase

## 项目概览

本项目用于根据北京地铁数据，查询从指定站点出发，在给定票价预算下可达的所有站点及路径。主要逻辑集中在 `main.py`，数据来源为远程 JSON 文件。

## 主要文件

- `main.py`：核心脚本，包含数据加载、票价计算、Dijkstra 路径搜索、命令行参数解析和结果输出。
- `pyproject.toml`：项目元数据和依赖声明（如 requests）。
- `uv.toml`：配置 PyPI 镜像。
- `README.md`：用法、依赖和功能说明。

## 架构与数据流

- 启动时通过命令行参数指定出发站和预算。
- 自动下载地铁数据（站点、距离、票价规则）。
- 使用 Dijkstra 算法遍历所有可达站点，计算路径、距离和票价。
- 仅输出票价等于预算的所有可达站点及路径。

## 依赖管理

- 推荐使用 [uv](https://github.com/astral-sh/uv) 管理依赖：
  ```bash
  uv pip install -r requirements.txt
  ```
- 也可用 pip3 安装：
  ```bash
  pip3 install -r requirements.txt
  ```

## 运行方式

 直接运行主脚本：
  ```bash
  python3 main.py <出发站名> <车费预算(元)> [--show-path]
  ```
  示例：
  ```bash
  python3 main.py 西直门 4
  ```
  默认不显示路径，若需显示路径可加 --show-path 选项：
  ```bash
  python3 main.py 西直门 4 --show-path
  ```

## 关键约定与模式

- 票价计算逻辑在 `calc_price`，与北京地铁实际规则保持一致。
- 路径搜索统一用 Dijkstra 算法，所有站点和路径均通过此算法生成。
- 地铁数据通过远程 JSON 获取，无需本地数据文件。
 结果输出格式为：站点名、距离、票价、所属线路。若加 --show-path 选项，则额外显示路径。

## 调试与扩展建议

- 如需调试或扩展票价规则，建议直接修改 `main.py` 的相关函数。
- 若需更换数据源，修改 `main.py` 顶部的 URL 即可。
- 所有核心逻辑均在单文件实现，便于快速理解和修改。

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
