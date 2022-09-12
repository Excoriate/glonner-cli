package actions

import (
	"github.com/glonner/pkg/common"
	"github.com/pkg/errors"
)

type GlobalRequiredArgs struct {
	Organization string
	Token        string
}

func GetRequiredOrgAndToken(args *GlobalRequiredArgs) (*GlobalRequiredArgs, error) {
	var org string
	var token string
	if args.Organization == "" {
		value, err := common.GetEnv("GITHUB_ORG")
		if err != nil {
			return nil, errors.New("Organization is required. " +
				"Please set the GitHub organization as an environment variable (" +
				"GITHUB_ORG) or pass it as" +
				" a flag --org")
		} else {
			org = value
		}
	} else {
		org = args.Organization
	}

	if args.Token == "" {
		value, err := common.GetEnv("GITHUB_TOKEN")
		if err != nil {
			return nil, errors.New("Token is required. " +
				"Please set the GitHub token as an environment variable (" +
				"GITHUB_TOKEN) or pass it as" +
				" a flag --token")
		} else {
			token = value
		}
	} else {
		token = args.Token
	}

	return &GlobalRequiredArgs{
		Organization: org,
		Token:        token,
	}, nil
}
