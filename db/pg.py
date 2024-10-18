import psycopg
import psycopg.rows

def conn(cfg):
    """
    连接 PostgreSQL 数据库。使用方法：

    import db.pg as pg
        with conn(cfg) as c:
        with c.cursor() as cur:
            r = cur.execute('SELECT * FROM topics').fetchone()
            print(r)
    """
    dsn = cfg['PG_DSN']
    conn = psycopg.connect(dsn, row_factory=psycopg.rows.dict_row)
    return conn

def close(conn):
    if conn:
        conn.commit()
        conn.close()