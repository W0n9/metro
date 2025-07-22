import duckdb
import polars as pl
from main import load_metro_data, dijkstra_all

DATA_URL = "https://static.qinxr.cn/Hyacinth/farecalc.json"


def main():
    data = load_metro_data()
    sta_dict = data['staDict']
    free_dis = int(data['freeDis'])
    all_sta_ids = list(sta_dict.keys())
    all_sta_names = {sid: sta_dict[sid]['name'] for sid in all_sta_ids}

    # 连接 DuckDB
    con = duckdb.connect('metro_fare.duckdb')

    for start_id in all_sta_ids:
        start_name = all_sta_names[start_id]
        result = dijkstra_all(sta_dict, start_id, free_dis)
        records = []
        for end_id, info in result.items():
            end_name = all_sta_names[end_id]
            records.append({
                'to_station': end_name,
                'distance': info['distance'],
                'price': info['price'],
                'path': '->'.join([all_sta_names[pid] for pid in info['path']])
            })
        df = pl.DataFrame(records)
        table_name = f"from_{start_name.replace(' ', '_')}"
        # DuckDB表名需转义
        con.execute(
            f"CREATE OR REPLACE TABLE '{table_name}' AS SELECT * FROM df")
    con.close()


if __name__ == "__main__":
    main()
