package repositories

type Seedable interface {
	Seed() error
}

type Migratable interface {
	Migrate() error
}
