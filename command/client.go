package command

import "github.com/ablease/zinn/api"

// Client ...
//
//go:generate counterfeiter . ApiClient
type ApiClient interface {
	Professions() ([]string, error)
	GetMasteryIDs() ([]int, error)
	Masteries(ids []int) ([]api.Mastery, error)
	AchievementIDs() ([]int, error)
	Achievements(ids []int) ([]api.Achievement, error)
	DailyCrafting() ([]string, error)
	MapChests() ([]string, error)
}
