// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/openshift/api/operator/v1"
	operatorv1 "github.com/openshift/client-go/operator/applyconfigurations/operator/v1"
	typedoperatorv1 "github.com/openshift/client-go/operator/clientset/versioned/typed/operator/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeClusterCSIDrivers implements ClusterCSIDriverInterface
type fakeClusterCSIDrivers struct {
	*gentype.FakeClientWithListAndApply[*v1.ClusterCSIDriver, *v1.ClusterCSIDriverList, *operatorv1.ClusterCSIDriverApplyConfiguration]
	Fake *FakeOperatorV1
}

func newFakeClusterCSIDrivers(fake *FakeOperatorV1) typedoperatorv1.ClusterCSIDriverInterface {
	return &fakeClusterCSIDrivers{
		gentype.NewFakeClientWithListAndApply[*v1.ClusterCSIDriver, *v1.ClusterCSIDriverList, *operatorv1.ClusterCSIDriverApplyConfiguration](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("clustercsidrivers"),
			v1.SchemeGroupVersion.WithKind("ClusterCSIDriver"),
			func() *v1.ClusterCSIDriver { return &v1.ClusterCSIDriver{} },
			func() *v1.ClusterCSIDriverList { return &v1.ClusterCSIDriverList{} },
			func(dst, src *v1.ClusterCSIDriverList) { dst.ListMeta = src.ListMeta },
			func(list *v1.ClusterCSIDriverList) []*v1.ClusterCSIDriver { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.ClusterCSIDriverList, items []*v1.ClusterCSIDriver) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
