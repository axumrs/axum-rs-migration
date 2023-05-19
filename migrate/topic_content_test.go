package migrate

import "testing"

func TestTopicContentRead(t *testing.T) {
	initTest(t)

	tc := &TopicContent{}

	ls, err := tc.Read()
	if err != nil {
		t.Fatal("read failed:", err)
	}

	for _, tcc := range ls {
		t.Logf("%#v\n", tcc)
	}
}

func TestTopicContentMigrate(t *testing.T) {
	initTest(t)

	tc := &TopicContent{Truncate: true}
	if err := tc.Migrate(); err != nil {
		t.Fatal("migrate failed:", err)
	}
}
