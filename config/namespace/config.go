package namespace

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tanzu-mission-control_namespace", func(r *config.Resource) {
		r.ShortGroup = "tmc"
		r.Kind = "Namespace"
		r.Version = "v1alpha1"
	})
}
