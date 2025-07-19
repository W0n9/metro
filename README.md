# 地铁票价可达站点查询

本项目用于根据北京地铁数据，查询从指定站点出发，在给定票价预算下可达的所有站点及路径。

## 依赖

- Python 3.13+
- requests

## 使用 uv 管理依赖

推荐使用 [uv](https://github.com/astral-sh/uv) 进行依赖管理和安装：

```bash
uv pip install -r requirements.txt
```

或直接安装：

```bash
uv pip install requests
```

## 安装依赖（pip 方式）

```bash
pip3 install -r requirements.txt
```

或直接安装：

```bash
pip3 install requests
```

## 用法

```bash
python3 main.py <出发站名> <车费预算(元)> [--show-path]
```

例如：

```bash
python3 main.py 西直门 4
```
默认不显示路径，若需显示路径可加 --show-path 选项：

```bash
python3 main.py 西直门 4 --show-path
```

## 功能说明

- 自动加载地铁数据（含站点、距离、票价规则）。
- 使用 Dijkstra 算法计算所有可达站点及路径。
- 默认输出所有票价等于预算的可达站点、距离和所属线路。
- 若加 --show-path 选项，则额外显示路径。

## 输出示例

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

## 数据来源

地铁数据来自：
https://static.qinxr.cn/Hyacinth/farecalc.json

## 许可证

MIT License
