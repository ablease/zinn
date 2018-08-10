package command

// ZinnClient ...
//go:generate counterfeiter . ApiClient
type ApiClient interface {
	Professions() ([]string, error)
}
