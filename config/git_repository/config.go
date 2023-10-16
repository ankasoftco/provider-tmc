
package git_repository

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("tanzu-mission-control_git_repository", func(r *config.Resource) {
		r.ShortGroup = "tmc"
		r.Kind = "Git_Repository"
		r.Version = "v1alpha1"
    })
}