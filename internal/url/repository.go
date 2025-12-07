package url

type Repository interface {
	Save(*URL) error
	FindByKey(string) (*URL, error)
}
