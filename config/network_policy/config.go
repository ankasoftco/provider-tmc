package network_policy

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tanzu-mission-control_network_policy", func(r *config.Resource) {
		r.ShortGroup = "tmc"
		r.Kind = "NetworkPolicy"
		r.Version = "v1alpha1"
	})
}
