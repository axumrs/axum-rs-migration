package model

type PGTopicContent struct {
	TopicID int64  `db:"topic_id"`
	MD      string `db:"md"`
	HTML    string `db:"html"`
}
type MySQLTopicContent struct {
	TopicID uint64 `db:"topic_id"`
	MD      string `db:"md"`
	HTML    string `db:"html"`
}

func MySQLTopicContentFrom(tc *PGTopicContent) *MySQLTopicContent {
	return &MySQLTopicContent{
		TopicID: uint64(tc.TopicID),
		MD:      tc.MD,
		HTML:    tc.HTML,
	}
}

func MySQLTopicContentListFrom(ls []*PGTopicContent) []*MySQLTopicContent {
	tcs := make([]*MySQLTopicContent, 0, len(ls))
	for _, tc := range ls {
		tcs = append(tcs, MySQLTopicContentFrom(tc))
	}
	return tcs
}
