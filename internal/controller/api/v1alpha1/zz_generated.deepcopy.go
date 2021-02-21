// +build !ignore_autogenerated

/*
Copyright 2020.

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
	conditionsv1 "github.com/openshift/custom-resource-status/conditions/v1"
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Agent) DeepCopyInto(out *Agent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Agent.
func (in *Agent) DeepCopy() *Agent {
	if in == nil {
		return nil
	}
	out := new(Agent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Agent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentList) DeepCopyInto(out *AgentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Agent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentList.
func (in *AgentList) DeepCopy() *AgentList {
	if in == nil {
		return nil
	}
	out := new(AgentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AgentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentReference) DeepCopyInto(out *AgentReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentReference.
func (in *AgentReference) DeepCopy() *AgentReference {
	if in == nil {
		return nil
	}
	out := new(AgentReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentSpec) DeepCopyInto(out *AgentSpec) {
	*out = *in
	if in.ClusterDeploymentName != nil {
		in, out := &in.ClusterDeploymentName, &out.ClusterDeploymentName
		*out = new(ClusterReference)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentSpec.
func (in *AgentSpec) DeepCopy() *AgentSpec {
	if in == nil {
		return nil
	}
	out := new(AgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentStatus) DeepCopyInto(out *AgentStatus) {
	*out = *in
	if in.StateUpdatedTime != nil {
		in, out := &in.StateUpdatedTime, &out.StateUpdatedTime
		*out = (*in).DeepCopy()
	}
	if in.LogsCollectedTime != nil {
		in, out := &in.LogsCollectedTime, &out.LogsCollectedTime
		*out = (*in).DeepCopy()
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = (*in).DeepCopy()
	}
	if in.CheckedInTime != nil {
		in, out := &in.CheckedInTime, &out.CheckedInTime
		*out = (*in).DeepCopy()
	}
	in.Inventory.DeepCopyInto(&out.Inventory)
	out.ValidationInfo = in.ValidationInfo
	in.Progress.DeepCopyInto(&out.Progress)
	if in.Connectivity != nil {
		in, out := &in.Connectivity, &out.Connectivity
		*out = make([]HostConnectivityValidationInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NtpSources != nil {
		in, out := &in.NtpSources, &out.NtpSources
		*out = make([]HostNTPSources, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentStatus.
func (in *AgentStatus) DeepCopy() *AgentStatus {
	if in == nil {
		return nil
	}
	out := new(AgentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterProgressInfo) DeepCopyInto(out *ClusterProgressInfo) {
	*out = *in
	if in.LastProgressUpdateTime != nil {
		in, out := &in.LastProgressUpdateTime, &out.LastProgressUpdateTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterProgressInfo.
func (in *ClusterProgressInfo) DeepCopy() *ClusterProgressInfo {
	if in == nil {
		return nil
	}
	out := new(ClusterProgressInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterReference) DeepCopyInto(out *ClusterReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterReference.
func (in *ClusterReference) DeepCopy() *ClusterReference {
	if in == nil {
		return nil
	}
	out := new(ClusterReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.MachineNetworkCidr != nil {
		in, out := &in.MachineNetworkCidr, &out.MachineNetworkCidr
		*out = new(string)
		**out = **in
	}
	if in.PullSecretRef != nil {
		in, out := &in.PullSecretRef, &out.PullSecretRef
		*out = new(v1.SecretReference)
		**out = **in
	}
	in.ProvisionRequirements.DeepCopyInto(&out.ProvisionRequirements)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	if in.HostNetworks != nil {
		in, out := &in.HostNetworks, &out.HostNetworks
		*out = make([]HostNetwork, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InstallationStartTime != nil {
		in, out := &in.InstallationStartTime, &out.InstallationStartTime
		*out = (*in).DeepCopy()
	}
	if in.InstallationCompletionTime != nil {
		in, out := &in.InstallationCompletionTime, &out.InstallationCompletionTime
		*out = (*in).DeepCopy()
	}
	in.Progress.DeepCopyInto(&out.Progress)
	if in.LastUpdateTime != nil {
		in, out := &in.LastUpdateTime, &out.LastUpdateTime
		*out = (*in).DeepCopy()
	}
	if in.ControllerLogsCollectionTime != nil {
		in, out := &in.ControllerLogsCollectionTime, &out.ControllerLogsCollectionTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HardwareValidationInfo) DeepCopyInto(out *HardwareValidationInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HardwareValidationInfo.
func (in *HardwareValidationInfo) DeepCopy() *HardwareValidationInfo {
	if in == nil {
		return nil
	}
	out := new(HardwareValidationInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostBoot) DeepCopyInto(out *HostBoot) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostBoot.
func (in *HostBoot) DeepCopy() *HostBoot {
	if in == nil {
		return nil
	}
	out := new(HostBoot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostCPU) DeepCopyInto(out *HostCPU) {
	*out = *in
	if in.Flags != nil {
		in, out := &in.Flags, &out.Flags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostCPU.
func (in *HostCPU) DeepCopy() *HostCPU {
	if in == nil {
		return nil
	}
	out := new(HostCPU)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostConnectivityValidationInfo) DeepCopyInto(out *HostConnectivityValidationInfo) {
	*out = *in
	if in.HostDeploymentName != nil {
		in, out := &in.HostDeploymentName, &out.HostDeploymentName
		*out = new(AgentReference)
		**out = **in
	}
	out.L2Connectivity = in.L2Connectivity
	out.L3Connectivity = in.L3Connectivity
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostConnectivityValidationInfo.
func (in *HostConnectivityValidationInfo) DeepCopy() *HostConnectivityValidationInfo {
	if in == nil {
		return nil
	}
	out := new(HostConnectivityValidationInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostDisk) DeepCopyInto(out *HostDisk) {
	*out = *in
	in.InstallationEligibility.DeepCopyInto(&out.InstallationEligibility)
	out.IoPerf = in.IoPerf
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostDisk.
func (in *HostDisk) DeepCopy() *HostDisk {
	if in == nil {
		return nil
	}
	out := new(HostDisk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostIOPerf) DeepCopyInto(out *HostIOPerf) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostIOPerf.
func (in *HostIOPerf) DeepCopy() *HostIOPerf {
	if in == nil {
		return nil
	}
	out := new(HostIOPerf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostInstallationEligibility) DeepCopyInto(out *HostInstallationEligibility) {
	*out = *in
	if in.NotEligibleReasons != nil {
		in, out := &in.NotEligibleReasons, &out.NotEligibleReasons
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostInstallationEligibility.
func (in *HostInstallationEligibility) DeepCopy() *HostInstallationEligibility {
	if in == nil {
		return nil
	}
	out := new(HostInstallationEligibility)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostInterface) DeepCopyInto(out *HostInterface) {
	*out = *in
	if in.IPV6Addresses != nil {
		in, out := &in.IPV6Addresses, &out.IPV6Addresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IPV4Addresses != nil {
		in, out := &in.IPV4Addresses, &out.IPV4Addresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Flags != nil {
		in, out := &in.Flags, &out.Flags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostInterface.
func (in *HostInterface) DeepCopy() *HostInterface {
	if in == nil {
		return nil
	}
	out := new(HostInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostInventory) DeepCopyInto(out *HostInventory) {
	*out = *in
	if in.ReportTime != nil {
		in, out := &in.ReportTime, &out.ReportTime
		*out = (*in).DeepCopy()
	}
	out.Memory = in.Memory
	in.Cpu.DeepCopyInto(&out.Cpu)
	if in.Interfaces != nil {
		in, out := &in.Interfaces, &out.Interfaces
		*out = make([]HostInterface, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Disks != nil {
		in, out := &in.Disks, &out.Disks
		*out = make([]HostDisk, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Boot = in.Boot
	out.SystemVendor = in.SystemVendor
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostInventory.
func (in *HostInventory) DeepCopy() *HostInventory {
	if in == nil {
		return nil
	}
	out := new(HostInventory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostMemory) DeepCopyInto(out *HostMemory) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostMemory.
func (in *HostMemory) DeepCopy() *HostMemory {
	if in == nil {
		return nil
	}
	out := new(HostMemory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostNTPSources) DeepCopyInto(out *HostNTPSources) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostNTPSources.
func (in *HostNTPSources) DeepCopy() *HostNTPSources {
	if in == nil {
		return nil
	}
	out := new(HostNTPSources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostNetwork) DeepCopyInto(out *HostNetwork) {
	*out = *in
	if in.HostIds != nil {
		in, out := &in.HostIds, &out.HostIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostNetwork.
func (in *HostNetwork) DeepCopy() *HostNetwork {
	if in == nil {
		return nil
	}
	out := new(HostNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostProgressInfo) DeepCopyInto(out *HostProgressInfo) {
	*out = *in
	in.StageStartTime.DeepCopyInto(&out.StageStartTime)
	in.StageUpdateTime.DeepCopyInto(&out.StageUpdateTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostProgressInfo.
func (in *HostProgressInfo) DeepCopy() *HostProgressInfo {
	if in == nil {
		return nil
	}
	out := new(HostProgressInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostSystemVendor) DeepCopyInto(out *HostSystemVendor) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostSystemVendor.
func (in *HostSystemVendor) DeepCopy() *HostSystemVendor {
	if in == nil {
		return nil
	}
	out := new(HostSystemVendor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostValidationInfo) DeepCopyInto(out *HostValidationInfo) {
	*out = *in
	out.Hardware = in.Hardware
	out.Network = in.Network
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostValidationInfo.
func (in *HostValidationInfo) DeepCopy() *HostValidationInfo {
	if in == nil {
		return nil
	}
	out := new(HostValidationInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Image) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageList) DeepCopyInto(out *ImageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Image, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageList.
func (in *ImageList) DeepCopy() *ImageList {
	if in == nil {
		return nil
	}
	out := new(ImageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
	if in.ClusterRef != nil {
		in, out := &in.ClusterRef, &out.ClusterRef
		*out = new(ClusterReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStatus) DeepCopyInto(out *ImageStatus) {
	*out = *in
	if in.ExpirationTime != nil {
		in, out := &in.ExpirationTime, &out.ExpirationTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStatus.
func (in *ImageStatus) DeepCopy() *ImageStatus {
	if in == nil {
		return nil
	}
	out := new(ImageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallEnv) DeepCopyInto(out *InstallEnv) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallEnv.
func (in *InstallEnv) DeepCopy() *InstallEnv {
	if in == nil {
		return nil
	}
	out := new(InstallEnv)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstallEnv) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallEnvList) DeepCopyInto(out *InstallEnvList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InstallEnv, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallEnvList.
func (in *InstallEnvList) DeepCopy() *InstallEnvList {
	if in == nil {
		return nil
	}
	out := new(InstallEnvList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstallEnvList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallEnvSpec) DeepCopyInto(out *InstallEnvSpec) {
	*out = *in
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(Proxy)
		**out = **in
	}
	if in.AdditionalNTPSources != nil {
		in, out := &in.AdditionalNTPSources, &out.AdditionalNTPSources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SSHAuthorizedKeys != nil {
		in, out := &in.SSHAuthorizedKeys, &out.SSHAuthorizedKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PullSecretRef != nil {
		in, out := &in.PullSecretRef, &out.PullSecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	in.AgentLabelSelector.DeepCopyInto(&out.AgentLabelSelector)
	if in.AgentLabels != nil {
		in, out := &in.AgentLabels, &out.AgentLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.NMStateConfigLabelSelector.DeepCopyInto(&out.NMStateConfigLabelSelector)
	if in.ClusterRef != nil {
		in, out := &in.ClusterRef, &out.ClusterRef
		*out = new(ClusterReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallEnvSpec.
func (in *InstallEnvSpec) DeepCopy() *InstallEnvSpec {
	if in == nil {
		return nil
	}
	out := new(InstallEnvSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallEnvStatus) DeepCopyInto(out *InstallEnvStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditionsv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallEnvStatus.
func (in *InstallEnvStatus) DeepCopy() *InstallEnvStatus {
	if in == nil {
		return nil
	}
	out := new(InstallEnvStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L2Connectivity) DeepCopyInto(out *L2Connectivity) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L2Connectivity.
func (in *L2Connectivity) DeepCopy() *L2Connectivity {
	if in == nil {
		return nil
	}
	out := new(L2Connectivity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L3Connectivity) DeepCopyInto(out *L3Connectivity) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L3Connectivity.
func (in *L3Connectivity) DeepCopy() *L3Connectivity {
	if in == nil {
		return nil
	}
	out := new(L3Connectivity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NMStateConfig) DeepCopyInto(out *NMStateConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NMStateConfig.
func (in *NMStateConfig) DeepCopy() *NMStateConfig {
	if in == nil {
		return nil
	}
	out := new(NMStateConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NMStateConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NMStateConfigList) DeepCopyInto(out *NMStateConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NMStateConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NMStateConfigList.
func (in *NMStateConfigList) DeepCopy() *NMStateConfigList {
	if in == nil {
		return nil
	}
	out := new(NMStateConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NMStateConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NMStateConfigSpec) DeepCopyInto(out *NMStateConfigSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(RawConfig, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NMStateConfigSpec.
func (in *NMStateConfigSpec) DeepCopy() *NMStateConfigSpec {
	if in == nil {
		return nil
	}
	out := new(NMStateConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkValidationInfo) DeepCopyInto(out *NetworkValidationInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkValidationInfo.
func (in *NetworkValidationInfo) DeepCopy() *NetworkValidationInfo {
	if in == nil {
		return nil
	}
	out := new(NetworkValidationInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionRequirements) DeepCopyInto(out *ProvisionRequirements) {
	*out = *in
	in.AgentSelector.DeepCopyInto(&out.AgentSelector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionRequirements.
func (in *ProvisionRequirements) DeepCopy() *ProvisionRequirements {
	if in == nil {
		return nil
	}
	out := new(ProvisionRequirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Proxy) DeepCopyInto(out *Proxy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Proxy.
func (in *Proxy) DeepCopy() *Proxy {
	if in == nil {
		return nil
	}
	out := new(Proxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in RawConfig) DeepCopyInto(out *RawConfig) {
	{
		in := &in
		*out = make(RawConfig, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RawConfig.
func (in RawConfig) DeepCopy() RawConfig {
	if in == nil {
		return nil
	}
	out := new(RawConfig)
	in.DeepCopyInto(out)
	return *out
}
