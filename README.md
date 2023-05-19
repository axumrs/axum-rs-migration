# axum.rs 数据迁移工具

本工具用于将 AXUM.RS 旧版的 PostgreSQL 数据迁移到新版的 MySQL。

## 环境变量

请在 `.env` 设置两个环境变量：

- `PG_DSN`：旧版的 PostgreSQL 的连接字符串

- `MYSQL_DSN`：新版的 MySQL 的连接字符串

例如：

```
PG_DSN='host=127.0.0.1 port=25432 user=axum_rs password=axum_rs dbname=axum_rs sslmode=disable'
MYSQL_DSN='root:root@tcp(127.0.0.1:23306)/axum_rs?charset=utf8mb4&collation=utf8mb4_unicode_ci&loc=PRC&parseTime=true'
```

## 运行

```bash
go run main.go
```

## FAQ

- 为什么不使用 rust 编写本工具：因为累 😮‍💨

- 为什么不提供预编译的二进制文件：（同上）

- 没有 Go 环境怎么办：装一个
