***WARNING: THIS REPO IS AN AUTO-GENERATED COPY.*** *This repo has been copied from [Gruntwork’s](https://gruntwork.io/) GitHub repositories so that you can consume it from your company’s own internal Git repositories. This copy is automatically created and updated by the `repo-copier` CLI tool. If you need to make changes to this repo, you should make the changes in a separate fork, and NOT make changes directly in this repo, as otherwise, the `repo-copier` will overwrite your changes! Please see the `repo-copier` [documentation](https://github.com/terraform-modules-krish/repo-copier) for more information on how the code is copied, how cross-references are updated, how the changelog is handled, etc.*

***

_You may find it valuable to view the following resources in the original repo. If these links give you a 404, visit https://app.gruntwork.io to gain access or email support@gruntwork.io if you need assistance._

[Home Page](https://github.com/gruntwork-io/kubergrunt/) |
[Pull Requests](https://github.com/gruntwork-io/kubergrunt/pulls) |
[Issues](https://github.com/gruntwork-io/kubergrunt/issues) |
[Releases and Assets](https://github.com/gruntwork-io/kubergrunt/releases)

_Alternatively, you can view a copied version of the resources listed above._

[Pull Requests](https://github.com/terraform-modules-krish/kubergrunt/blob/main/.github/PULL_REQUESTS.md) |
[Issues](https://github.com/terraform-modules-krish/kubergrunt/blob/main/.github/ISSUES.md) |
[ChangeLog](https://github.com/terraform-modules-krish/kubergrunt/blob/main/.github/CHANGELOG.md)

***

[![Maintained by Gruntwork.io](https://img.shields.io/badge/maintained%20by-gruntwork.io-%235849a6.svg)](https://gruntwork.io/?ref=repo_kubergrunt)

# kubergrunt

`kubergrunt` is an encompassing tool that attempts to fill in the gaps between Terraform, Helm, and Kubectl for managing
a Kubernetes Cluster. The binaries are all built as part of the CI pipeline on each release of the package, and is
appended to the corresponding release in the [Releases Page](https://github.com/terraform-modules-krish/kubergrunt/blob/v0.1.4/../../releases).

Some of the features of `kubergrunt` includes:

* configuring `kubectl` to authenticate with a given EKS cluster. Learn more about authenticating `kubectl` to EKS
  in the [eks-cluster module README](../eks-cluster/README.md#how-to-authenticate-kubectl).
* managing Helm and associated TLS certificates.
* setting up Helm client with TLS certificates.


## Installation

You can install `kubergrunt` using the [Gruntwork Installer](https://github.com/terraform-modules-krish/gruntwork-installer):

```bash
gruntwork-install --binary-name "kubergrunt" --repo "https://github.com/terraform-modules-krish/kubergrunt" --tag "v0.0.1"
```

Alternatively, you can download the corresponding binary for your platform directly from the [Releases Page](https://github.com/terraform-modules-krish/kubergrunt/blob/v0.1.4/../../releases).


## Commands

The following commands are available as part of `kubergrunt`:

1. [eks](#eks)
    * [configure](#configure)
    * [token](#token)
    * [deploy](#deploy)

### eks

The `eks` subcommand of `kubergrunt` is used to setup the operator machine to interact with a Kubernetes cluster running
on EKS.

#### verify

This subcommand verifies that the specified EKS cluster is up and ready. An EKS cluster is considered ready when:

- The cluster status reaches ACTIVE state.
- The cluster Kubernetes API server endpoint responds to http requests.

When passing `--wait` to the command, this command will wait until the EKS cluster reaches the ready state, or it
times out. The timeout parameters are configurable with the `--max-retries` and `--sleep-between-retries` options, where
`--max-retries` specifies the number of times the command will try to verify a specific condition before giving up, and
`--sleep-between-retries` specifies the duration of time (e.g 10m = 10 minutes) to wait between each trial. So for
example, if you ran the command:

```bash
kubergrunt eks verify --eks-cluster-arn $EKS_CLUSTER_ARN --wait --max-retries 10 --sleep-between-retries 15s
```

and the cluster was not active yet, this command will query the AWS API up to 10 times, waiting 15 seconds inbetween
each try for a total of 150 seconds (2.5 minutes) before timing out.

Run `kubergrunt eks verify --help` to see all the available options.

#### configure

This subcommand will setup the installed `kubectl` with config contexts that will allow it to authenticate to a specified
EKS cluster. This binary is designed to be used as part of one of the modules in the package, although this binary
supports running as a standalone binary. For example, this binary might be used to setup a new operator machine to be
able to talk to an existing EKS cluster.

For example to setup a `kubectl` install on an operator machine to authenticate with EKS:

```bash
kubergrunt eks configure --eks-cluster-arn $EKS_CLUSTER_ARN
```

Run `kubergrunt eks configure --help` to see all the available options.

#### token

This subcommand is used by `kubectl` to retrieve an authentication token using the AWS API authenticated with IAM
credentials. This token is then used to authenticate to the Kubernetes cluster. This command embeds the
`aws-iam-authenticator` tool into `kubergrunt` so that operators don't have to install a separate tool to manage
authentication into Kubernetes.

The `configure` subcommand of `kubergrunt eks` assumes you will be using this method to authenticate with the Kubernetes
cluster provided by EKS. If you wish to use `aws-iam-authenticator` instead, replace the auth info clause of the `kubectl`
config context.

See [How do I authenticate kubectl to the EKS cluster?](../eks-cluster-control-plane/README.md#how-to-authenticate-kubectl) for more information on
authenticating `kubectl` with EKS.

#### deploy

This subcommand will initiate a rolling deployment of the current AMI config to the EC2 instances in your EKS cluster.
This command will not deploy or update an application deployed on your Kubernetes cluster (e.g `Deployment` resource,
`Pod` resource, etc). For that, refer to the [`k8s-service` module documentation](../k8s-service). Instead, this command
is for managing and deploying an update to the EC2 instances underlying your EKS cluster.

Terraform and AWS do not provide a way to automatically roll out a change to the Instances in an EKS Cluster. Due to
Terraform limitations (see [here for a
discussion](https://github.com/terraform-providers/terraform-provider-aws/issues/567)), there is currently no way to
implement this purely in Terraform code. Therefore, we've created this subcommand that can do a zero-downtime roll out
for you.

To deploy a change (such as rolling out a new AMI) to all EKS workers using this command:

1. Make sure the `cluster_max_size` is at least twice the size of `cluster_min_size`. The extra capacity will be used to
   deploy the updated instances.
1. Update the Terraform code with your changes (e.g. update the `cluster_instance_ami` variable to a new AMI).
1. Run `terraform apply`.
1. Run the command:

```bash
kubergrunt eks deploy --eks-cluster-arn CLUSTER_ARN --asg-name ASG_NAME
```

When you call the command, it will:

1. Double the desired capacity of the Auto Scaling Group that powers the EKS Cluster. This will launch new EKS workers
   with the new launch configuration.
1. Wait for the new nodes to be ready for Pod scheduling in Kubernetes. This includes waiting for the new nodes to be
   registered to any external load balancers managed by Kubernetes.
1. Drain the pods scheduled on the old EKS workers (using the equivalent of `kubectl drain`), so that they will be
   rescheduled on the new EKS workers.
1. Wait for all the pods to migrate off of the old EKS workers.
1. Set the desired capacity down to the original value and remove the old EKS workers from the ASG.

Note that to minimize service disruption from this command, your services should setup [a
PodDisruptionBudget](https://kubernetes.io/docs/tasks/run-application/configure-pdb/), [a readiness
probe](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/#define-readiness-probes)
that fails on container shutdown events, and implement graceful handling of SIGTERM in the container. You can learn more
about these features in [our blog post covering them](TODO).

Currently `kubergrunt` does not implement any checks for these resources to be implemented. However in the future, we
plan to bake in checks into the deployment command to verify that all services have a disruption budget set, and warn
the user of any services that do not have a check.

## Who maintains this project?

`kubergrunt` is maintained by [Gruntwork](http://www.gruntwork.io/). If you are looking for help or commercial support,
send an email to [support@gruntwork.io](mailto:support@gruntwork.io?Subject=kubergrunt).

Gruntwork can help with:

* Setup, customization, and support for this Module.
* Modules and submodules for other types of infrastructure, such as VPCs, Docker clusters, databases, and continuous
  integration.
* Modules and Submodules that meet compliance requirements, such as HIPAA.
* Consulting & Training on AWS, Terraform, and DevOps.


## How do I contribute?

Contributions are very welcome! Check out the [Contribution Guidelines](https://github.com/terraform-modules-krish/kubergrunt/blob/v0.1.4/CONTRIBUTING.md) for instructions.


## How is this project versioned?

This project follows the principles of [Semantic Versioning](http://semver.org/). You can find each new release, along
with the changelog, in the [Releases Page](../../releases).

During initial development, the major version will be 0 (e.g., `0.x.y`), which indicates the code does not yet have a
stable API. Once we hit `1.0.0`, we will make every effort to maintain a backwards compatible API and use the MAJOR,
MINOR, and PATCH versions on each release to indicate any incompatibilities.


## License

Please see [LICENSE](https://github.com/terraform-modules-krish/kubergrunt/blob/v0.1.4/LICENSE) for how the code in this repo is licensed.

Copyright &copy; 2018 Gruntwork, Inc.
