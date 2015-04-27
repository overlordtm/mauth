package mauth

import "net/url"

func extractPath(next string) string {
	n, err := url.Parse(next)
	if err != nil {
		return "/"
	}

	return n.Path
}

func autofillConfig(config *Config) {

	if config.Context == nil {
		config.Context = GetContext
	}

	if config.PathError == "" {
		config.PathError = PathError
	}

	if config.CodeRedirect == 0 {
		config.CodeRedirect = CodeRedirect
	}
}
