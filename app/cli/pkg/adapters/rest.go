package rest

import (
	logger "github.com/glonner/pkg/log"
	"net/http"
	"time"
)

type IRestAdapter interface {
	Get(url string, headers map[string][]string) (*http.Response, error)
	// TODO: Complement more operations in the future. Till now, isn't required ;)
}

type Adapter struct {
	logger *logger.ILogger
}

func (r *Adapter) Get(url string, headers map[string][]string) (*http.Response, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	req, _ := http.NewRequest("GET", url, nil)

	req.Header = headers

	return client.Do(req)
}

func NewRestAdapter(logger *logger.ILogger) IRestAdapter {
	return &Adapter{
		logger: logger,
	}
}
