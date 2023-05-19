package model

type PGSubject struct {
	ID      int32
	Name    string
	Slug    string
	Summary string
	IsDel   bool `db:"is_del"`
}

type MySQLSubjectStatus uint8

const (
	MySQLSubjectStatusWriting MySQLSubjectStatus = iota
	MySQLSubjectStatusFinished
)

type MySQLSubject struct {
	ID      uint32
	Name    string
	Slug    string
	Summary string
	IsDel   bool `db:"is_del"`
	Cover   string
	Status  MySQLSubjectStatus
	Price   uint32
}

func MySQLSubjectFrom(ps *PGSubject) *MySQLSubject {
	return &MySQLSubject{
		ID:      uint32(ps.ID),
		Name:    ps.Name,
		Slug:    ps.Slug,
		Summary: ps.Summary,
		IsDel:   ps.IsDel,
		Cover:   "",
		Status:  MySQLSubjectStatusFinished,
		Price:   0,
	}
}

func MySQLSubjectListFrom(list []*PGSubject) []*MySQLSubject {
	ss := make([]*MySQLSubject, 0, len(list))

	for _, s := range list {
		ss = append(ss, MySQLSubjectFrom(s))
	}
	return ss
}
