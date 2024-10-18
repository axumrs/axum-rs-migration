import db.pg as pg
import db.mysql as mysql
from utils import xid

def migration(cfg,truncate=True):
    rows = list()

    with mysql.conn(cfg, ) as c:
        with c.cursor(dictionary=True) as cur:
            cur.execute('SELECT id, name, is_del FROM tag ORDER BY id ASC')
            r = cur.fetchall()
            for x in r:
                rows.append(x)
    
    rows = tuple(map(lambda x: (xid.new(), x['name'], x['is_del']==1), rows))

    with pg.conn(cfg) as c:
        with c.cursor() as cur:
            if truncate:
                cur.execute('TRUNCATE TABLE tags')
            
            for row in rows:
                cur.execute('INSERT INTO tags (id, "name", is_del) VALUES(%s,%s,%s)', row)
    
    return len(rows)

def get_name(cfg, mysql_id):
    with mysql.conn(cfg, ) as c:
        with c.cursor(dictionary=True) as cur:
            cur.execute('SELECT name FROM tag WHERE id = %s', (mysql_id, ))
            r = cur.fetchone()
            if r:
                return r['name']
            else:
                return None

def get_id(cfg, name):
    with pg.conn(cfg) as c:
        with c.cursor() as cur:
            cur.execute('SELECT id FROM tags WHERE "name" = %s', (name, ))
            r = cur.fetchone()
            if r:
                return r['id']
            else:
                return None