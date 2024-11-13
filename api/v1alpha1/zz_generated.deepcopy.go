//go:build !ignore_autogenerated

/*
Copyright 2024, Tax Administration of The Netherlands.
Licensed under the EUPL 1.2.
See LICENSE.md for details.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/belastingdienst/opr-paas/internal/quota"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigArgoPermissions) DeepCopyInto(out *ConfigArgoPermissions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigArgoPermissions.
func (in *ConfigArgoPermissions) DeepCopy() *ConfigArgoPermissions {
	if in == nil {
		return nil
	}
	out := new(ConfigArgoPermissions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ConfigCapPerm) DeepCopyInto(out *ConfigCapPerm) {
	{
		in := &in
		*out = make(ConfigCapPerm, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigCapPerm.
func (in ConfigCapPerm) DeepCopy() ConfigCapPerm {
	if in == nil {
		return nil
	}
	out := new(ConfigCapPerm)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ConfigCapabilities) DeepCopyInto(out *ConfigCapabilities) {
	{
		in := &in
		*out = make(ConfigCapabilities, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigCapabilities.
func (in ConfigCapabilities) DeepCopy() ConfigCapabilities {
	if in == nil {
		return nil
	}
	out := new(ConfigCapabilities)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigCapability) DeepCopyInto(out *ConfigCapability) {
	*out = *in
	in.QuotaSettings.DeepCopyInto(&out.QuotaSettings)
	if in.ExtraPermissions != nil {
		in, out := &in.ExtraPermissions, &out.ExtraPermissions
		*out = make(ConfigCapPerm, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.DefaultPermissions != nil {
		in, out := &in.DefaultPermissions, &out.DefaultPermissions
		*out = make(ConfigCapPerm, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigCapability.
func (in *ConfigCapability) DeepCopy() *ConfigCapability {
	if in == nil {
		return nil
	}
	out := new(ConfigCapability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigLdap) DeepCopyInto(out *ConfigLdap) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigLdap.
func (in *ConfigLdap) DeepCopy() *ConfigLdap {
	if in == nil {
		return nil
	}
	out := new(ConfigLdap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigQuotaSettings) DeepCopyInto(out *ConfigQuotaSettings) {
	*out = *in
	if in.DefQuota != nil {
		in, out := &in.DefQuota, &out.DefQuota
		*out = make(map[v1.ResourceName]resource.Quantity, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.MinQuotas != nil {
		in, out := &in.MinQuotas, &out.MinQuotas
		*out = make(map[v1.ResourceName]resource.Quantity, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.MaxQuotas != nil {
		in, out := &in.MaxQuotas, &out.MaxQuotas
		*out = make(map[v1.ResourceName]resource.Quantity, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigQuotaSettings.
func (in *ConfigQuotaSettings) DeepCopy() *ConfigQuotaSettings {
	if in == nil {
		return nil
	}
	out := new(ConfigQuotaSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ConfigRoleMappings) DeepCopyInto(out *ConfigRoleMappings) {
	{
		in := &in
		*out = make(ConfigRoleMappings, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigRoleMappings.
func (in ConfigRoleMappings) DeepCopy() ConfigRoleMappings {
	if in == nil {
		return nil
	}
	out := new(ConfigRoleMappings)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ConfigRolesSas) DeepCopyInto(out *ConfigRolesSas) {
	{
		in := &in
		*out = make(ConfigRolesSas, len(*in))
		for key, val := range *in {
			var outVal map[string]bool
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make(map[string]bool, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigRolesSas.
func (in ConfigRolesSas) DeepCopy() ConfigRolesSas {
	if in == nil {
		return nil
	}
	out := new(ConfigRolesSas)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedName) DeepCopyInto(out *NamespacedName) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedName.
func (in *NamespacedName) DeepCopy() *NamespacedName {
	if in == nil {
		return nil
	}
	out := new(NamespacedName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Paas) DeepCopyInto(out *Paas) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Paas.
func (in *Paas) DeepCopy() *Paas {
	if in == nil {
		return nil
	}
	out := new(Paas)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Paas) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PaasCapabilities) DeepCopyInto(out *PaasCapabilities) {
	{
		in := &in
		*out = make(PaasCapabilities, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PaasCapabilities.
func (in PaasCapabilities) DeepCopy() PaasCapabilities {
	if in == nil {
		return nil
	}
	out := new(PaasCapabilities)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PaasCapability) DeepCopyInto(out *PaasCapability) {
	*out = *in
	if in.Quota != nil {
		in, out := &in.Quota, &out.Quota
		*out = make(quota.Quota, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.SshSecrets != nil {
		in, out := &in.SshSecrets, &out.SshSecrets
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PaasCapability.
func (in *PaasCapability) DeepCopy() *PaasCapability {
	if in == nil {
		return nil
	}
	out := new(PaasCapability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PaasConfig) DeepCopyInto(out *PaasConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PaasConfig.
func (in *PaasConfig) DeepCopy() *PaasConfig {
	if in == nil {
		return nil
	}
	out := new(PaasConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PaasConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PaasConfigList) DeepCopyInto(out *PaasConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PaasConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PaasConfigList.
func (in *PaasConfigList) DeepCopy() *PaasConfigList {
	if in == nil {
		return nil
	}
	out := new(PaasConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PaasConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PaasConfigSpec) DeepCopyInto(out *PaasConfigSpec) {
	*out = *in
	if in.DecryptKeyPaths != nil {
		in, out := &in.DecryptKeyPaths, &out.DecryptKeyPaths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Capabilities != nil {
		in, out := &in.Capabilities, &out.Capabilities
		*out = make(ConfigCapabilities, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	out.Whitelist = in.Whitelist
	out.LDAP = in.LDAP
	out.ArgoPermissions = in.ArgoPermissions
	if in.RoleMappings != nil {
		in, out := &in.RoleMappings, &out.RoleMappings
		*out = make(ConfigRoleMappings, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PaasConfigSpec.
func (in *PaasConfigSpec) DeepCopy() *PaasConfigSpec {
	if in == nil {
		return nil
	}
	out := new(PaasConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PaasConfigStatus) DeepCopyInto(out *PaasConfigStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PaasConfigStatus.
func (in *PaasConfigStatus) DeepCopy() *PaasConfigStatus {
	if in == nil {
		return nil
	}
	out := new(PaasConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PaasGroup) DeepCopyInto(out *PaasGroup) {
	*out = *in
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PaasGroup.
func (in *PaasGroup) DeepCopy() *PaasGroup {
	if in == nil {
		return nil
	}
	out := new(PaasGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PaasGroups) DeepCopyInto(out *PaasGroups) {
	{
		in := &in
		*out = make(PaasGroups, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PaasGroups.
func (in PaasGroups) DeepCopy() PaasGroups {
	if in == nil {
		return nil
	}
	out := new(PaasGroups)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PaasList) DeepCopyInto(out *PaasList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Paas, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PaasList.
func (in *PaasList) DeepCopy() *PaasList {
	if in == nil {
		return nil
	}
	out := new(PaasList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PaasList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PaasNS) DeepCopyInto(out *PaasNS) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PaasNS.
func (in *PaasNS) DeepCopy() *PaasNS {
	if in == nil {
		return nil
	}
	out := new(PaasNS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PaasNS) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PaasNSList) DeepCopyInto(out *PaasNSList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PaasNS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PaasNSList.
func (in *PaasNSList) DeepCopy() *PaasNSList {
	if in == nil {
		return nil
	}
	out := new(PaasNSList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PaasNSList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PaasNSSpec) DeepCopyInto(out *PaasNSSpec) {
	*out = *in
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SshSecrets != nil {
		in, out := &in.SshSecrets, &out.SshSecrets
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PaasNSSpec.
func (in *PaasNSSpec) DeepCopy() *PaasNSSpec {
	if in == nil {
		return nil
	}
	out := new(PaasNSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PaasNsStatus) DeepCopyInto(out *PaasNsStatus) {
	*out = *in
	if in.Messages != nil {
		in, out := &in.Messages, &out.Messages
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PaasNsStatus.
func (in *PaasNsStatus) DeepCopy() *PaasNsStatus {
	if in == nil {
		return nil
	}
	out := new(PaasNsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PaasSpec) DeepCopyInto(out *PaasSpec) {
	*out = *in
	if in.Capabilities != nil {
		in, out := &in.Capabilities, &out.Capabilities
		*out = make(PaasCapabilities, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make(PaasGroups, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Quota != nil {
		in, out := &in.Quota, &out.Quota
		*out = make(quota.Quota, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SshSecrets != nil {
		in, out := &in.SshSecrets, &out.SshSecrets
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PaasSpec.
func (in *PaasSpec) DeepCopy() *PaasSpec {
	if in == nil {
		return nil
	}
	out := new(PaasSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PaasStatus) DeepCopyInto(out *PaasStatus) {
	*out = *in
	if in.Messages != nil {
		in, out := &in.Messages, &out.Messages
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Quota != nil {
		in, out := &in.Quota, &out.Quota
		*out = make(map[string]quota.Quota, len(*in))
		for key, val := range *in {
			var outVal map[v1.ResourceName]resource.Quantity
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make(quota.Quota, len(*in))
				for key, val := range *in {
					(*out)[key] = val.DeepCopy()
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PaasStatus.
func (in *PaasStatus) DeepCopy() *PaasStatus {
	if in == nil {
		return nil
	}
	out := new(PaasStatus)
	in.DeepCopyInto(out)
	return out
}
