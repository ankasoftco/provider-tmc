/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this PackageInstall.
func (mg *PackageInstall) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Namespace),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NamespaceRef,
		Selector:     mg.Spec.ForProvider.NamespaceSelector,
		To: reference.To{
			List:    &NamespaceList{},
			Managed: &Namespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Namespace")
	}
	mg.Spec.ForProvider.Namespace = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NamespaceRef = rsp.ResolvedReference

	return nil
}
