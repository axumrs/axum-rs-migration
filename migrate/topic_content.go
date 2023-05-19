package migrate

import (
	"github.com/axumrs/axum-rs-migration/db"
	"github.com/axumrs/axum-rs-migration/model"
)

type TopicContent struct {
	Truncate bool
}

func (tc *TopicContent) Read() (ls []*model.PGTopicContent, err error) {
	if err := db.PG.Select(&ls, "SELECT * FROM topic_content ORDER BY topic_id ASC"); err != nil {
		return nil, err
	}
	return ls, nil
}

func (tc *TopicContent) Write(ls []*model.MySQLTopicContent) error {
	tx, err := db.MY.Beginx()
	if err != nil {
		return err
	}

	if tc.Truncate {
		if _, err := tx.Exec("TRUNCATE TABLE topic_content"); err != nil {
			return tx.Rollback()
		}
	}

	if _, err := tx.NamedExec("INSERT INTO topic_content (topic_id, md, html) VALUES(:topic_id, :md, :html)", ls); err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}

func (tc *TopicContent) Migrate() error {
	ls, err := tc.Read()
	if err != nil {
		return err
	}
	return tc.Write(model.MySQLTopicContentListFrom(ls))
}
