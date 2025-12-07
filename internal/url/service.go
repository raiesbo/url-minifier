package url

import (
	"errors"
	"fmt"
	"path/filepath"
	"time"

	"github.com/raiesbo/url-minifier/internal/utils"
)

var (
	ErrMissingRepository = errors.New("repository not provided")
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	if r == nil {
		panic(ErrMissingRepository)
	}

	return &Service{repo: r}
}

func (s *Service) NewURL(longURL string, host string) (*URL, error) {
	urlKey := utils.CreateURLKey(10)

	url := &URL{
		Slug:      urlKey,
		Original:  longURL,
		Short:     filepath.Join(host, urlKey),
		SecretKey: utils.CreateURLKey(10),
		Clicks:    0,
		CreatedAt: time.Now(),
	}

	return url, s.repo.Save(url)
}

// Store triggers the corresponding repo function passing down the args
func (s *Service) Store(url *URL) error {
	return s.repo.Save(url)
}

// FindBySlug triggers the corresponding repo function passing down the args
func (s *Service) FindBySlug(slug string) (*URL, error) {
	return s.repo.FindByKey(slug)
}

// FindByLongURL searches based on the long version of the URL
func (s *Service) FindByLongURL(originalURL string) (*URL, error) {
	original := fmt.Sprintf("idx:original:%s", originalURL)
	return s.repo.FindByKey(original)
}

// UpdateCounter increments the visit counter of a URL
func (s *Service) UpdateCounter(key string) error {
	url, err := s.repo.FindByKey(key)
	if err != nil {
		return err
	}
	url.Clicks++
	return s.repo.Save(url)
}
