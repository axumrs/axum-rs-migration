import db.pg as pg
import db.mysql as mysql
from utils import xid,sha256
from .subject import get_id
from bs4 import BeautifulSoup


def migration(cfg,truncate=True):
    rows = list()

    with mysql.conn(cfg, ) as c:
        with c.cursor(dictionary=True) as cur:
            cur.execute('''SELECT
	t.id, title, t.slug, t.summary, author, src, hit, dateline, try_readable, t.is_del, t.cover, t.pin
	,md,html
	,s.slug AS subject_slug
FROM 
	topic AS t
INNER JOIN
	topic_content AS tc 
ON	tc.topic_id  = t.id
INNER JOIN 
	subject AS s
ON t.subject_id  = s.id

ORDER BY t.id ASC ''')
            r = cur.fetchall()
            for x in r:
                rows.append(x)
    
   

   
    newRows = list()
    htmls = list()
    for row in rows:
         # 获取专题ID
        subject_id = get_id(cfg, row['subject_slug'])
        topic_id = xid.new()
        newRows.append((topic_id, row['title'], row['slug'], row['summary'], row['author'], row['src'], row['hit'], row['dateline'], row['try_readable']==1, row['is_del']==1,row['cover'],row['md'],row['pin'],subject_id))
        htmls.append((topic_id,row['html']))


    with pg.conn(cfg) as c:
        with c.cursor() as cur:
            if truncate:
                cur.execute('TRUNCATE TABLE topics')
            
            for row in newRows:
                cur.execute('INSERT INTO topics (id, title, slug, summary, author, src, hit, dateline, try_readable, is_del, cover, md, pin, subject_id) VALUES (%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)', row)
    
    return (len(newRows), htmls)


def migration_content(cfg,htmls,truncate=True):
    datas = list()

    for (topic_id,html) in htmls:
            soup = BeautifulSoup(html, 'html5lib')
            secs = soup.find_all(cfg['SECTION_ELEMENTS'].split(','))
            for idx,s in enumerate(secs):
                hashed_str = sha256.hash_str(s,cfg['HASH_SECRET_KEY'])
                sec_id = xid.new()
                s['data-section'] = sec_id
                data = (sec_id, topic_id, idx, hashed_str, str(s))
                datas.append(data)
                
    with pg.conn(cfg) as c:
        with c.cursor() as cur:
            if truncate:
                cur.execute('TRUNCATE TABLE topic_sections')
            
            for data in datas:
                cur.execute('INSERT INTO topic_sections (id, topic_id, sort, hash, "content") VALUES (%s,%s,%s,%s,%s)', data)

    return len(datas)

def get_slug(cfg,mysql_id):
    with mysql.conn(cfg, ) as c:
        with c.cursor(dictionary=True) as cur:
            cur.execute('SELECT slug FROM topic WHERE id = %s', (mysql_id, ))
            r = cur.fetchone()
            if r:
                return r['slug']
            else:
                return None

def get_id(cfg, slug):
    with pg.conn(cfg) as c:
        with c.cursor() as cur:
            cur.execute('SELECT id FROM topics WHERE slug = %s', (slug, ))
            r = cur.fetchone()
            if r:
                return r['id']
            else:
                return None