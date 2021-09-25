// Copyright 2019 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"io"
	"net/http"
	"sync"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin/config"
	"github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

var (
	once        = sync.Once{}
	oauthClient *http.Client
)

// New returns a new config plugin.
func New(token string) config.Plugin {
	return &plugin{
		// TODO replace or remove these configuration
		// parameters. They are for demo purposes only.
		token: token,
	}
}

type plugin struct {
	// TODO replace or remove these configuration
	// parameters. They are for demo purposes only.
	token string
}

func (p *plugin) Find(ctx context.Context, req *config.Request) (*drone.Config, error) {
	// creates a github client used to fetch the yaml.
	logrus.Infoln("got request", req)
	trans := http.DefaultClient
	if len(p.token) > 0 {
		once.Do(func() {
			oauthClient = oauth2.NewClient(ctx, oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: p.token},
			))
		})
		trans = oauthClient
	}
	c := github.NewClient(trans)
	repo, _, err := c.Repositories.Get(ctx, req.Repo.Namespace, req.Repo.Name)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	reader, err := c.Repositories.DownloadContents(ctx, req.Repo.Namespace, req.Repo.Name, req.Repo.Config,
		&github.RepositoryContentGetOptions{
			Ref: *repo.DefaultBranch,
		})
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	bs, err := io.ReadAll(reader)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	logrus.Infoln("request master config success")
	return &drone.Config{
		Data: string(bs),
	}, nil
}
