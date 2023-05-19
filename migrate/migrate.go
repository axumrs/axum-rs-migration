package migrate

type Migrater interface {
	Migrate() error
}
