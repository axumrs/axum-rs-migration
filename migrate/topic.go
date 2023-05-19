package migrate

import (
	"github.com/axumrs/axum-rs-migration/db"
	"github.com/axumrs/axum-rs-migration/model"
)

type Topic struct {
	Truncate bool
}

func (t *Topic) Read() (ls []*model.PGTopic, err error) {
	if err := db.PG.Select(&ls, "SELECT * FROM topic ORDER BY id ASC"); err != nil {
		return nil, err
	}
	return ls, nil
}

func (t *Topic) Write(ls []*model.MySQLTopic) error {
	tx, err := db.MY.Beginx()
	if err != nil {
		return err
	}

	if t.Truncate {
		if _, err := tx.Exec("TRUNCATE TABLE topic"); err != nil {
			return tx.Rollback()
		}
	}

	if _, err := tx.NamedExec("INSERT INTO topic (id,title, subject_id, slug, summary, author, src, hit, dateline, try_readable, is_del, cover) VALUES(:id,:title, :subject_id, :slug, :summary, :author, :src, :hit, :dateline, :try_readable, :is_del, :cover)", ls); err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}

func (t *Topic) Migrate() error {
	ls, err := t.Read()
	if err != nil {
		return err
	}
	return t.Write(model.MySQLTopicListFrom(ls))
}
