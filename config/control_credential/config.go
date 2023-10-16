package control_credential

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("tanzu-mission-control_credential", func(r *config.Resource) {
		r.ShortGroup = "tmc"
		r.Kind = "Control_Credential"
		r.Version = "v1alpha1"
    })
}