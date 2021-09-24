// Copyright 2019 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"testing"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin/config"
)

func TestPlugin(t *testing.T) {
	t.Skip()
	p := plugin{
		token: "<secret>",
	}
	conf, _ := p.Find(context.Background(), &config.Request{
		Repo: drone.Repo{
			Namespace: "Pivot-Studio",
			Name:      "HUSTHoleBackEnd",
		},
	})
	if conf == nil {
		t.Errorf("conf should not be nil")
	}
}
