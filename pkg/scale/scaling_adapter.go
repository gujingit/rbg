package scale

import (
	"fmt"

	workloadsv1alpha "sigs.k8s.io/rbgs/api/workloads/v1alpha1"
)

func GenerateScalingAdapterName(rbgName, roleName string) string {
	return fmt.Sprintf("%s-%s", rbgName, roleName)
}

func IsScalingAdapterManagedByRBG(
	scalingAdapter *workloadsv1alpha.RoleBasedGroupScalingAdapter,
	rbg *workloadsv1alpha.RoleBasedGroup,
) bool {
	if scalingAdapter == nil {
		return false
	}

	// If rbg is nil, it could be that the rbg has not been created yet. Just checking if the scalingAdapter's
	// OwnerReference is set to the rolebasedgroup controller.
	// If the rbg already exists, then check if the uid in the ownerReference is equal.
	return scalingAdapter.ContainsRBGOwner(rbg)
}

func IsScalingAdapterEnable(roleSpec *workloadsv1alpha.RoleSpec) bool {
	if roleSpec == nil || roleSpec.ScalingAdapter == nil {
		return false
	}
	return roleSpec.ScalingAdapter.Enable
}
