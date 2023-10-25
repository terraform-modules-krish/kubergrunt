package kubectl

import (
	"testing"

	"github.com/terraform-modules-krish/terratest/modules/k8s"
	"github.com/stretchr/testify/require"
)

func GetTestKubectlOptions(t *testing.T) *KubectlOptions {
	kubeConfigPath, err := k8s.GetKubeConfigPathE(t)
	require.NoError(t, err)
	return NewKubectlOptions("", kubeConfigPath)
}
