package repositories

type ISeed interface {
	Seed() error
}

type IMigrate interface {
	Migrate() error
}
