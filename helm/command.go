package helm

import (
	"os"

	gruntwork-cli "github.com/terraform-modules-krish/go-commons/shell"

	"github.com/terraform-modules-krish/kubergrunt/kubectl"
)

// RunHelm will make a call to helm, setting the config and context to the ones specified in the provided options.
func RunHelm(options *kubectl.KubectlOptions, args ...string) error {
	_, err := RunHelmAndGetOutput(options, args...)
	return err
}

// RunHelmAndGetOutput will make a call to helm, setting the config and context to the ones specified in the provided
// options, and return the output of the command.
func RunHelmAndGetOutput(options *kubectl.KubectlOptions, args ...string) (string, error) {
	shellOptions := shell.NewShellOptions()
	cmdArgs := []string{}

	// If KubectlOptions is configured with a direct auth info, use that instead of the context. Note that since helm
	// does not support directly using auth infos as CLI args, we create a tmp file that holds this info.
	scheme := options.AuthScheme()
	switch scheme {
	case kubectl.ConfigBased:
		if options.ContextName != "" {
			cmdArgs = append(cmdArgs, "--kube-context", options.ContextName)
		}
		if options.ConfigPath != "" {
			cmdArgs = append(cmdArgs, "--kubeconfig", options.ConfigPath)
		}
	default:
		tmpfile, err := options.TempConfigFromAuthInfo()
		if tmpfile != "" {
			// Make sure to delete the tmp file at the end
			defer os.Remove(tmpfile)
		}
		if err != nil {
			return "", err
		}
		cmdArgs = append(cmdArgs, "--kubeconfig", tmpfile)
	}

	cmdArgs = append(cmdArgs, args...)
	return shell.RunShellCommandAndGetAndStreamOutput(shellOptions, "helm", cmdArgs...)
}
