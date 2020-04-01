module github.com/devigned/go-autorest-mods-test

go 1.14

require (
	github.com/docker/distribution v2.7.1+incompatible
	sigs.k8s.io/cluster-api-provider-azure v0.4.0
)

replace github.com/docker/distribution => github.com/devigned/distribution v0.0.0-20200401145234-a46fa76a3251
