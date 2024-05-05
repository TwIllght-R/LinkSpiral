package usecases

type UrlMappingRepo interface {
	CreateUrlMapping(originalURL string) (string, error)
}
