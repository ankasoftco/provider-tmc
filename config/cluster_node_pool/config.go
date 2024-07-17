package cluster_node_pool

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("tanzu-mission-control_cluster_node_pool", func(r *config.Resource) {
		r.ShortGroup = "tmc"
		r.Kind = "ClusterNodePool"
		r.Version = "v1alpha1"
    })
}