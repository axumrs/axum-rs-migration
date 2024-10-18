import mysql.connector

def conn(cfg):
    """
    连接 MySQL 数据库。使用方法：
    
    import db.mysql as mysql
      with mysql.conn(cfg) as c:
        with c.cursor(dictionary =True) as cur:
             cur.execute('SELECT * FROM topic')
             r = cur.fetchall()
             for x in r:
                 print(x)
    """
    dsn = {'host': cfg['MYSQL_HOST'], 'user': cfg['MYSQL_USER'], 'password': cfg['MYSQL_PASSWORD'], 'database': cfg['MYSQL_DATABASE'], 'port':cfg['MYSQL_PORT'],'charset':'utf8mb4', 'collation':'utf8mb4_unicode_ci', 'time_zone':'PRC'}
    c = mysql.connector.connect(**dsn)
    return c