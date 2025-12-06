package issue

import (
	"context"
	"fmt"
	"sync"

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

// GoGetById gets all issues with their IDs synchronously.
// Causes an early-exit if the error is met, and the function returns it.
func GoGetById(ctx context.Context, conf *base.Config, ids []int) ([]*github.Issue, error) {
	var (
		issues []*github.Issue
		goterr error
		wg     sync.WaitGroup
		m      sync.Mutex
	)

	for _, id := range ids {
		wg.Go(func() {
			i, err := GetById(ctx, conf, id)
			m.Lock()
			defer m.Unlock() 
			if err != nil {
				if goterr == nil {
					goterr = err
				}
				return
			}
			issues = append(issues, i)
		})
	}

	wg.Wait()

	return issues, goterr
}
