import db.pg as pg
import db.mysql as mysql
from utils import xid
from .topic import get_slug,get_id
from .tag import get_name as get_tag_name,get_id as get_tag_id

def migration(cfg,truncate=True):
    rows = list()
    topic_id_data = dict()
    tag_id_data = dict()

    with mysql.conn(cfg, ) as c:
        with c.cursor(dictionary=True) as cur:
            cur.execute('SELECT topic_id, tag_id, is_del FROM topic_tag WHERE is_del=false ORDER BY topic_id ASC,tag_id ASC')
            r = cur.fetchall()
            for x in r:
                rows.append(x)
    
    newRows = list()
    for row in rows:
        topic_id = topic_id_data.get(row['topic_id'], None)
        tag_id = tag_id_data.get(row['tag_id'], None)
        if topic_id is None:
            topic_slug = get_slug(cfg, row['topic_id'])
            topic_id = get_id(cfg, topic_slug)
            topic_id_data[row['topic_id']] = topic_id
        
        if tag_id is None:
            tag_name = get_tag_name(cfg, row['tag_id'])
            tag_id = get_tag_id(cfg, tag_name)
            tag_id_data[row['tag_id']] = tag_id
        newRows.append((xid.new(), topic_id, tag_id))

 

    with pg.conn(cfg) as c:
        with c.cursor() as cur:
            if truncate:
                cur.execute('TRUNCATE TABLE topic_tags')
            
            for row in newRows:
                cur.execute('INSERT INTO topic_tags (id, topic_id, tag_id) VALUES(%s,%s,%s) ON CONFLICT (topic_id, tag_id) DO UPDATE SET id = EXCLUDED.id', row)
    
    return len(newRows)