package main

import (
	"log"

	"github.com/axumrs/axum-rs-migration/global"
	"github.com/axumrs/axum-rs-migration/migrate"
)

func init() {
	global.Init()
}

func main() {
	s := migrate.Subject{Truncate: true}
	t := migrate.Topic{Truncate: true}
	tc := migrate.TopicContent{Truncate: true}
	tg := migrate.Tag{Truncate: true}
	tctg := migrate.TopicTag{Truncate: true}

	log.Println("开始迁移")

	if err := s.Migrate(); err != nil {
		log.Fatal("subject 迁移失败:", err)
	}
	log.Println("subject 迁移成功")

	if err := t.Migrate(); err != nil {
		log.Fatal("topic 迁移失败:", err)
	}
	log.Println("topic 迁移成功")

	if err := tc.Migrate(); err != nil {
		log.Fatal("topic_content 迁移失败:", err)
	}
	log.Println("topic_content 迁移成功")

	if err := tg.Migrate(); err != nil {
		log.Fatal("tag 迁移失败:", err)
	}
	log.Println("tag 迁移成功")

	if err := tctg.Migrate(); err != nil {
		log.Fatal("topic_tag 迁移失败:", err)
	}
	log.Println("topic_tag 迁移成功")
}
