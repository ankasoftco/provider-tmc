package custom_policy

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("tanzu-mission-control_custom_policy", func(r *config.Resource) {
		r.ShortGroup = "tmc"
		r.Kind = "Custom_Policy"
		r.Version = "v1alpha1"
    })
}