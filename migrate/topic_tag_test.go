package migrate

import "testing"

func TestTopicTagRead(t *testing.T) {
	initTest(t)

	tt := &TopicTag{}

	ls, err := tt.Read()
	if err != nil {
		t.Fatal(err)
	}

	for _, i := range ls {
		t.Logf("%#v\n", i)
	}
}

func TestTopicTagMigrate(t *testing.T) {
	initTest(t)

	tt := &TopicTag{Truncate: true}

	if err := tt.Migrate(); err != nil {
		t.Fatal(err)
	}
}
