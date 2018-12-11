package db

//go:generate mockgen -destination mockFoo/mock_foo.go -source foo.go

// Repository is ...
type Repository interface {
	Create(key string, value []byte) error
	Retrieve(key string) ([]byte, error)
	Update(key string, value []byte) error
	Delete(key string) error
}
