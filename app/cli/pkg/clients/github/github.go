package github

import (
	"fmt"
	rest "github.com/glonner/pkg/adapters"
	"github.com/glonner/pkg/common"
	logger "github.com/glonner/pkg/log"
	"net/http"
)

const (
	base string = "https://api.github.com"
	api  string = "repos"
)

type IGitHubClient interface {
	GetURL(org string, page int) string
	GetAuth(token string) (map[string][]string, error)
	GetData(uri string, auth map[string][]string) (*http.Response, error)
}

type Options struct {
	DisablePagination bool
	OwnerOrOrg        string
}

type Client struct {
	Token  string
	Owner  string
	logger logger.ILogger
	rest   rest.IRestAdapter
}

type Auth struct {
	Auth map[string][]string
}

func (g Client) GetURL(org string, page int) string {
	var uri string
	if page > 1 {
		uri = fmt.Sprintf("%s/orgs/%s/%s?type=all&per_page=100&page=%d", base, org, api, page)
	} else {
		uri = fmt.Sprintf("%s/orgs/%s/%s?type=all&per_page=100", base, org, api)
	}
	g.logger.LogDebugF("GitHub URL formed: %s", uri)
	return uri
}

func (g Client) GetAuth(token string) (map[string][]string, error) {
	auth := make(map[string][]string)

	var tokenResolved string

	if token == "" {
		tokenFromEnv, err := common.GetEnv("GITHUB_TOKEN")

		if err != nil {
			return auth, err
		}

		tokenResolved = tokenFromEnv
	} else {
		tokenResolved = token
	}

	auth["Authorization"] = []string{fmt.Sprintf("token %s", tokenResolved)}
	auth["Accept"] = []string{"application/vnd.github+jso"}
	auth["Content-Type"] = []string{"application/json"}

	return auth, nil
}

func (g Client) GetData(uri string, auth map[string][]string) (*http.Response, error) {
	resp, err := g.rest.Get(uri, auth)
	if err != nil {
		g.logger.LogError(err)
		return nil, err
	}

	return resp, nil
}

func NewGitHubClient(log *logger.ILogger) IGitHubClient {
	return &Client{
		logger: *log,
		rest:   rest.NewRestAdapter(log),
	}
}
