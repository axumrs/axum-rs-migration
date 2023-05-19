package migrate

import (
	"testing"

	"github.com/axumrs/axum-rs-migration/db"
)

const (
	pgDSN    = "host=127.0.0.1 port=25432 user=axum_rs password=axum_rs dbname=axum_rs sslmode=disable"
	mysqlDSN = "root:root@tcp(127.0.0.1:23306)/axum_rs?charset=utf8mb4&collation=utf8mb4_unicode_ci&loc=PRC&parseTime=true"
)

func initTest(t *testing.T) {
	if err := db.InitPG(pgDSN); err != nil {
		t.Fatal("connect pg failed:", err)
	}
	if err := db.InitMySQL(mysqlDSN); err != nil {
		t.Fatal("connect mysql failed:", err)
	}
}
