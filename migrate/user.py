import db.pg as pg
import db.mysql as mysql
from utils import xid


def migration(cfg, truncate=True):
    users = list()
    # userStatus = ('Pending', 'Actived', 'Freezed')
    userTypes = ('Normal', 'Subscriber')

    with mysql.conn(cfg) as c:
        with c.cursor(dictionary=True) as cur:
            cur.execute('SELECT id, email, nickname, password, status, dateline, types, sub_exp, points, allow_device_num, jwt_exp, is_del FROM `user` WHERE is_del=false ORDER BY id ASC')
            r = cur.fetchall()
            for x in r:
                users.append(x)
    
    users = tuple(map(lambda x: (xid.new(), x['email'], x['nickname'], x['password'], 'Pending', x['dateline'], userTypes[x['types']], x['sub_exp'], x['points'], 1, 20), users))
 
    with pg.conn(cfg) as c:
        with c.cursor() as cur:
            if truncate:
                cur.execute('TRUNCATE TABLE users')
            
            for user in users:
                cur.execute('INSERT INTO users (id, email, nickname, "password", status, dateline, kind, sub_exp, points, allow_device_num, session_exp) VALUES( %s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)', user)
    
    return len(users)