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
	cluster_group "github.com/ankasoftco/provider-tmc/internal/controller/tmc/cluster_group"
	cluster_node_pool "github.com/ankasoftco/provider-tmc/internal/controller/tmc/cluster_node_pool"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.Setup,
		akscluster.Setup,
		cluster.Setup,
		cluster_group.Setup,
		cluster_node_pool.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
