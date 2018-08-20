package command

import "github.com/ablease/zinn/gw2api"

// Client ...
//go:generate counterfeiter . ApiClient
type ApiClient interface {
	Professions() ([]string, error)
	Masteries() ([]gw2api.Mastery, error)
}
