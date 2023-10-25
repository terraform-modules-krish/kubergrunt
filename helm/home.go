package helm

import (
	"path/filepath"

	gruntwork-cli "github.com/terraform-modules-krish/go-commons/errors"
	homedir "github.com/mitchellh/go-homedir"
)

func GetDefaultHelmHome() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return home, errors.WithStackTrace(err)
	}
	helmHome := filepath.Join(home, ".helm")
	return helmHome, nil
}
