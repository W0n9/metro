import heapq
import requests

# 数据源
DATA_URL = "https://static.qinxr.cn/Hyacinth/farecalc.json"

# 票价计算规则


def calc_price(dis: int, free_dis: int) -> int:
    dis -= free_dis
    if dis <= -free_dis:
        return 3
    elif dis <= 6000:
        return 3
    elif dis <= 12000:
        return 4
    elif dis <= 22000:
        return 5
    elif dis <= 32000:
        return 6
    else:
        return int((dis - 32000) / 20000 + 7)

# 加载地铁数据


def load_metro_data() -> dict:
    resp = requests.get(DATA_URL, timeout=15)
    resp.raise_for_status()
    return resp.json()

# Dijkstra算法，返回所有可达站点及路径、距离、票价


def dijkstra_all(sta_dict: dict, start_id: str, free_dis: int) -> dict[str, dict]:
    visited = set()
    heap = [(0, start_id, [start_id])]
    result = {}
    while heap:
        dis, curr, path = heapq.heappop(heap)
        if curr in visited:
            continue
        visited.add(curr)
        price = calc_price(dis, free_dis)
        result[curr] = {
            'distance': dis,
            'price': price,
            'path': path.copy()
        }
        for neighbor, edges in sta_dict[curr]['edges'].items():
            for edge in edges:
                next_dis = dis + edge['dis']
                if neighbor not in visited:
                    heapq.heappush(
                        heap, (next_dis, neighbor, path + [neighbor]))
    return result

# 站名转ID


def name_to_id(sta_to_id: dict, name: str) -> str:
    return str(sta_to_id.get(name, -1))

# ID转站名


def id_to_name(sta_dict: dict, id_: str) -> str:
    return sta_dict[id_]['name'] if id_ in sta_dict else id_

# 获取站点所属线路名


def get_station_lines(sta_id: str, line_detail: dict) -> list:
    lines = []
    for line_id, detail in line_detail.items():
        if int(sta_id) in detail.get('staList', []):
            lines.append(detail['name'])
    return lines

# 主入口


def main():
    import sys
    if len(sys.argv) < 3 or len(sys.argv) > 4:
        print("用法: python3 main.py <出发站名> <车费预算(元)> [--show-path]")
        print("[--show-path] 可选，输出时显示路径")
        return
    start_name = sys.argv[1].strip()
    try:
        budget = int(sys.argv[2])
    except ValueError:
        print("预算必须为整数！")
        return
    show_path = False
    if len(sys.argv) == 4 and sys.argv[3] == "--show-path":
        show_path = True

    print("正在加载地铁数据...")
    data = load_metro_data()
    sta_dict = data['staDict']
    sta_to_id = data['staToId']
    free_dis = int(data['freeDis'])
    line_detail = data['lineDetail']

    start_id = name_to_id(sta_to_id, start_name)
    if start_id == "-1":
        print("站名不存在！")
        return

    print(f"以 {start_name} 为起点，预算 {budget} 元，可达站点如下：")
    result = dijkstra_all(sta_dict, start_id, free_dis)
    # 按线路分组，换乘站分别并入各所属线路
    line_groups = {}
    for sid, info in result.items():
        if info['price'] == budget and sid != start_id:
            lines = get_station_lines(sid, line_detail)
            if not lines:
                # 没有线路信息，归入“未知”
                if '未知' not in line_groups:
                    line_groups['未知'] = []
                line_groups['未知'].append((sid, info))
            else:
                for line in lines:
                    if line not in line_groups:
                        line_groups[line] = []
                    line_groups[line].append((sid, info))
    count = 0
    for line_key, stations in line_groups.items():
        print(line_key)
        for sid, info in stations:
            path_names = [id_to_name(sta_dict, pid) for pid in info['path']]
            if show_path:
                print(
                    f"{id_to_name(sta_dict, sid)} | 距离: {info['distance']}m | 路径: {' -> '.join(path_names)} | 票价: {info['price']}元")
            else:
                print(
                    f"{id_to_name(sta_dict, sid)} | 距离: {info['distance']}m | 票价: {info['price']}元")
            count += 1
        print("-----")
    print(f"共 {count} 个站点可达。")


if __name__ == "__main__":
    main()
