package command

// Client ...
//go:generate counterfeiter . ApiClient
type ApiClient interface {
	Professions() ([]string, error)
}
