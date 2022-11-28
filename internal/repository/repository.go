package repository

type Repository interface {
	FetchUrl(string) (string, error)
	AddUrl(string) (string, error)
}
