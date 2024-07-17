/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	providerconfig "github.com/ankasoftco/provider-tmc/internal/controller/providerconfig"
	akscluster "github.com/ankasoftco/provider-tmc/internal/controller/tmc/akscluster"
	cluster "github.com/ankasoftco/provider-tmc/internal/controller/tmc/cluster"
	clustergroup "github.com/ankasoftco/provider-tmc/internal/controller/tmc/clustergroup"
	clusternodepool "github.com/ankasoftco/provider-tmc/internal/controller/tmc/clusternodepool"
	credential "github.com/ankasoftco/provider-tmc/internal/controller/tmc/credential"
	custompolicy "github.com/ankasoftco/provider-tmc/internal/controller/tmc/custompolicy"
	ekscluster "github.com/ankasoftco/provider-tmc/internal/controller/tmc/ekscluster"
	gitrepository "github.com/ankasoftco/provider-tmc/internal/controller/tmc/gitrepository"
	iampolicy "github.com/ankasoftco/provider-tmc/internal/controller/tmc/iampolicy"
	imagepolicy "github.com/ankasoftco/provider-tmc/internal/controller/tmc/imagepolicy"
	integration "github.com/ankasoftco/provider-tmc/internal/controller/tmc/integration"
	kubernetessecret "github.com/ankasoftco/provider-tmc/internal/controller/tmc/kubernetessecret"
	kustomization "github.com/ankasoftco/provider-tmc/internal/controller/tmc/kustomization"
	namespace "github.com/ankasoftco/provider-tmc/internal/controller/tmc/namespace"
	namespacequotapolicy "github.com/ankasoftco/provider-tmc/internal/controller/tmc/namespacequotapolicy"
	networkpolicy "github.com/ankasoftco/provider-tmc/internal/controller/tmc/networkpolicy"
	packageinstall "github.com/ankasoftco/provider-tmc/internal/controller/tmc/packageinstall"
	packagerepository "github.com/ankasoftco/provider-tmc/internal/controller/tmc/packagerepository"
	repositorycredential "github.com/ankasoftco/provider-tmc/internal/controller/tmc/repositorycredential"
	securitypolicy "github.com/ankasoftco/provider-tmc/internal/controller/tmc/securitypolicy"
	workspace "github.com/ankasoftco/provider-tmc/internal/controller/tmc/workspace"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.Setup,
		akscluster.Setup,
		cluster.Setup,
		clustergroup.Setup,
		clusternodepool.Setup,
		credential.Setup,
		custompolicy.Setup,
		ekscluster.Setup,
		gitrepository.Setup,
		iampolicy.Setup,
		imagepolicy.Setup,
		integration.Setup,
		kubernetessecret.Setup,
		kustomization.Setup,
		namespace.Setup,
		namespacequotapolicy.Setup,
		networkpolicy.Setup,
		packageinstall.Setup,
		packagerepository.Setup,
		repositorycredential.Setup,
		securitypolicy.Setup,
		workspace.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
