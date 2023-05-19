package migrate

import (
	"github.com/axumrs/axum-rs-migration/db"
	"github.com/axumrs/axum-rs-migration/model"
)

type Subject struct {
	Truncate bool
}

func (s *Subject) Read() (list []*model.PGSubject, err error) {
	if err := db.PG.Select(&list, "SELECT * FROM subject ORDER BY id ASC"); err != nil {
		return nil, err
	}
	return list, nil
}

func (s *Subject) Write(list []*model.MySQLSubject) error {
	tx, err := db.MY.Beginx()
	if err != nil {
		return err
	}

	if s.Truncate {
		if _, err := tx.Exec("TRUNCATE TABLE subject"); err != nil {
			return tx.Rollback()
		}
	}

	if _, err := tx.NamedExec("INSERT INTO subject (id,name, slug, summary, is_del, cover, status, price) VALUES(:id, :name, :slug, :summary, :is_del, :cover, :status, :price)", list); err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}

func (s *Subject) Migrate() error {
	list, err := s.Read()
	if err != nil {
		return err
	}

	if err := s.Write(model.MySQLSubjectListFrom(list)); err != nil {
		return err
	}
	return nil
}
