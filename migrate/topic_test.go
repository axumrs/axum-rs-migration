package migrate

import (
	"testing"

	"github.com/axumrs/axum-rs-migration/model"
)

func TestTopicRead(t *testing.T) {
	initTest(t)

	tp := &Topic{}
	ls, err := tp.Read()
	if err != nil {
		t.Fatal("read failed:", err)
	}
	for _, tt := range ls {
		t.Logf("%#v\n", tt)
	}
}

func TestTopicWrite(t *testing.T) {
	initTest(t)

	tp := &Topic{Truncate: true}

	ls, err := tp.Read()
	if err != nil {
		t.Fatal("read failed:", err)
	}

	ms := model.MySQLTopicListFrom(ls)

	if err := tp.Write(ms); err != nil {
		t.Fatal("write failed:", err)
	}
}

func TestTopicMigrate(t *testing.T) {
	initTest(t)

	tp := &Topic{Truncate: true}
	if err := tp.Migrate(); err != nil {
		t.Fatal("migrate failed:", err)
	}
}
