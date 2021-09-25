// Copyright 2019 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"os"
	"testing"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin/config"
)

func TestPlugin(t *testing.T) {
	if len(os.Getenv("CI")) > 0 {
		t.Skip("this test may fail in ci env due to github connection issue, we'll skip it")
	}
	p := plugin{
		token: "", // your secret here
	}
	conf, err := p.Find(context.Background(), &config.Request{
		Repo: drone.Repo{
			Namespace: "Pivot-Studio",
			Name:      "masterext",
			Config:    ".drone.yml",
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	if conf == nil {
		t.Errorf("conf should not be nil")
	}
}
