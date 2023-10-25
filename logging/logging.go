// Logging package includes code for managing the logger for kubergrunt
package logging

import (
	gruntwork-cli "github.com/terraform-modules-krish/go-commons/logging"
	"github.com/sirupsen/logrus"
)

func GetProjectLogger() *logrus.Entry {
	logger := logging.GetLogger("")
	return logger.WithField("name", "kubergrunt")
}
