from dotenv import dotenv_values
from migrate import user, subject,tag,topic,topic_tag


def main():
    cfg = dotenv_values(".env")

    n = subject.migration(cfg)
    print('已迁移 {} 条专题数据'.format(n))

    n = tag.migration(cfg)
    print('已迁移 {} 条标签数据'.format(n))

    (n,htmls) = topic.migration(cfg)
    print('已迁移 {} 条文章数据'.format(n))
    n = topic.migration_content(cfg,htmls)
    print('已迁移 {} 条文章章节数据'.format(n))

    n = topic_tag.migration(cfg)
    print('已迁移 {} 条文章标签数据'.format(n))

    # n = user.migration(cfg)
    # print('已迁移 {} 条用户数据'.format(n))

if __name__ == "__main__":
    main()