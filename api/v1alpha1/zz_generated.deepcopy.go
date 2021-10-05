// +build !ignore_autogenerated

/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBSEConfig) DeepCopyInto(out *DBSEConfig) {
	*out = *in
	if in.AllowedRoles != nil {
		in, out := &in.AllowedRoles, &out.AllowedRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RootRotationStatements != nil {
		in, out := &in.RootRotationStatements, &out.RootRotationStatements
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DatabaseSpecificConfig != nil {
		in, out := &in.DatabaseSpecificConfig, &out.DatabaseSpecificConfig
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBSEConfig.
func (in *DBSEConfig) DeepCopy() *DBSEConfig {
	if in == nil {
		return nil
	}
	out := new(DBSEConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBSERole) DeepCopyInto(out *DBSERole) {
	*out = *in
	out.DefaultTTL = in.DefaultTTL
	out.MaxTTL = in.MaxTTL
	if in.CreationStatements != nil {
		in, out := &in.CreationStatements, &out.CreationStatements
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RevocationStatements != nil {
		in, out := &in.RevocationStatements, &out.RevocationStatements
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RollbackStatements != nil {
		in, out := &in.RollbackStatements, &out.RollbackStatements
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RenewStatements != nil {
		in, out := &in.RenewStatements, &out.RenewStatements
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBSERole.
func (in *DBSERole) DeepCopy() *DBSERole {
	if in == nil {
		return nil
	}
	out := new(DBSERole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseSecretEngineConfig) DeepCopyInto(out *DatabaseSecretEngineConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseSecretEngineConfig.
func (in *DatabaseSecretEngineConfig) DeepCopy() *DatabaseSecretEngineConfig {
	if in == nil {
		return nil
	}
	out := new(DatabaseSecretEngineConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseSecretEngineConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseSecretEngineConfigList) DeepCopyInto(out *DatabaseSecretEngineConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DatabaseSecretEngineConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseSecretEngineConfigList.
func (in *DatabaseSecretEngineConfigList) DeepCopy() *DatabaseSecretEngineConfigList {
	if in == nil {
		return nil
	}
	out := new(DatabaseSecretEngineConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseSecretEngineConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseSecretEngineConfigSpec) DeepCopyInto(out *DatabaseSecretEngineConfigSpec) {
	*out = *in
	out.Authentication = in.Authentication
	in.DBSEConfig.DeepCopyInto(&out.DBSEConfig)
	if in.RootCredentialsFromVaultSecret != nil {
		in, out := &in.RootCredentialsFromVaultSecret, &out.RootCredentialsFromVaultSecret
		*out = new(VaultSecretReference)
		**out = **in
	}
	if in.RootCredentialsFromSecret != nil {
		in, out := &in.RootCredentialsFromSecret, &out.RootCredentialsFromSecret
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.RootCredentialsFromRandomSecret != nil {
		in, out := &in.RootCredentialsFromRandomSecret, &out.RootCredentialsFromRandomSecret
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseSecretEngineConfigSpec.
func (in *DatabaseSecretEngineConfigSpec) DeepCopy() *DatabaseSecretEngineConfigSpec {
	if in == nil {
		return nil
	}
	out := new(DatabaseSecretEngineConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseSecretEngineConfigStatus) DeepCopyInto(out *DatabaseSecretEngineConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseSecretEngineConfigStatus.
func (in *DatabaseSecretEngineConfigStatus) DeepCopy() *DatabaseSecretEngineConfigStatus {
	if in == nil {
		return nil
	}
	out := new(DatabaseSecretEngineConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseSecretEngineRole) DeepCopyInto(out *DatabaseSecretEngineRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseSecretEngineRole.
func (in *DatabaseSecretEngineRole) DeepCopy() *DatabaseSecretEngineRole {
	if in == nil {
		return nil
	}
	out := new(DatabaseSecretEngineRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseSecretEngineRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseSecretEngineRoleList) DeepCopyInto(out *DatabaseSecretEngineRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DatabaseSecretEngineRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseSecretEngineRoleList.
func (in *DatabaseSecretEngineRoleList) DeepCopy() *DatabaseSecretEngineRoleList {
	if in == nil {
		return nil
	}
	out := new(DatabaseSecretEngineRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseSecretEngineRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseSecretEngineRoleSpec) DeepCopyInto(out *DatabaseSecretEngineRoleSpec) {
	*out = *in
	out.Authentication = in.Authentication
	in.DBSERole.DeepCopyInto(&out.DBSERole)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseSecretEngineRoleSpec.
func (in *DatabaseSecretEngineRoleSpec) DeepCopy() *DatabaseSecretEngineRoleSpec {
	if in == nil {
		return nil
	}
	out := new(DatabaseSecretEngineRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseSecretEngineRoleStatus) DeepCopyInto(out *DatabaseSecretEngineRoleStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseSecretEngineRoleStatus.
func (in *DatabaseSecretEngineRoleStatus) DeepCopy() *DatabaseSecretEngineRoleStatus {
	if in == nil {
		return nil
	}
	out := new(DatabaseSecretEngineRoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeAuthConfiguration) DeepCopyInto(out *KubeAuthConfiguration) {
	*out = *in
	out.ServiceAccount = in.ServiceAccount
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeAuthConfiguration.
func (in *KubeAuthConfiguration) DeepCopy() *KubeAuthConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeAuthConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Mount) DeepCopyInto(out *Mount) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Mount.
func (in *Mount) DeepCopy() *Mount {
	if in == nil {
		return nil
	}
	out := new(Mount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MountConfig) DeepCopyInto(out *MountConfig) {
	*out = *in
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.AuditNonHMACRequestKeys != nil {
		in, out := &in.AuditNonHMACRequestKeys, &out.AuditNonHMACRequestKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AuditNonHMACResponseKeys != nil {
		in, out := &in.AuditNonHMACResponseKeys, &out.AuditNonHMACResponseKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PassthroughRequestHeaders != nil {
		in, out := &in.PassthroughRequestHeaders, &out.PassthroughRequestHeaders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowedResponseHeaders != nil {
		in, out := &in.AllowedResponseHeaders, &out.AllowedResponseHeaders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountConfig.
func (in *MountConfig) DeepCopy() *MountConfig {
	if in == nil {
		return nil
	}
	out := new(MountConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PasswordPolicy) DeepCopyInto(out *PasswordPolicy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PasswordPolicy.
func (in *PasswordPolicy) DeepCopy() *PasswordPolicy {
	if in == nil {
		return nil
	}
	out := new(PasswordPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PasswordPolicyFormat) DeepCopyInto(out *PasswordPolicyFormat) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]PasswordPolicyRule, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PasswordPolicyFormat.
func (in *PasswordPolicyFormat) DeepCopy() *PasswordPolicyFormat {
	if in == nil {
		return nil
	}
	out := new(PasswordPolicyFormat)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PasswordPolicyRule) DeepCopyInto(out *PasswordPolicyRule) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PasswordPolicyRule.
func (in *PasswordPolicyRule) DeepCopy() *PasswordPolicyRule {
	if in == nil {
		return nil
	}
	out := new(PasswordPolicyRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Policy) DeepCopyInto(out *Policy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Policy.
func (in *Policy) DeepCopy() *Policy {
	if in == nil {
		return nil
	}
	out := new(Policy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Policy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyList) DeepCopyInto(out *PolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Policy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyList.
func (in *PolicyList) DeepCopy() *PolicyList {
	if in == nil {
		return nil
	}
	out := new(PolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpec) DeepCopyInto(out *PolicySpec) {
	*out = *in
	out.Authentication = in.Authentication
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpec.
func (in *PolicySpec) DeepCopy() *PolicySpec {
	if in == nil {
		return nil
	}
	out := new(PolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyStatus) DeepCopyInto(out *PolicyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyStatus.
func (in *PolicyStatus) DeepCopy() *PolicyStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RandomSecret) DeepCopyInto(out *RandomSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RandomSecret.
func (in *RandomSecret) DeepCopy() *RandomSecret {
	if in == nil {
		return nil
	}
	out := new(RandomSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RandomSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RandomSecretList) DeepCopyInto(out *RandomSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RandomSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RandomSecretList.
func (in *RandomSecretList) DeepCopy() *RandomSecretList {
	if in == nil {
		return nil
	}
	out := new(RandomSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RandomSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RandomSecretSpec) DeepCopyInto(out *RandomSecretSpec) {
	*out = *in
	out.Authentication = in.Authentication
	if in.SecretFormat != nil {
		in, out := &in.SecretFormat, &out.SecretFormat
		*out = make(map[string]PasswordPolicy, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.RefreshPeriod = in.RefreshPeriod
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RandomSecretSpec.
func (in *RandomSecretSpec) DeepCopy() *RandomSecretSpec {
	if in == nil {
		return nil
	}
	out := new(RandomSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RandomSecretStatus) DeepCopyInto(out *RandomSecretStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RandomSecretStatus.
func (in *RandomSecretStatus) DeepCopy() *RandomSecretStatus {
	if in == nil {
		return nil
	}
	out := new(RandomSecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretEngineMount) DeepCopyInto(out *SecretEngineMount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretEngineMount.
func (in *SecretEngineMount) DeepCopy() *SecretEngineMount {
	if in == nil {
		return nil
	}
	out := new(SecretEngineMount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretEngineMount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretEngineMountList) DeepCopyInto(out *SecretEngineMountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecretEngineMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretEngineMountList.
func (in *SecretEngineMountList) DeepCopy() *SecretEngineMountList {
	if in == nil {
		return nil
	}
	out := new(SecretEngineMountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretEngineMountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretEngineMountSpec) DeepCopyInto(out *SecretEngineMountSpec) {
	*out = *in
	out.Authentication = in.Authentication
	in.Mount.DeepCopyInto(&out.Mount)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretEngineMountSpec.
func (in *SecretEngineMountSpec) DeepCopy() *SecretEngineMountSpec {
	if in == nil {
		return nil
	}
	out := new(SecretEngineMountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretEngineMountStatus) DeepCopyInto(out *SecretEngineMountStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretEngineMountStatus.
func (in *SecretEngineMountStatus) DeepCopy() *SecretEngineMountStatus {
	if in == nil {
		return nil
	}
	out := new(SecretEngineMountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VRole) DeepCopyInto(out *VRole) {
	*out = *in
	if in.TargetServiceAccounts != nil {
		in, out := &in.TargetServiceAccounts, &out.TargetServiceAccounts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TokenBoundCIDRs != nil {
		in, out := &in.TokenBoundCIDRs, &out.TokenBoundCIDRs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VRole.
func (in *VRole) DeepCopy() *VRole {
	if in == nil {
		return nil
	}
	out := new(VRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultRole) DeepCopyInto(out *VaultRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultRole.
func (in *VaultRole) DeepCopy() *VaultRole {
	if in == nil {
		return nil
	}
	out := new(VaultRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VaultRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultRoleList) DeepCopyInto(out *VaultRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VaultRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultRoleList.
func (in *VaultRoleList) DeepCopy() *VaultRoleList {
	if in == nil {
		return nil
	}
	out := new(VaultRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VaultRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultRoleSpec) DeepCopyInto(out *VaultRoleSpec) {
	*out = *in
	out.Authentication = in.Authentication
	in.VRole.DeepCopyInto(&out.VRole)
	if in.TargetNamespaceSelector != nil {
		in, out := &in.TargetNamespaceSelector, &out.TargetNamespaceSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.TargetNamespaces != nil {
		in, out := &in.TargetNamespaces, &out.TargetNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultRoleSpec.
func (in *VaultRoleSpec) DeepCopy() *VaultRoleSpec {
	if in == nil {
		return nil
	}
	out := new(VaultRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultRoleStatus) DeepCopyInto(out *VaultRoleStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultRoleStatus.
func (in *VaultRoleStatus) DeepCopy() *VaultRoleStatus {
	if in == nil {
		return nil
	}
	out := new(VaultRoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultSecretReference) DeepCopyInto(out *VaultSecretReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultSecretReference.
func (in *VaultSecretReference) DeepCopy() *VaultSecretReference {
	if in == nil {
		return nil
	}
	out := new(VaultSecretReference)
	in.DeepCopyInto(out)
	return out
}
