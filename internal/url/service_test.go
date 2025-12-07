package url

import (
	"testing"
)

type MockRepository struct{}

func (r *MockRepository) Save(_ *URL) error {
	return nil
}

func (r *MockRepository) FindByKey(_ string) (*URL, error) {
	return nil, nil
}

func TestNewService(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Log("new service panics without a repository")
		}
	}()

	_ = NewService(nil)

	t.Errorf("new service did not panic when not passed a repository")
}

func TestNewService2(t *testing.T) {
	repo := MockRepository{}
	service := NewService(&repo)

	if service.repo == nil {
		t.Errorf("the repository is not correctly assigned to the service")
	}
}

func TestService_NewURL(t *testing.T) {
	originalURL := "https://random.com/test-url"
	service := NewService(&MockRepository{})

	url, err := service.NewURL(originalURL, "g.cd")

	if err != nil {
		t.Fatal("the URL creation thrown a error")
	}

	if url.Original != originalURL {
		t.Errorf("unable to set the correct properties when creating a new URL")
	}
}
