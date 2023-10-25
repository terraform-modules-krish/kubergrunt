module github.com/terraform-modules-krish/kubergrunt

go 1.14

require (
	github.com/aws/aws-sdk-go v1.34.27
	github.com/terraform-modules-krish/go-commons v0.7.0
	github.com/terraform-modules-krish/terratest v0.30.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/sirupsen/logrus v1.6.0
	github.com/stretchr/testify v1.6.1
	github.com/urfave/cli v1.22.4

	// EKS is still on k8s v1.17
	k8s.io/api v0.18.3
	k8s.io/apimachinery v0.18.3
	k8s.io/client-go v0.18.3
	sigs.k8s.io/aws-iam-authenticator v0.5.1
)
