package client

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/yohan460/go-jamf-api"
)

type Client struct {
	logger     zerolog.Logger
	Spec       Spec
	JamfClient *jamf.Client
}

func (c *Client) ID() string {
	return "jamf:instance:{" + c.Spec.InstanceDomain + "}"
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func New(ctx context.Context, logger zerolog.Logger, s *Spec) (Client, error) {
	c := Client{
		logger: logger.With().
			Str("plugin", "jamf").
			Logger(),
		Spec: *s,
	}

	if err := s.validate(); err != nil {
		return c, err
	}

	var authMethod jamf.Option

	if s.AuthMethod == "basic" {
		authMethod = jamf.WithBasicAuth(s.Username, s.Password)
	} else if s.AuthMethod == "oauth" {
		authMethod = jamf.WithOAuth(s.ClientID, s.ClientSecret)
	}

	// Remove https:// from the instance domain if it exists
	// The go-jamf-api client automatically adds https:// to the instance domain
	cleanedInstanceDomain := s.InstanceDomain
	if cleanedInstanceDomain[:8] == "https://" {
		cleanedInstanceDomain = cleanedInstanceDomain[8:]
	}

	client, err := jamf.NewClient(
		cleanedInstanceDomain,
		authMethod,
	)

	if err != nil {
		return c, err
	}

	c.JamfClient = client

	return c, nil
}
