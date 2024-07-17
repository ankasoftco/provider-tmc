package repository_credential

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tanzu-mission-control_repository_credential", func(r *config.Resource) {
		r.ShortGroup = "tmc"
		r.Kind = "RepositoryCredential"
		r.Version = "v1alpha1"
	})
}