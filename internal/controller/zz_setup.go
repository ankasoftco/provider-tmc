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
	credential "github.com/ankasoftco/provider-tmc/internal/controller/tmc/credential"
	custom_policy "github.com/ankasoftco/provider-tmc/internal/controller/tmc/custom_policy"
	ekscluster "github.com/ankasoftco/provider-tmc/internal/controller/tmc/ekscluster"
	git_repository "github.com/ankasoftco/provider-tmc/internal/controller/tmc/git_repository"
	iam_policy "github.com/ankasoftco/provider-tmc/internal/controller/tmc/iam_policy"
	image_policy "github.com/ankasoftco/provider-tmc/internal/controller/tmc/image_policy"
	integration "github.com/ankasoftco/provider-tmc/internal/controller/tmc/integration"
	kubernetes_secret "github.com/ankasoftco/provider-tmc/internal/controller/tmc/kubernetes_secret"
	kustomization "github.com/ankasoftco/provider-tmc/internal/controller/tmc/kustomization"
	namespace "github.com/ankasoftco/provider-tmc/internal/controller/tmc/namespace"
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
		credential.Setup,
		custom_policy.Setup,
		ekscluster.Setup,
		git_repository.Setup,
		iam_policy.Setup,
		image_policy.Setup,
		integration.Setup,
		kubernetes_secret.Setup,
		kustomization.Setup,
		namespace.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
