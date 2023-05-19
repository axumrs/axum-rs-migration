package model

type tagBase struct {
	Name  string
	IsDel bool `db:"is_del"`
}

type PGTag struct {
	ID int32
	tagBase
}
type MySQLTag struct {
	ID uint32
	tagBase
}

func MySQLTagFrom(t *PGTag) *MySQLTag {
	return &MySQLTag{
		ID:      uint32(t.ID),
		tagBase: t.tagBase,
	}
}

func MySQLTagListFrom(ls []*PGTag) []*MySQLTag {
	ts := make([]*MySQLTag, 0, len(ls))
	for _, t := range ls {
		ts = append(ts, MySQLTagFrom(t))
	}
	return ts
}
