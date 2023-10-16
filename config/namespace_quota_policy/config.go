package namespace_quota_policy

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tanzu-mission-control_namespace_quota_policy", func(r *config.Resource) {
		r.ShortGroup = "tmc"
		r.Kind = "Namespace_Quota_Policy"
		r.Version = "v1alpha1"
	})
}