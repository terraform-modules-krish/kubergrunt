# Changelog


# v0.9.3-alpha
_Released Sep 6, 2022_
Alpha release for regression testing `kubergrunt` support for `1.23` in `terraform-aws-eks`.
<br>

# v0.9.3
_Released Sep 8, 2022_
## Commands

- `eks sync-core-components`

## Description

Add support for Kubernetes version 1.23 and drop support for 1.19 in the `sync-core-components` command.

## Related links

#174 
<br>

# v0.9.2
_Released Aug 1, 2022_
## Commands

- `eks configure`

## Description

Updated dependency `aws-iam-authenticator` to incorporate security patches.

## Related links

https://github.com/gruntwork-io/kubergrunt/pull/170
<br>

# v0.9.1
_Released Jun 20, 2022_
## Commands

- `eks configure`

## Description

Updated dependency `aws-iam-authenticator` so that `v1beta1` versioned tokens can be generated.

## Related links

https://github.com/gruntwork-io/kubergrunt/pull/166
<br>

# v0.9.0
_Released Apr 15, 2022_
## Commands

- `eks sync-core-components`

## Description

This adds support for Kubernetes version 1.22 in the `sync-core-components` command. In turn, support for Kubernetes 1.16, 1.17, and 1.18 has been dropped as they reached end of support in EKS.

## Related links

https://github.com/gruntwork-io/kubergrunt/pull/157
<br>

# v0.8.0
_Released Jan 25, 2022_
## Commands

- `eks deploy`

## Description

Improve fault tolerance for `eks deploy` subcommand by storing deploy state in a recovery file and being able to resume from the last successful stage.

## Related links

#154 
<br>

# v0.7.9
_Released Aug 25, 2021_
## Commands

- `eks sync-core-components`

## Description
Updated `eks sync-core-components` to deploy `aws-vpc-cni` version 1.9.

## Related links
https://github.com/gruntwork-io/kubergrunt/pull/144
<br>

# v0.7.8
_Released Aug 25, 2021_
## Commands

- `eks deploy`

## Description
Fixed a bug in `eks deploy` where it did not correctly identify ALBs managed by the AWS LoadBalancer Controller.

## Related links
https://github.com/gruntwork-io/kubergrunt/pull/142
<br>

# v0.7.7
_Released Aug 20, 2021_
## Commands

- `eks sync-core-components`

## Description
Fixed a bug in `eks sync-core-components` where the `eksbuild` tag to use for `coredns` and `kube-proxy` depended on the region of the EKS cluster. `kubergrunt` will now identify the correct `eksbuild` tag to use for the specific region of the cluster.

## Related links
https://github.com/gruntwork-io/kubergrunt/pull/137
<br>

# v0.7.6
_Released Aug 17, 2021_
## Commands

- `eks sync-core-components`

## Description
Updated `eks sync-core-components` command to use the correct version of coredns deployed on EKS clusters with kubernetes version 1.21 (CoreDNS `v1.8.4`).

## Related links
https://github.com/gruntwork-io/kubergrunt/pull/134
<br>

# v0.7.5
_Released Aug 17, 2021_
## Commands

- `eks deploy`

## Description

Fixed a bug where `eks deploy` crashes with NLBs that have more than one target group.

## Related links
https://github.com/gruntwork-io/kubergrunt/pull/135
<br>

# v0.7.4
_Released Aug 12, 2021_
## Commands

- `eks deploy`

## Description

Fixed a bug where `eks deploy` gets stuck waiting for load balancers to register when the `Service` load balancer is using a NLB.

## Related links

https://github.com/gruntwork-io/kubergrunt/pull/132
<br>

# v0.7.3
_Released Jul 21, 2021_
## Commands

- `eks sync-core-components`

## Description

This adds support for Kubernetes version 1.21 in the `sync-core-components` command.

## Related links

https://github.com/gruntwork-io/kubergrunt/pull/131
<br>

# v0.7.2
_Released Jun 15, 2021_
## Commands

- `eks sync-core-components`

## Description

This fixes a bug in the core components update for Kubernetes version `1.20`, where `coredns` required additional permissions. The `sync-core-components` command will now patch the `ClusterRole` with the necessary permissions.


## Related links

https://github.com/gruntwork-io/kubergrunt/pull/130
<br>

# v0.7.11
_Released Nov 19, 2021_
## Commands

- `eks sync`

## Description
Updated `eks sync` to work with all regions.

## Related links
https://github.com/gruntwork-io/kubergrunt/pull/151
<br>

# v0.7.10
_Released Oct 6, 2021_
## Commands

- `eks deploy`

## Description
Updated `eks deploy` to automatically expand the max size of the ASG if there is not enough capacity to replace all the current nodes in the ASG.

## Related links
https://github.com/gruntwork-io/kubergrunt/pull/146
<br>

# v0.7.1
_Released Jun 3, 2021_
## Commands

- `eks sync-core-components`

## Description

This adds support for skipping the specific components in the `sync-core-components` command. Review the updated command docs (`kubergrunt eks sync-core-components --help`) for more info.

## Related links

https://github.com/gruntwork-io/kubergrunt/pull/128
<br>

# v0.7.0
_Released May 28, 2021_
## Commands

- `eks sync-core-components`

## Description

This adds support for Kubernetes version 1.20 in the `sync-core-components` command. In turn, support for Kubernetes 1.14 and 1.15 has been dropped.

## Related links

https://github.com/gruntwork-io/kubergrunt/pull/127
<br>

# v0.6.9
_Released Dec 19, 2020_
## Commands Affected

- `eks cleanup-security-group` ([docs](https://github.com/gruntwork-io/kubergrunt/blob/master/README.md#cleanup-security-group))

## Description
Fix bug where the cleanup routine did not work in accounts without the default VPC configured.

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/106
<br>

# v0.6.8
_Released Dec 16, 2020_
## Commands Affected

- `eks cleanup-security-group` ([docs](https://github.com/gruntwork-io/kubergrunt/blob/master/README.md#cleanup-security-group))

## Description
We discovered a bug in the retry logic that handles ENIs detached before the request to detach can be sent, but there is no error returned.

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/105
<br>

# v0.6.7
_Released Dec 10, 2020_
## Commands Affected

- `eks cleanup-security-group` ([docs](https://github.com/gruntwork-io/kubergrunt/blob/master/README.md#cleanup-security-group))

## Description
We discovered a bug in the logic that handles errors for the case when the ENI is detached before the request to detach can be sent.

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/104
<br>

# v0.6.6
_Released Dec 10, 2020_
## Commands Affected

- `eks cleanup-security-group` ([docs](https://github.com/gruntwork-io/kubergrunt/blob/master/README.md#cleanup-security-group))

## Description
We increased the timeout for deleting resources leftover after destroying the EKS cluster.

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/103
<br>

# v0.6.5
_Released Dec 8, 2020_
## New Commands Added

- `eks cleanup-security-group` ([docs](https://github.com/gruntwork-io/kubergrunt/blob/master/README.md#cleanup-security-group))
- `eks schedule-coredns ec2` ([docs](https://github.com/gruntwork-io/kubergrunt/blob/master/README.md#schedule-coredns))
- `eks schedule-coredns fargate` ([docs](https://github.com/gruntwork-io/kubergrunt/blob/master/README.md#schedule-coredns))

## Description
We have added new commands to help with cleaning up security groups associated with an EKS cluster, and to toggle the CoreDNS service between scheduling on Fargate and EC2 worker types.

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/100
- https://github.com/gruntwork-io/kubergrunt/pull/102
<br>

# v0.6.4
_Released Oct 28, 2020_
## Commands Affected

- `eks verify`

## Description
`eks verify` now uses the right kubergrunt executable, which is the currently running kubergrunt binary. Before, it was assuming that the currently running kubergrunt binary was in the PATH environment variable. That is not always the case, and it's safer to use os.Executable() to get the absolute path of kubergrunt.

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/99
<br>

# v0.6.3
_Released Oct 22, 2020_
## Commands Affected

- `eks sync-core-components`

## Description
`eks sync-core-components` now supports Kubernetes version 1.18.

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/98
<br>

# v0.6.2
_Released Oct 13, 2020_
## Commands Affected

- `eks sync-core-components`

## Description

The VPC CNI version that `kubergrunt eks sync-core-components` will upgrade to has been bumped to the latest version recommended by AWS (`v1.7.5`).

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/97
<br>

# v0.6.16
_Released May 13, 2021_
## Commands

- `eks drain` [**NEW**]

## Description

This introduces the `eks drain` subcommand which can be used to cordon and drain all the workers that are associated with the provided AWS Auto Scaling Groups. Refer to the [command docs](https://github.com/gruntwork-io/kubergrunt#drain) for more information.


## Related links

https://github.com/gruntwork-io/kubergrunt/pull/126
<br>

# v0.6.15
_Released May 3, 2021_
## Description

All commands that require AWS access will now properly support credentials configuration defined in `~/.aws/config`.

## Special thanks

Special thanks to @jtdoepke for their contribution!

## Related links

https://github.com/gruntwork-io/kubergrunt/pull/123
<br>

# v0.6.14
_Released Apr 29, 2021_
## Commands

- `eks oidc-thumbprint`

## Description

The `eks oidc-thumbprint` command now supports HTTP proxies.

## Special thanks

Special thanks to @DanielRis for their contribution!

## Related links

https://github.com/gruntwork-io/kubergrunt/pull/124
<br>

# v0.6.13
_Released Apr 16, 2021_
## Description

Starting this release, we will publish darwin/arm64 (support for M1 chipset) binaries and drop support for darwin/386.

## Related links

https://github.com/gruntwork-io/kubergrunt/pull/120
<br>

# v0.6.12
_Released Apr 12, 2021_
## Description

Update the `aws-sdk-go` version used to v1.38.14.

## Special thanks

Special thanks to @sylr for their contribution!

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/118
<br>

# v0.6.11
_Released Mar 30, 2021_
## Description

This is an internal code cleanup release, where some functions were moved to a common library. There is no change in behavior to `kubergrunt` in this release.

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/115
<br>

# v0.6.10
_Released Feb 23, 2021_
## Commands Affected

- `eks sync-core-components` ([docs](https://github.com/gruntwork-io/kubergrunt#sync-core-components))

## Description
`eks sync-core-components` now supports k8s 1.19.

## Special thanks

Special thanks to @FriedCircuits for their contribution!

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/112
<br>

# v0.6.1
_Released Sep 29, 2020_
## Commands Affected

- `eks sync-core-components` [**NEW**]

## Description

This introduces `kubergrunt eks sync-core-components`, which can be used to roll out updates to the core components of an EKS cluster after upgrading the Kubernetes version. Refer to [the AWS docs on the topic](https://docs.aws.amazon.com/eks/latest/userguide/update-cluster.html) for more details.

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/96
<br>

# v0.6.0
_Released Sep 21, 2020_
## Commands Affected

- `helm` [**REMOVED**]

## Description

The `helm` subcommand was removed in this release, now that [Helm v2 is scheduled for end of life](https://helm.sh/blog/helm-v2-deprecation-timeline/).

The compilation and build steps have also been switched to go modules instead of `dep`.

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/95
<br>

# v0.5.9
_Released Feb 26, 2020_
## Commands Affected

- `eks deploy`

## Description

This adds a new CLI arg `--delete-local-data` so that the rolling deployment can continue even if there are pods using emptyDir.

## Special Thanks

Special thanks to @adamrbennett for their contribution!

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/81
<br>

# v0.5.8
_Released Oct 22, 2019_
## Commands Affected

- `k8s kubectl` [**NEW**]

## Description

This introduces a new subcommand `kubergrunt k8s kubectl`, which is a thin wrapper around `kubectl`. This allows you to use the `kubergrunt` CLI args for authenticating to the Kubernetes cluster. This is most useful when you want to pass in the server certificate authority data as a base64 encoded arg on the CLI, which you currently can not do with `kubectl` directly.

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/74
<br>

# v0.5.7
_Released Oct 9, 2019_
## Commands Affected

- `helm grant`
- `helm configure`

## Description

The `helm grant` command will now additionally grant permissions to get the `Deployment` resource that corresponds to the tiller deployment. This is necessary to use the terraform helm provider.

The `--helm-home` option of the `helm configure` command can now be set using the environment variable, `HELM_HOME`.

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/71
- https://github.com/gruntwork-io/kubergrunt/pull/73
<br>

# v0.5.6
_Released Oct 7, 2019_
## Commands Affected

- `tls gen`

## Description

The `tls gen` command now supports setting the DNS names in the Subject Alternative Name (SAN) of the certificate. You can configure this using the new `--tls-dns-name` arg for the command.

## Special thanks

Special thanks to @KyleMartin901 for their contribution!

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/70
<br>

# v0.5.5
_Released Oct 3, 2019_
#68 Update to AWS SDK 1.25.4 (latest version) so kubergrunt can be used in an EKS container that is assuming an IAM role
<br>

# v0.5.4
_Released Sep 17, 2019_
## Commands Affected

- `eks deploy`

## Description

The `eks deploy` command will now cordon all the old nodes before starting to drain them so that the evicted pods get scheduled on the new nodes in the cluster.

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/65
<br>

# v0.5.3
_Released Sep 13, 2019_
## Commands Affected

- `eks oidc-thumbprint` [**NEW**]

## Description

This release introduces a new command `kubergrunt eks oidc-thumbprint` which can be used to compute the root CA thumbprint for a given OpenID Connect Issuer. This thumbprint is necessary when setting up an OpenID Connect Provider for use with the IAM Roles for Service Accounts feature provided by EKS.

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/62
<br>

# v0.5.2
_Released Aug 15, 2019_
## Commands Affected

Any command that auths to Kubernetes API.

## Description

You can now specify the kubernetes config path to use using the environment variable `KUBECONFIG`.

## Special Thanks

Special thanks to @KyleMartin901 for their contribution!

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/59
<br>

# v0.5.14
_Released Jan 8, 2021_
## Description

This is a backport fix to ensure that the helm 2 functionality works with the new location for the stable repository. See [the helm announcement on the new location for stable repository](https://helm.sh/blog/new-location-stable-incubator-charts/) for more information.

NOTE: This release is marked as a pre-release to prevent it from bubbling as the latest release on the repository.
<br>

# v0.5.13
_Released Apr 14, 2020_
## Commands Affected

None

## Description

Starting this release, we will publish sha256 checksums with our binaries.

## Special thanks

Special thanks to @DrFaust92 for their contribution!

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/89
<br>

# v0.5.12
_Released Mar 27, 2020_
## Commands Affected

All commands that take in kubernetes auth options.

## Description

This release introduces a new authentication option for kubernetes API calls to EKS, where it will create a temporary kubeconfig file using the AWS SDK.

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/87
<br>

# v0.5.11
_Released Mar 10, 2020_
## Commands Affected

- `eks deploy`

## Description

This fixes a bug in the `eks deploy` command where it did not handle `LoadBalancer` Services that are internal.

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/83
<br>

# v0.5.10
_Released Mar 3, 2020_
## Commands Affected

- `eks oidc`

## Description

This fixes a nondeterministic failure in retrieving the thumbprint, where depending on timing, you are likely to fail retrieval of the OIDC thumbprint when the EKS cluster is first launching.

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/82
<br>

# v0.5.1
_Released Jul 13, 2019_
## Commands Affected

Any command that requires TLS subject info.

## Description

The TLS subject info JSON previously only allowed one name for some of the fields (e.g `org` for organization), but it now allows using alternate names:

- Common Name: `common_name`
- Organization: `org` OR `organization`
- Organizational Unit: `org_unit` OR `organizational_unit`
- City: `city` OR `locality`
- State: `state` OR `province`
- Country: `country`

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/57
<br>

# v0.5.0
_Released Jun 3, 2019_
## Commands Affected

- `helm grant`
- `helm revoke`

## Description

This release makes the labels used to tag the Kubernetes resources more robust so that they fit within the constraints of the label system. For example, this will modify special characters before including them into as labels.

**WARNING**: This revoke command is incompatible with any grant commands that were run before this version. See the "Migration Notes" for details on how to revoke access to already granted entities.

## Migration Notes

The revoke command depends on labels that are set on the `Role`, `RoleBinding`, and `Secret` to be able to find them for deletion. The new labels are incompatible with the previous versions, preventing you from revoking entities that were granted prior to this version.

However, entities granted access on `v0.4.0` can be revoked using the `kubergrunt helm revoke` command from `v0.4.0`.

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/55
<br>

# v0.4.0
_Released May 28, 2019_
## Commands Affected

- `helm grant`
- `helm revoke` [**NEW**]

## Description

This release introduces the `helm revoke` command, which will remove access to Tiller from the specified RBAC entities. This is done by removing:

- The `RoleBinding` that binds permissions to access the Kubernetes `Secret` containing the TLS key pair and the Tiller Pod.
- The `Role` that grants permissions to access the Kubernetes `Secret` containing the TLS key pair and the Tiller Pod.
- The Kubernetes `Secret` containing the client TLS certificate.

To support the revoke action, this release updates the `grant` command to add labels to the `Role`, `RoleBinding`, and `Secret` that is generated.

**WARNING**: This revoke command is incompatible with any grant commands that were run before this version. See the "Migration Notes" for details on how to revoke access to already granted entities.

**NOTE**: This revoke command does not mark the TLS certificate as unusable. This means that if the entity is able to open a connection to the Tiller pod, they will be able to reuse the certificate. For a full revocation, you will need to reissue the certificates under a new CA that is installed in the Tiller pod.

## Migration Notes

The revoke command depends on labels that are set on the `Role`, `RoleBinding`, and `Secret` to be able to find them for deletion. Previous versions of `kubergrunt` did not set labels on these resources during the `grant` operation, making them incompatible with the implementation of `revoke`. If you wish to `revoke` access to entities that were granted access using a previous version of `grant`, follow these steps:

- Get the entity name. For RBAC Users and Groups, this is the name of the User or Group. For ServiceAccounts, this will be the string `NAMESPACE/NAME`, where `NAMESPACE` is the namespace where the ServiceAccount is defined, and `NAME` is the name of the ServiceAccount.
- Take the entity name and generate a md5 hash of the string. We will refer to this as the environment variable `ENTITY_ID_MD5`.
- Record the name of the tiller namespace in the environment variable `TILLER_NAMESPACE`.
- Delete the objects using `kubectl`:

```
kubectl delete rolebinding "$ENTITY_ID_MD5-$ENTITY_ID_MD5-$TILLER_NAMESPACE-tiller-access-binding" -n "$TILLER_NAMESPACE"
kubectl delete role "$ENTITY_ID_MD5-$TILLER_NAMESPACE-tiller-access" -n "$TILLER_NAMESPACE"
kubectl delete secret "tiller-client-$ENTITY_ID_MD5-certs" -n "$TILLER_NAMESPACE"
```

## Special Thanks

Special thanks to @bwhaley for their contribution!

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/50
- https://github.com/gruntwork-io/kubergrunt/pull/51
<br>

# v0.3.9
_Released May 6, 2019_
## Commands Affected

- `helm configure`
- Any commands that accept TLS configurations

## Description

This release updates `kubergrunt helm configure` with a new option `--as-tf-data`, which enables you to call it in an external data source. Passing this flag will cause the command to output the configured helm home directory in the output json on `stdout` at the end of the command.

Additionally, `kubergrunt helm configure` now supports a new option you can set for `--helm-home` (E.g `--helm-home __TMP__`) to inform kubergrunt that you intend to use a tmp directory for the configured helm home directory. Passing this option in will cause the command to generate a new directory in the tmp folder for the platform to use as the helm home.

The above two features, when used in combination, allows you to dynamically configure the `helm` provider on the host machine in Terraform.

Finally, this release also updates the TLS subject json parsing for the CLI such that it can accept equivalent arg names in the terraform TLS subject block: https://www.terraform.io/docs/providers/tls/r/cert_request.html#common_name

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/47
- https://github.com/gruntwork-io/kubergrunt/pull/46
<br>

# v0.3.8
_Released May 1, 2019_
## Commands Affected

- `tls gen`
- `helm wait-for-tiller` [**NEW**]

## Description

This release updates the `tls gen` command to use the new way of authenticating to Kubernetes (specifically passing in server and token info directly) and using JSON to configure the TLS subject.

This release also introduces a new command `helm wait-for-tiller` which can be used to wait for a tiller deployment to roll out Pods, and have at least one Pod that can be pinged. This enables chaining calls to helm after helm is deployed when using a different helm deployment process that doesn't rely on the helm client (e.g creating deployment resources manually).

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/44
<br>

# v0.3.7
_Released Apr 26, 2019_
## Commands Affected

- `k8s wait-for-ingress` [**NEW**]

## Description

This release introduces the `k8s wait-for-ingress` sub command which can be used to wait until an `Ingress` resource has an endpoint associated with it.


## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/43
<br>

# v0.3.6
_Released Apr 15, 2019_
## Description

All commands that setup TLS certificates (e.g `helm deploy` and `helm grant`) can now take in the TLS subject information as a json.

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/40
<br>

# v0.3.5
_Released Apr 10, 2019_
## Description

All commands that need to authenticate to Kubernetes (e.g `eks deploy`, `helm deploy`, etc) now take in direct authentication parameters as an alternative means to auth. You can now pass in the authentication parameters `--kubectl-server-endpoint`, `--kubectl-certificate-authority` and `--kubectl-token` to the command to setup authentication to the Kubernetes cluster, as opposed to setting up the config.

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/38
- https://github.com/gruntwork-io/kubergrunt/pull/39
<br>

# v0.3.4
_Released Apr 8, 2019_
## Commands Affected

- `eks token`

## Description

The `eks token` command now accepts a new parameter `--as-tf-data`, which will encode the token output in a format that can be used as an [external data source](https://www.terraform.io/docs/providers/external/data_source.html) in terraform. This allows you to configure the kubernetes and helm providers without configuring a kubectl context.

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/37
<br>

# v0.3.3
_Released Apr 2, 2019_
## Commands Affected

- `tls gen` [**NEW**]

## Description

This release introduces the `tls gen` sub command which can be used to generate TLS certificate key pairs that are managed using Kubernetes Secrets.


## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/34
- https://github.com/gruntwork-io/kubergrunt/pull/35
<br>

# v0.3.2
_Released Feb 13, 2019_
## Commands Affected

- `helm deploy`

## Description

This release fixes the following bugs:

- `helm deploy` was not able to successfully authenticate against a GKE cluster on Google Cloud Platform.

Additional improvements:

- Updated the `Gopkg.lock` file.

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/25
- https://github.com/gruntwork-io/kubergrunt/pull/29
- https://github.com/gruntwork-io/kubergrunt/pull/32
<br>

# v0.3.1
_Released Feb 8, 2019_
## Commands Affected

- `helm deploy`
- `helm configure`

## Description

This release introduces the following new feature:

- You can now specify the Tiller container image to use with `helm deploy`, using the `--tiller-image` and `--tiller-version` flags.

This release fixes the following bugs:

- `helm deploy` had a bug where the client was configured after Tiller was confirmed to be deployed. However, the Tiller deployment check requires connecting to Tiller, leading to a chicken and egg problem. This fixes that by delaying the polling step until after the client has been granted access and configured.
- `helm configure` did not setup the helm home directory with the initial set of repository, leading the operator to run the repo update command immediately after configuring. This fixes that by installing the stable repo as part of configure.
- `helm deploy` required finding the executable `kubergrunt` in the `PATH`, exactly as is. This removes that need.

Additional improvements:

- `helm deploy` now uses the helm go library instead of calling out to the helm client.

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/21
- https://github.com/gruntwork-io/kubergrunt/pull/22
- https://github.com/gruntwork-io/kubergrunt/pull/23
<br>

# v0.3.0
_Released Feb 2, 2019_
## Commands Affected

- `helm deploy` [**BREAKING CHANGE**]

## Description

This release fixes two bugs:

- `helm deploy` had a bug where the client certificates were generated using the same parameters as the server certificates. Fixing this requires new required parameters to the command: `--client-tls-common-name` and `--client-tls-org`.
- `helm deploy` had a bug where it ignored the `--helm-home` directory when deploying Tiller. This prevented a custom home directory from being initialized.

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/16
- https://github.com/gruntwork-io/kubergrunt/pull/17
<br>

# v0.2.0
_Released Feb 1, 2019_
## Commands Affected

- `helm deploy` [**BREAKING CHANGE**]

## Description

This release updates the `helm deploy` command to now require an RBAC entity to be passed in (one of `--rbac-user`, `--rbac-group`, or `--rbac-service-account`), which will be used to configure the local helm client.

This ensures that you can interact with `helm` immediately after a `deploy`, without the need to `grant` and `configure`.

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/14
<br>

# v0.12.1-test-signing-binaries
_Released Aug 11, 2023_
Testing changes in https://github.com/gruntwork-io/kubergrunt/pull/209
<br>

# v0.12.1
_Released Aug 11, 2023_
## What's Changed
* Update CircleCI config to sign MacOS binaries by @marinalimeira in https://github.com/gruntwork-io/kubergrunt/pull/209

## New Contributors
* @marinalimeira made their first contribution in https://github.com/gruntwork-io/kubergrunt/pull/209

**Full Changelog**: https://github.com/gruntwork-io/kubergrunt/compare/v0.12.0...v0.12.1
<br>

# v0.12.0-alpha.1
_Released Jul 7, 2023_
## What's Changed
* [skip ci] Remove Zack from CODEOWNERS by @zackproser in https://github.com/gruntwork-io/kubergrunt/pull/199
* Issue/88 by @pras111gw in https://github.com/gruntwork-io/kubergrunt/pull/200
Added k8s 1.27 and removed 1.22 (deprecated)

**Full Changelog**: https://github.com/gruntwork-io/kubergrunt/compare/v0.11.3...v0.12.0-alpha.1
<br>

# v0.12.0
_Released Jul 20, 2023_
## What's Changed
* [skip ci] Remove Zack from CODEOWNERS by @zackproser in https://github.com/gruntwork-io/kubergrunt/pull/199
* Issue/88 by @pras111gw in https://github.com/gruntwork-io/kubergrunt/pull/200
Added k8s 1.27 and removed 1.22 (deprecated)

**Full Changelog**: https://github.com/gruntwork-io/kubergrunt/compare/v0.11.3...v0.12.0
<br>

# v0.11.3-alpha.1
_Released May 6, 2023_
[feature/k8s-126 (https://github.com/gruntwork-io/kubergrunt/pull/197)](https://github.com/gruntwork-io/kubergrunt/pull/197)

**Full Changelog**: https://github.com/gruntwork-io/kubergrunt/compare/v0.11.2...v0.11.3-alpha.1
<br>

# v0.11.3 Support kubernetes 1.26 in EKS
_Released Jun 2, 2023_
## What's Changed
* Feature/k8s 126 by @pras111gw in https://github.com/gruntwork-io/kubergrunt/pull/197
* Chore/update187 eksbuild5 by @pras111gw in https://github.com/gruntwork-io/kubergrunt/pull/198

Add support for EKS/Kubernetes version 1.26

**Full Changelog**: https://github.com/gruntwork-io/kubergrunt/compare/v0.11.2...v0.11.3
<br>

# v0.11.2
_Released Apr 28, 2023_
## What's Changed
* Bump github.com/prometheus/client_golang from 1.11.0 to 1.11.1 by @dependabot in https://github.com/gruntwork-io/kubergrunt/pull/183
* Bump golang.org/x/net from 0.5.0 to 0.7.0 by @dependabot in https://github.com/gruntwork-io/kubergrunt/pull/188
* Feature/k8s 125 hotfix by @pras111gw in https://github.com/gruntwork-io/kubergrunt/pull/187
* [skip ci] Refactor contexts by @eak12913 in https://github.com/gruntwork-io/kubergrunt/pull/190
* test fix - new version of kubeproxy by @pras111gw in https://github.com/gruntwork-io/kubergrunt/pull/191
* Change deprecated delete-local-data flag to delete-emptydir-data by @autero1 in https://github.com/gruntwork-io/kubergrunt/pull/161
* Update CODEOWNERS by @yorinasub17 in https://github.com/gruntwork-io/kubergrunt/pull/175
* Feature/issue 194 by @pras111gw in https://github.com/gruntwork-io/kubergrunt/pull/195

**Full Changelog**: https://github.com/gruntwork-io/kubergrunt/compare/v0.11.1...v0.11.2
<br>

# v0.11.1
_Released Mar 28, 2023_
## What's Changed
* [skip ci] Added FUNDING.yml by @eak12913 in https://github.com/gruntwork-io/kubergrunt/pull/180
* feature/k8s 125 by @pras111gw in https://github.com/gruntwork-io/kubergrunt/pull/185

Add support for Kubernetes version 1.25 and drop support for deprecated 1.21.
EKS no longer supports versions below 1.22 - https://docs.aws.amazon.com/eks/latest/userguide/kubernetes-versions.html#available-versions

**Full Changelog**: https://github.com/gruntwork-io/kubergrunt/compare/v0.10.1...v0.11.1

## What's Changed
* [skip ci] Added FUNDING.yml by @eak12913 in https://github.com/gruntwork-io/kubergrunt/pull/180
* feature/k8s 125 by @pras111gw in https://github.com/gruntwork-io/kubergrunt/pull/185


**Full Changelog**: https://github.com/gruntwork-io/kubergrunt/compare/v0.10.1...v0.11.1
<br>

# v0.10.2
_Released Jan 18, 2023_
## Description

Updated multiple dependencies to incorporate security patches.

- https://security.snyk.io/vuln/SNYK-GOLANG-GOPKGINYAMLV3-2952714
- https://security.snyk.io/vuln/SNYK-GOLANG-GOPKGINYAMLV3-2841557

## Related links

#179 
<br>

# v0.10.0-alpha.3
_Released Dec 1, 2022_
Alpha release for regression testing kubergrunt support for 1.24 in terraform-aws-eks.
Fixed VPC CNI manifest download.
<br>

# v0.10.0-alpha.2
_Released Nov 16, 2022_
Alpha release for regression testing kubergrunt support for 1.24 in terraform-aws-eks.
(NOTE: alpha.1 release was mistakenly created from `main`)
<br>

# v0.10.0-alpha.1
_Released Sep 28, 2022_
Alpha release for regression testing `kubergrunt` support for `1.24` in `terraform-aws-eks`.
<br>

# v0.10.0
_Released Dec 15, 2022_
## Commands

- `eks sync-core-components`

## Description

Add support for Kubernetes version 1.24 and drop support for 1.20 in the `sync-core-components` command.

## Related links

#177 
<br>

# v0.1.5
_Released Jan 31, 2019_
## Commands Affected

- `helm deploy` [**NEW**]
- `helm undeploy` [**NEW**]
- `helm configure` [**NEW**]
- `helm grant` [**NEW**]

## Description

This release introduces a suite of subcommands that can be used to deploy helm with all the official security best practices considered. Take a look at the command help text for more information.

## Reference

- https://github.com/gruntwork-io/kubergrunt/pull/1
<br>

# v0.1.4
_Released Jan 14, 2019_
## Commands Affected

- `eks configure`
- `eks token`
- `eks deploy`

## Description

This is a compatible release of `kubergrunt` with [`package-k8s` v0.1.4](https://github.com/gruntwork-io/package-k8s/releases/tag/v0.1.4).
<br>

# Alpha release for regression testing kubergrunt support for 1.25 in terraform-aws-eks.
_Released Mar 7, 2023_
## What's Changed
* Upgrade Go, dependencies and k8s version by @autero1 in https://github.com/gruntwork-io/kubergrunt/pull/179
* [skip ci] Added FUNDING.yml by @eak12913 in https://github.com/gruntwork-io/kubergrunt/pull/180
* feature/k8s 125 by @pras111gw in https://github.com/gruntwork-io/kubergrunt/pull/185


**Full Changelog**: https://github.com/gruntwork-io/kubergrunt/compare/v0.10.0...v0.11.0-alpha.1
<br>

