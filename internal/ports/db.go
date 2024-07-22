package ports

type DB interface {
	CloseDB() (err error)
}
