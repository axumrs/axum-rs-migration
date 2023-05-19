package migrate

import (
	"github.com/axumrs/axum-rs-migration/db"
	"github.com/axumrs/axum-rs-migration/model"
)

type Tag struct {
	Truncate bool
}

func (t *Tag) Read() (ls []*model.PGTag, err error) {
	if err := db.PG.Select(&ls, "SELECT * FROM tag ORDER BY id ASC"); err != nil {
		return nil, err
	}
	return ls, nil
}

func (t *Tag) Write(ls []*model.MySQLTag) error {
	tx, err := db.MY.Beginx()
	if err != nil {
		return err
	}

	if t.Truncate {
		if _, err := tx.Exec("TRUNCATE TABLE tag"); err != nil {
			return tx.Rollback()
		}
	}

	if _, err := tx.NamedExec("INSERT INTO tag (id,name, is_del) VALUES(:id,:name, :is_del)", ls); err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}

func (t *Tag) Migrate() error {
	ls, err := t.Read()
	if err != nil {
		return err
	}
	return t.Write(model.MySQLTagListFrom(ls))
}
