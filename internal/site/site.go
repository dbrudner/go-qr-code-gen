package site

type Site struct {
	ID          string
	Description string
	URL         string
}

func (s Site) GetURL() string {
	return s.URL
}

type URLHolder interface {
	GetURL() string
}
