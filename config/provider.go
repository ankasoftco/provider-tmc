/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	akscluster "github.com/ankasoftco/provider-tmc/config/akscluster"
	cluster "github.com/ankasoftco/provider-tmc/config/cluster"
	cluster_group "github.com/ankasoftco/provider-tmc/config/cluster_group"
	cluster_node_pool "github.com/ankasoftco/provider-tmc/config/cluster_node_pool"
	credential "github.com/ankasoftco/provider-tmc/config/credential"
	custom_policy "github.com/ankasoftco/provider-tmc/config/custom_policy"
	ekscluster "github.com/ankasoftco/provider-tmc/config/ekscluster"
	git_repository "github.com/ankasoftco/provider-tmc/config/git_repository"
	iam_policy "github.com/ankasoftco/provider-tmc/config/iam_policy"
	image_policy "github.com/ankasoftco/provider-tmc/config/image_policy"
	integration "github.com/ankasoftco/provider-tmc/config/integration"
	kubernetes_secret "github.com/ankasoftco/provider-tmc/config/kubernetes_secret"
	kustomization "github.com/ankasoftco/provider-tmc/config/kustomization"
	namespace "github.com/ankasoftco/provider-tmc/config/namespace"


)

const (
	resourcePrefix = "tmc"
	modulePath     = "github.com/ankasoftco/provider-tmc"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		akscluster.Configure,
		cluster.Configure,
		cluster_group.Configure,
		cluster_node_pool.Configure,
		credential.Configure,
		custom_policy.Configure,
		ekscluster.Configure,
		git_repository.Configure,
		iam_policy.Configure,
		image_policy.Configure,
		integration.Configure,
		kubernetes_secret.Configure,
		kustomization.Configure,
		name
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
