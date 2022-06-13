// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"errors"
	"github.com/docker/cli/cli/config"
	"github.com/docker/cli/cli/config/types"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"log"
	"os"
)

// Args provides plugin execution arguments.
type Args struct {
	Pipeline

	// Level defines the plugin log level.
	Level string `envconfig:"PLUGIN_LOG_LEVEL"`

	ServerAddress string `envconfig:"PLUGIN_SERVER_ADDRESS"`
	User          string `envconfig:"PLUGIN_USERNAME"`
	Password      string `envconfig:"PLUGIN_PASSWORD"`
}

// Exec executes the plugin.
func Exec(ctx context.Context, args Args) error {

	if args.User == "" && args.Password == "" {
		return errors.New("username and password required")
	}

	// Some code reused from https://github.com/google/go-containerregistry/blob/main/cmd/crane/cmd/auth.go
	cf, err := config.Load(os.Getenv("DOCKER_CONFIG"))
	if err != nil {
		return err
	}
	creds := cf.GetCredentialsStore(args.ServerAddress)
	if args.ServerAddress == name.DefaultRegistry {
		args.ServerAddress = authn.DefaultAuthKey
	}
	if err := creds.Store(types.AuthConfig{
		ServerAddress: args.ServerAddress,
		Username:      args.User,
		Password:      args.Password,
	}); err != nil {
		return err
	}

	if err := cf.Save(); err != nil {
		return err
	}
	log.Printf("logged in via %s", cf.Filename)

	return nil
}
