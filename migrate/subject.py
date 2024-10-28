import db.pg as pg
import db.mysql as mysql
from utils import xid

def migration(cfg,truncate=True):
    rows = list()
    subjectState = ('Writing', 'Finished')

    with mysql.conn(cfg, ) as c:
        with c.cursor(dictionary=True) as cur:
            cur.execute('SELECT id, name, slug, summary, is_del, cover, status, price, pin FROM subject ORDER BY id ASC')
            r = cur.fetchall()
            for x in r:
                rows.append(x)
    
    rows = tuple(map(lambda x: (xid.new(), x['name'], x['slug'], x['summary'], x['cover'].replace('https://file.axum.rs', cfg['FILE_URL']),subjectState[x['status']], x['price']//100, x['pin'],x['is_del']==1), rows))

    with pg.conn(cfg) as c:
        with c.cursor() as cur:
            if truncate:
                cur.execute('TRUNCATE TABLE subjects')
            
            for row in rows:
                cur.execute('INSERT INTO subjects (id, "name", slug, summary, cover, status, price, pin, is_del) VALUES( %s,%s,%s,%s,%s,%s,%s,%s,%s)', row)
    
    return len(rows)

def get_id(cfg, slug):
    with pg.conn(cfg) as c:
        with c.cursor() as cur:
            r = cur.execute('SELECT id FROM subjects WHERE slug = %s', (slug, )).fetchone()
            if r:
                return r['id']
            else:
                return None