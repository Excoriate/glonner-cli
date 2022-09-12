package services

import (
	"encoding/json"
	"github.com/glonner/pkg/clients/github"
	logger "github.com/glonner/pkg/log"
	"github.com/glonner/pkg/system"
	"github.com/pkg/errors"
	"strings"
	"time"
)

type RepoDataSimplified struct {
	Name        string
	CloneURL    string
	Description string
	Size        int
}

type RepoData struct {
	ID       int    `json:"id"`
	NodeID   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Owner    struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"owner"`
	Private          bool        `json:"private"`
	HTMLURL          string      `json:"html_url"`
	Description      string      `json:"description"`
	Fork             bool        `json:"fork"`
	URL              string      `json:"url"`
	ArchiveURL       string      `json:"archive_url"`
	AssigneesURL     string      `json:"assignees_url"`
	BlobsURL         string      `json:"blobs_url"`
	BranchesURL      string      `json:"branches_url"`
	CollaboratorsURL string      `json:"collaborators_url"`
	CommentsURL      string      `json:"comments_url"`
	CommitsURL       string      `json:"commits_url"`
	CompareURL       string      `json:"compare_url"`
	ContentsURL      string      `json:"contents_url"`
	ContributorsURL  string      `json:"contributors_url"`
	DeploymentsURL   string      `json:"deployments_url"`
	DownloadsURL     string      `json:"downloads_url"`
	EventsURL        string      `json:"events_url"`
	ForksURL         string      `json:"forks_url"`
	GitCommitsURL    string      `json:"git_commits_url"`
	GitRefsURL       string      `json:"git_refs_url"`
	GitTagsURL       string      `json:"git_tags_url"`
	GitURL           string      `json:"git_url"`
	IssueCommentURL  string      `json:"issue_comment_url"`
	IssueEventsURL   string      `json:"issue_events_url"`
	IssuesURL        string      `json:"issues_url"`
	KeysURL          string      `json:"keys_url"`
	LabelsURL        string      `json:"labels_url"`
	LanguagesURL     string      `json:"languages_url"`
	MergesURL        string      `json:"merges_url"`
	MilestonesURL    string      `json:"milestones_url"`
	NotificationsURL string      `json:"notifications_url"`
	PullsURL         string      `json:"pulls_url"`
	ReleasesURL      string      `json:"releases_url"`
	SSHURL           string      `json:"ssh_url"`
	StargazersURL    string      `json:"stargazers_url"`
	StatusesURL      string      `json:"statuses_url"`
	SubscribersURL   string      `json:"subscribers_url"`
	SubscriptionURL  string      `json:"subscription_url"`
	TagsURL          string      `json:"tags_url"`
	TeamsURL         string      `json:"teams_url"`
	TreesURL         string      `json:"trees_url"`
	CloneURL         string      `json:"clone_url"`
	MirrorURL        string      `json:"mirror_url"`
	HooksURL         string      `json:"hooks_url"`
	SvnURL           string      `json:"svn_url"`
	Homepage         string      `json:"homepage"`
	Language         interface{} `json:"language"`
	ForksCount       int         `json:"forks_count"`
	StargazersCount  int         `json:"stargazers_count"`
	WatchersCount    int         `json:"watchers_count"`
	Size             int         `json:"size"`
	DefaultBranch    string      `json:"default_branch"`
	OpenIssuesCount  int         `json:"open_issues_count"`
	IsTemplate       bool        `json:"is_template"`
	Topics           []string    `json:"topics"`
	HasIssues        bool        `json:"has_issues"`
	HasProjects      bool        `json:"has_projects"`
	HasWiki          bool        `json:"has_wiki"`
	HasPages         bool        `json:"has_pages"`
	HasDownloads     bool        `json:"has_downloads"`
	Archived         bool        `json:"archived"`
	Disabled         bool        `json:"disabled"`
	Visibility       string      `json:"visibility"`
	PushedAt         time.Time   `json:"pushed_at"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
	Permissions      struct {
		Admin bool `json:"admin"`
		Push  bool `json:"push"`
		Pull  bool `json:"pull"`
	} `json:"permissions"`
	TemplateRepository interface{} `json:"template_repository"`
}

type IGitHubSvc interface {
	GetRepositories(org string, token string) ([]RepoData,
		error)
	GetSimplifiedRepositoryData(repo RepoData) []RepoDataSimplified
	CloneRepository(cloneURL string, dirPath string) (bool, error)
}

type GitHubSvc struct {
	client github.IGitHubClient
	logger *logger.ILogger
}

func (s GitHubSvc) GetRepositories(org, token string) ([]RepoData,
	error) {
	auth, err := s.client.GetAuth(token)
	if err != nil {
		return nil, err
	}

	var repositories []RepoData
	page := 1
	for {
		var reposIntermediateData []RepoData
		uri := s.client.GetURL(org, page)
		res, err := s.client.GetData(uri, auth)

		if err != nil {
			return nil, err
		}

		err = json.NewDecoder(res.Body).Decode(&reposIntermediateData)

		if len(reposIntermediateData) == 0 {
			break
		} else {
			repositories = append(repositories, reposIntermediateData...)
		}

		if err != nil {
			return nil, err
		}

		page++
	}

	return repositories, nil
}

func (s GitHubSvc) GetSimplifiedRepositoryData(repo RepoData) []RepoDataSimplified {
	var data []RepoDataSimplified
	data = append(data, RepoDataSimplified{
		Name:        repo.Name,
		Description: repo.Description,
		CloneURL:    repo.CloneURL,
		Size:        repo.Size,
	})

	return data
}

func (s GitHubSvc) CloneRepository(cloneULR, dirPath string) (bool, error) {
	out, err := system.RunCMD(*s.logger, "git", "clone", cloneULR, dirPath)

	if err != nil {
		return false, err
	}

	if out != "" {
		if strings.Contains(out, "fatal") {
			return false, errors.New(out)
		}
	}

	return true, nil
}

func NewGitHubSvc(log *logger.ILogger) IGitHubSvc {
	return &GitHubSvc{
		client: github.NewGitHubClient(log),
		logger: log,
	}
}
