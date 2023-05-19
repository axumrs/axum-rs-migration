package migrate

import "testing"

func TestTagRead(t *testing.T) {
	initTest(t)

	tg := &Tag{}
	ls, err := tg.Read()
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range ls {
		t.Logf("%#v\n", tt)
	}
}

func TestTagMigrate(t *testing.T) {
	initTest(t)

	tg := &Tag{Truncate: true}
	if err := tg.Migrate(); err != nil {
		t.Fatal(err)
	}
}
