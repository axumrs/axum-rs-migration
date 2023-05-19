package model

type PGTopicTag struct {
	TopicID int64 `db:"topic_id"`
	TagID   int32 `db:"tag_id"`
	IsDel   bool  `db:"is_del"`
}
type MySQLTopicTag struct {
	TopicID uint64 `db:"topic_id"`
	TagID   uint32 `db:"tag_id"`
	IsDel   bool   `db:"is_del"`
}

func MySQLTopicTagFrom(tt *PGTopicTag) *MySQLTopicTag {

	return &MySQLTopicTag{
		TopicID: uint64(tt.TopicID),
		TagID:   uint32(tt.TagID),
		IsDel:   tt.IsDel,
	}
}

func MySQLTopicTagListFrom(ls []*PGTopicTag) []*MySQLTopicTag {
	tts := make([]*MySQLTopicTag, 0, len(ls))
	for _, tt := range ls {
		tts = append(tts, MySQLTopicTagFrom(tt))
	}
	return tts
}
