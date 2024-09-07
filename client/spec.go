package client

import (
	"fmt"
)

type Spec struct {
	InstanceDomain string `json:"instance_domain"`
	AuthMethod     string `json:"auth_method"`
	ClientID       string `json:"client_id"`
	ClientSecret   string `json:"client_secret"`
	Username       string `json:"basic_auth_username"`
	Password       string `json:"basic_auth_password"`
}

func (s *Spec) validate() error {
	if s == nil {
		return fmt.Errorf(`required field "jamf_client_config" is missing`)
	}

	if len(s.InstanceDomain) == 0 {
		return fmt.Errorf(`required field "jamf_client_config.instance_domain" is missing`)
	}

	if len(s.AuthMethod) == 0 {
		return fmt.Errorf(`required field "jamf_client_config.auth_method" is missing`)
	}

	if s.AuthMethod != "basic" && s.AuthMethod != "oauth" {
		return fmt.Errorf(`invalid value for "jamf_client_config.auth_method". Must be one of: "basic", "oauth"`)
	}

	if s.AuthMethod == "basic" {
		if len(s.Username) == 0 {
			return fmt.Errorf(`required field "jamf_client_config.username" is missing`)
		}

		if len(s.Password) == 0 {
			return fmt.Errorf(`required field "jamf_client_config.password" is missing`)
		}
	}

	if s.AuthMethod == "oauth" {
		if len(s.ClientID) == 0 {
			return fmt.Errorf(`required field "jamf_client_config.client_id" is missing`)
		}

		if len(s.ClientSecret) == 0 {
			return fmt.Errorf(`required field "jamf_client_config.client_secret" is missing`)
		}
	}

	return nil
}
