package package_install

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tanzu-mission-control_package_install", func(r *config.Resource) {
		r.ShortGroup = "tmc"
		r.Kind = "PackageInstall"
		r.Version = "v1alpha1"
		r.References["namespace"] = config.Reference{
			Type: "github.com/ankasoftco/provider-tmc/apis/tmc/v1alpha1.Namespace",
		}
	})
}