package issue

import (
	"context"
	"fmt"

	"github.com/dywoq/gh-utility/base"
	"github.com/google/go-github/v80/github"
	"golang.org/x/oauth2"
)

// GetById gets the issue by ID.
// Returns an error if there is no issue with this ID.
func GetById(ctx context.Context, conf *base.Config, id int) (*github.Issue, error) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: conf.Token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	issue, _, err := client.Issues.Get(ctx, conf.Owner, conf.Repository, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get issue #%d: %w", id, err)
	}
	return issue, nil
}
