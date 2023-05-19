package migrate

import (
	"testing"

	"github.com/axumrs/axum-rs-migration/model"
)

func TestSubjectRead(t *testing.T) {
	initTest(t)

	s := &Subject{}

	list, err := s.Read()
	if err != nil {
		t.Fatal("read failed:", err)
	}

	for _, s := range list {
		t.Logf("%#v\n", s)
	}
}

func TestSubjectWrite(t *testing.T) {
	initTest(t)
	s := &Subject{Truncate: true}

	list, err := s.Read()
	if err != nil {
		t.Fatal("read failed:", err)
	}

	mList := model.MySQLSubjectListFrom(list)

	if err := s.Write(mList); err != nil {
		t.Fatal("write failed:", err)
	}
}

func TestSubjectMigrate(t *testing.T) {
	initTest(t)

	s := &Subject{Truncate: true}

	if err := s.Migrate(); err != nil {
		t.Fatal("migrate failed:", err)
	}
}
