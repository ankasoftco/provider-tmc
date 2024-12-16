package package_repository

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tanzu-mission-control_package_repository", func(r *config.Resource) {
		r.ShortGroup = "tmc"
		r.Kind = "PackageRepository"
		r.Version = "v1alpha1"
	})
}
