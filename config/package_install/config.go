package package_install

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tanzu-mission-control_package_install", func(r *config.Resource) {
		r.ShortGroup = "tmc"
		r.Kind = "Package_Install"
		r.Version = "v1alpha1"
	})
}