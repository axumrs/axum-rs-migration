package migrate

import (
	"github.com/axumrs/axum-rs-migration/db"
	"github.com/axumrs/axum-rs-migration/model"
)

type TopicTag struct {
	Truncate bool
}

func (tt *TopicTag) Read() (ls []*model.PGTopicTag, err error) {
	if err := db.PG.Select(&ls, "SELECT * FROM topic_tag ORDER BY topic_id ASC"); err != nil {
		return nil, err
	}
	return ls, nil
}

func (tt *TopicTag) Write(ls []*model.MySQLTopicTag) error {
	tx, err := db.MY.Beginx()
	if err != nil {
		return err
	}

	if tt.Truncate {
		if _, err := tx.Exec("TRUNCATE TABLE topic_tag"); err != nil {
			return tx.Rollback()
		}
	}

	if _, err := tx.NamedExec("INSERT INTO topic_tag (topic_id, tag_id, is_del) VALUES(:topic_id, :tag_id, :is_del)", ls); err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}

func (tt *TopicTag) Migrate() error {
	ls, err := tt.Read()
	if err != nil {
		return err
	}
	return tt.Write(model.MySQLTopicTagListFrom(ls))
}
