package model

import "time"

type PGTopic struct {
	ID        int64
	Title     string
	SubjectID int32 `db:"subject_id"`
	Slug      string
	Summary   string
	Src       string
	Author    string
	Hit       int32
	Dateline  int32
	IsDel     bool `db:"is_del"`
}

type MySQLTopic struct {
	ID          uint64
	Title       string
	SubjectID   uint32 `db:"subject_id"`
	Slug        string
	Summary     string
	Src         string
	Author      string
	Hit         uint64
	Dateline    time.Time
	IsDel       bool `db:"is_del"`
	TryReadable bool `db:"try_readable"`
	Cover       string
}

func MySQLTopicFrom(pt *PGTopic) *MySQLTopic {
	return &MySQLTopic{
		ID:          uint64(pt.ID),
		Title:       pt.Title,
		SubjectID:   uint32(pt.SubjectID),
		Slug:        pt.Slug,
		Summary:     pt.Summary,
		Src:         pt.Src,
		Author:      pt.Author,
		Hit:         uint64(pt.Hit),
		Dateline:    time.Unix(int64(pt.Dateline), 0),
		IsDel:       pt.IsDel,
		TryReadable: false,
		Cover:       "",
	}
}

func MySQLTopicListFrom(ls []*PGTopic) []*MySQLTopic {
	ms := make([]*MySQLTopic, 0, len(ls))

	for _, t := range ls {
		ms = append(ms, MySQLTopicFrom(t))
	}

	return ms
}
