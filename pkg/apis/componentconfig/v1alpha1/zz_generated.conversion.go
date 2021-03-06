// +build !ignore_autogenerated

/*
Copyright 2018 (c) 2018 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	componentconfig "github.com/gardener/gardener/pkg/apis/componentconfig"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_BackupInfrastructureControllerConfiguration_To_componentconfig_BackupInfrastructureControllerConfiguration,
		Convert_componentconfig_BackupInfrastructureControllerConfiguration_To_v1alpha1_BackupInfrastructureControllerConfiguration,
		Convert_v1alpha1_ClientConnectionConfiguration_To_componentconfig_ClientConnectionConfiguration,
		Convert_componentconfig_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration,
		Convert_v1alpha1_CloudProfileControllerConfiguration_To_componentconfig_CloudProfileControllerConfiguration,
		Convert_componentconfig_CloudProfileControllerConfiguration_To_v1alpha1_CloudProfileControllerConfiguration,
		Convert_v1alpha1_ControllerManagerConfiguration_To_componentconfig_ControllerManagerConfiguration,
		Convert_componentconfig_ControllerManagerConfiguration_To_v1alpha1_ControllerManagerConfiguration,
		Convert_v1alpha1_ControllerManagerControllerConfiguration_To_componentconfig_ControllerManagerControllerConfiguration,
		Convert_componentconfig_ControllerManagerControllerConfiguration_To_v1alpha1_ControllerManagerControllerConfiguration,
		Convert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration,
		Convert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration,
		Convert_v1alpha1_MetricsConfiguration_To_componentconfig_MetricsConfiguration,
		Convert_componentconfig_MetricsConfiguration_To_v1alpha1_MetricsConfiguration,
		Convert_v1alpha1_QuotaControllerConfiguration_To_componentconfig_QuotaControllerConfiguration,
		Convert_componentconfig_QuotaControllerConfiguration_To_v1alpha1_QuotaControllerConfiguration,
		Convert_v1alpha1_SecretBindingControllerConfiguration_To_componentconfig_SecretBindingControllerConfiguration,
		Convert_componentconfig_SecretBindingControllerConfiguration_To_v1alpha1_SecretBindingControllerConfiguration,
		Convert_v1alpha1_SeedControllerConfiguration_To_componentconfig_SeedControllerConfiguration,
		Convert_componentconfig_SeedControllerConfiguration_To_v1alpha1_SeedControllerConfiguration,
		Convert_v1alpha1_ServerConfiguration_To_componentconfig_ServerConfiguration,
		Convert_componentconfig_ServerConfiguration_To_v1alpha1_ServerConfiguration,
		Convert_v1alpha1_ShootCareControllerConfiguration_To_componentconfig_ShootCareControllerConfiguration,
		Convert_componentconfig_ShootCareControllerConfiguration_To_v1alpha1_ShootCareControllerConfiguration,
		Convert_v1alpha1_ShootControllerConfiguration_To_componentconfig_ShootControllerConfiguration,
		Convert_componentconfig_ShootControllerConfiguration_To_v1alpha1_ShootControllerConfiguration,
		Convert_v1alpha1_ShootMaintenanceControllerConfiguration_To_componentconfig_ShootMaintenanceControllerConfiguration,
		Convert_componentconfig_ShootMaintenanceControllerConfiguration_To_v1alpha1_ShootMaintenanceControllerConfiguration,
		Convert_v1alpha1_ShootQuotaControllerConfiguration_To_componentconfig_ShootQuotaControllerConfiguration,
		Convert_componentconfig_ShootQuotaControllerConfiguration_To_v1alpha1_ShootQuotaControllerConfiguration,
	)
}

func autoConvert_v1alpha1_BackupInfrastructureControllerConfiguration_To_componentconfig_BackupInfrastructureControllerConfiguration(in *BackupInfrastructureControllerConfiguration, out *componentconfig.BackupInfrastructureControllerConfiguration, s conversion.Scope) error {
	out.ConcurrentSyncs = in.ConcurrentSyncs
	out.SyncPeriod = in.SyncPeriod
	out.DeletionGracePeriodDays = (*int)(unsafe.Pointer(in.DeletionGracePeriodDays))
	return nil
}

// Convert_v1alpha1_BackupInfrastructureControllerConfiguration_To_componentconfig_BackupInfrastructureControllerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_BackupInfrastructureControllerConfiguration_To_componentconfig_BackupInfrastructureControllerConfiguration(in *BackupInfrastructureControllerConfiguration, out *componentconfig.BackupInfrastructureControllerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_BackupInfrastructureControllerConfiguration_To_componentconfig_BackupInfrastructureControllerConfiguration(in, out, s)
}

func autoConvert_componentconfig_BackupInfrastructureControllerConfiguration_To_v1alpha1_BackupInfrastructureControllerConfiguration(in *componentconfig.BackupInfrastructureControllerConfiguration, out *BackupInfrastructureControllerConfiguration, s conversion.Scope) error {
	out.ConcurrentSyncs = in.ConcurrentSyncs
	out.SyncPeriod = in.SyncPeriod
	out.DeletionGracePeriodDays = (*int)(unsafe.Pointer(in.DeletionGracePeriodDays))
	return nil
}

// Convert_componentconfig_BackupInfrastructureControllerConfiguration_To_v1alpha1_BackupInfrastructureControllerConfiguration is an autogenerated conversion function.
func Convert_componentconfig_BackupInfrastructureControllerConfiguration_To_v1alpha1_BackupInfrastructureControllerConfiguration(in *componentconfig.BackupInfrastructureControllerConfiguration, out *BackupInfrastructureControllerConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_BackupInfrastructureControllerConfiguration_To_v1alpha1_BackupInfrastructureControllerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_ClientConnectionConfiguration_To_componentconfig_ClientConnectionConfiguration(in *ClientConnectionConfiguration, out *componentconfig.ClientConnectionConfiguration, s conversion.Scope) error {
	out.KubeConfigFile = in.KubeConfigFile
	out.AcceptContentTypes = in.AcceptContentTypes
	out.ContentType = in.ContentType
	out.QPS = in.QPS
	out.Burst = in.Burst
	return nil
}

// Convert_v1alpha1_ClientConnectionConfiguration_To_componentconfig_ClientConnectionConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ClientConnectionConfiguration_To_componentconfig_ClientConnectionConfiguration(in *ClientConnectionConfiguration, out *componentconfig.ClientConnectionConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClientConnectionConfiguration_To_componentconfig_ClientConnectionConfiguration(in, out, s)
}

func autoConvert_componentconfig_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(in *componentconfig.ClientConnectionConfiguration, out *ClientConnectionConfiguration, s conversion.Scope) error {
	out.KubeConfigFile = in.KubeConfigFile
	out.AcceptContentTypes = in.AcceptContentTypes
	out.ContentType = in.ContentType
	out.QPS = in.QPS
	out.Burst = in.Burst
	return nil
}

// Convert_componentconfig_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration is an autogenerated conversion function.
func Convert_componentconfig_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(in *componentconfig.ClientConnectionConfiguration, out *ClientConnectionConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(in, out, s)
}

func autoConvert_v1alpha1_CloudProfileControllerConfiguration_To_componentconfig_CloudProfileControllerConfiguration(in *CloudProfileControllerConfiguration, out *componentconfig.CloudProfileControllerConfiguration, s conversion.Scope) error {
	out.ConcurrentSyncs = in.ConcurrentSyncs
	return nil
}

// Convert_v1alpha1_CloudProfileControllerConfiguration_To_componentconfig_CloudProfileControllerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_CloudProfileControllerConfiguration_To_componentconfig_CloudProfileControllerConfiguration(in *CloudProfileControllerConfiguration, out *componentconfig.CloudProfileControllerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_CloudProfileControllerConfiguration_To_componentconfig_CloudProfileControllerConfiguration(in, out, s)
}

func autoConvert_componentconfig_CloudProfileControllerConfiguration_To_v1alpha1_CloudProfileControllerConfiguration(in *componentconfig.CloudProfileControllerConfiguration, out *CloudProfileControllerConfiguration, s conversion.Scope) error {
	out.ConcurrentSyncs = in.ConcurrentSyncs
	return nil
}

// Convert_componentconfig_CloudProfileControllerConfiguration_To_v1alpha1_CloudProfileControllerConfiguration is an autogenerated conversion function.
func Convert_componentconfig_CloudProfileControllerConfiguration_To_v1alpha1_CloudProfileControllerConfiguration(in *componentconfig.CloudProfileControllerConfiguration, out *CloudProfileControllerConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_CloudProfileControllerConfiguration_To_v1alpha1_CloudProfileControllerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_ControllerManagerConfiguration_To_componentconfig_ControllerManagerConfiguration(in *ControllerManagerConfiguration, out *componentconfig.ControllerManagerConfiguration, s conversion.Scope) error {
	if err := Convert_v1alpha1_ClientConnectionConfiguration_To_componentconfig_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
		return err
	}
	out.GardenerClientConnection = (*componentconfig.ClientConnectionConfiguration)(unsafe.Pointer(in.GardenerClientConnection))
	if err := Convert_v1alpha1_ControllerManagerControllerConfiguration_To_componentconfig_ControllerManagerControllerConfiguration(&in.Controllers, &out.Controllers, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	out.LogLevel = in.LogLevel
	if err := Convert_v1alpha1_MetricsConfiguration_To_componentconfig_MetricsConfiguration(&in.Metrics, &out.Metrics, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ServerConfiguration_To_componentconfig_ServerConfiguration(&in.Server, &out.Server, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ControllerManagerConfiguration_To_componentconfig_ControllerManagerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ControllerManagerConfiguration_To_componentconfig_ControllerManagerConfiguration(in *ControllerManagerConfiguration, out *componentconfig.ControllerManagerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ControllerManagerConfiguration_To_componentconfig_ControllerManagerConfiguration(in, out, s)
}

func autoConvert_componentconfig_ControllerManagerConfiguration_To_v1alpha1_ControllerManagerConfiguration(in *componentconfig.ControllerManagerConfiguration, out *ControllerManagerConfiguration, s conversion.Scope) error {
	if err := Convert_componentconfig_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
		return err
	}
	out.GardenerClientConnection = (*ClientConnectionConfiguration)(unsafe.Pointer(in.GardenerClientConnection))
	if err := Convert_componentconfig_ControllerManagerControllerConfiguration_To_v1alpha1_ControllerManagerControllerConfiguration(&in.Controllers, &out.Controllers, s); err != nil {
		return err
	}
	if err := Convert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	out.LogLevel = in.LogLevel
	if err := Convert_componentconfig_MetricsConfiguration_To_v1alpha1_MetricsConfiguration(&in.Metrics, &out.Metrics, s); err != nil {
		return err
	}
	if err := Convert_componentconfig_ServerConfiguration_To_v1alpha1_ServerConfiguration(&in.Server, &out.Server, s); err != nil {
		return err
	}
	return nil
}

// Convert_componentconfig_ControllerManagerConfiguration_To_v1alpha1_ControllerManagerConfiguration is an autogenerated conversion function.
func Convert_componentconfig_ControllerManagerConfiguration_To_v1alpha1_ControllerManagerConfiguration(in *componentconfig.ControllerManagerConfiguration, out *ControllerManagerConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_ControllerManagerConfiguration_To_v1alpha1_ControllerManagerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_ControllerManagerControllerConfiguration_To_componentconfig_ControllerManagerControllerConfiguration(in *ControllerManagerControllerConfiguration, out *componentconfig.ControllerManagerControllerConfiguration, s conversion.Scope) error {
	out.CloudProfile = (*componentconfig.CloudProfileControllerConfiguration)(unsafe.Pointer(in.CloudProfile))
	out.SecretBinding = (*componentconfig.SecretBindingControllerConfiguration)(unsafe.Pointer(in.SecretBinding))
	out.Quota = (*componentconfig.QuotaControllerConfiguration)(unsafe.Pointer(in.Quota))
	out.Seed = (*componentconfig.SeedControllerConfiguration)(unsafe.Pointer(in.Seed))
	if err := Convert_v1alpha1_ShootControllerConfiguration_To_componentconfig_ShootControllerConfiguration(&in.Shoot, &out.Shoot, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ShootCareControllerConfiguration_To_componentconfig_ShootCareControllerConfiguration(&in.ShootCare, &out.ShootCare, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ShootMaintenanceControllerConfiguration_To_componentconfig_ShootMaintenanceControllerConfiguration(&in.ShootMaintenance, &out.ShootMaintenance, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ShootQuotaControllerConfiguration_To_componentconfig_ShootQuotaControllerConfiguration(&in.ShootQuota, &out.ShootQuota, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_BackupInfrastructureControllerConfiguration_To_componentconfig_BackupInfrastructureControllerConfiguration(&in.BackupInfrastructure, &out.BackupInfrastructure, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ControllerManagerControllerConfiguration_To_componentconfig_ControllerManagerControllerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ControllerManagerControllerConfiguration_To_componentconfig_ControllerManagerControllerConfiguration(in *ControllerManagerControllerConfiguration, out *componentconfig.ControllerManagerControllerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ControllerManagerControllerConfiguration_To_componentconfig_ControllerManagerControllerConfiguration(in, out, s)
}

func autoConvert_componentconfig_ControllerManagerControllerConfiguration_To_v1alpha1_ControllerManagerControllerConfiguration(in *componentconfig.ControllerManagerControllerConfiguration, out *ControllerManagerControllerConfiguration, s conversion.Scope) error {
	out.CloudProfile = (*CloudProfileControllerConfiguration)(unsafe.Pointer(in.CloudProfile))
	out.SecretBinding = (*SecretBindingControllerConfiguration)(unsafe.Pointer(in.SecretBinding))
	out.Quota = (*QuotaControllerConfiguration)(unsafe.Pointer(in.Quota))
	out.Seed = (*SeedControllerConfiguration)(unsafe.Pointer(in.Seed))
	if err := Convert_componentconfig_ShootControllerConfiguration_To_v1alpha1_ShootControllerConfiguration(&in.Shoot, &out.Shoot, s); err != nil {
		return err
	}
	if err := Convert_componentconfig_ShootCareControllerConfiguration_To_v1alpha1_ShootCareControllerConfiguration(&in.ShootCare, &out.ShootCare, s); err != nil {
		return err
	}
	if err := Convert_componentconfig_ShootMaintenanceControllerConfiguration_To_v1alpha1_ShootMaintenanceControllerConfiguration(&in.ShootMaintenance, &out.ShootMaintenance, s); err != nil {
		return err
	}
	if err := Convert_componentconfig_ShootQuotaControllerConfiguration_To_v1alpha1_ShootQuotaControllerConfiguration(&in.ShootQuota, &out.ShootQuota, s); err != nil {
		return err
	}
	if err := Convert_componentconfig_BackupInfrastructureControllerConfiguration_To_v1alpha1_BackupInfrastructureControllerConfiguration(&in.BackupInfrastructure, &out.BackupInfrastructure, s); err != nil {
		return err
	}
	return nil
}

// Convert_componentconfig_ControllerManagerControllerConfiguration_To_v1alpha1_ControllerManagerControllerConfiguration is an autogenerated conversion function.
func Convert_componentconfig_ControllerManagerControllerConfiguration_To_v1alpha1_ControllerManagerControllerConfiguration(in *componentconfig.ControllerManagerControllerConfiguration, out *ControllerManagerControllerConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_ControllerManagerControllerConfiguration_To_v1alpha1_ControllerManagerControllerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(in *LeaderElectionConfiguration, out *componentconfig.LeaderElectionConfiguration, s conversion.Scope) error {
	out.LeaderElect = in.LeaderElect
	out.LeaseDuration = in.LeaseDuration
	out.RenewDeadline = in.RenewDeadline
	out.RetryPeriod = in.RetryPeriod
	out.ResourceLock = in.ResourceLock
	out.LockObjectNamespace = in.LockObjectNamespace
	out.LockObjectName = in.LockObjectName
	return nil
}

// Convert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(in *LeaderElectionConfiguration, out *componentconfig.LeaderElectionConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(in, out, s)
}

func autoConvert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(in *componentconfig.LeaderElectionConfiguration, out *LeaderElectionConfiguration, s conversion.Scope) error {
	out.LeaderElect = in.LeaderElect
	out.LeaseDuration = in.LeaseDuration
	out.RenewDeadline = in.RenewDeadline
	out.RetryPeriod = in.RetryPeriod
	out.ResourceLock = in.ResourceLock
	out.LockObjectNamespace = in.LockObjectNamespace
	out.LockObjectName = in.LockObjectName
	return nil
}

// Convert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration is an autogenerated conversion function.
func Convert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(in *componentconfig.LeaderElectionConfiguration, out *LeaderElectionConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(in, out, s)
}

func autoConvert_v1alpha1_MetricsConfiguration_To_componentconfig_MetricsConfiguration(in *MetricsConfiguration, out *componentconfig.MetricsConfiguration, s conversion.Scope) error {
	out.Interval = in.Interval
	return nil
}

// Convert_v1alpha1_MetricsConfiguration_To_componentconfig_MetricsConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_MetricsConfiguration_To_componentconfig_MetricsConfiguration(in *MetricsConfiguration, out *componentconfig.MetricsConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_MetricsConfiguration_To_componentconfig_MetricsConfiguration(in, out, s)
}

func autoConvert_componentconfig_MetricsConfiguration_To_v1alpha1_MetricsConfiguration(in *componentconfig.MetricsConfiguration, out *MetricsConfiguration, s conversion.Scope) error {
	out.Interval = in.Interval
	return nil
}

// Convert_componentconfig_MetricsConfiguration_To_v1alpha1_MetricsConfiguration is an autogenerated conversion function.
func Convert_componentconfig_MetricsConfiguration_To_v1alpha1_MetricsConfiguration(in *componentconfig.MetricsConfiguration, out *MetricsConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_MetricsConfiguration_To_v1alpha1_MetricsConfiguration(in, out, s)
}

func autoConvert_v1alpha1_QuotaControllerConfiguration_To_componentconfig_QuotaControllerConfiguration(in *QuotaControllerConfiguration, out *componentconfig.QuotaControllerConfiguration, s conversion.Scope) error {
	out.ConcurrentSyncs = in.ConcurrentSyncs
	return nil
}

// Convert_v1alpha1_QuotaControllerConfiguration_To_componentconfig_QuotaControllerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_QuotaControllerConfiguration_To_componentconfig_QuotaControllerConfiguration(in *QuotaControllerConfiguration, out *componentconfig.QuotaControllerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_QuotaControllerConfiguration_To_componentconfig_QuotaControllerConfiguration(in, out, s)
}

func autoConvert_componentconfig_QuotaControllerConfiguration_To_v1alpha1_QuotaControllerConfiguration(in *componentconfig.QuotaControllerConfiguration, out *QuotaControllerConfiguration, s conversion.Scope) error {
	out.ConcurrentSyncs = in.ConcurrentSyncs
	return nil
}

// Convert_componentconfig_QuotaControllerConfiguration_To_v1alpha1_QuotaControllerConfiguration is an autogenerated conversion function.
func Convert_componentconfig_QuotaControllerConfiguration_To_v1alpha1_QuotaControllerConfiguration(in *componentconfig.QuotaControllerConfiguration, out *QuotaControllerConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_QuotaControllerConfiguration_To_v1alpha1_QuotaControllerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_SecretBindingControllerConfiguration_To_componentconfig_SecretBindingControllerConfiguration(in *SecretBindingControllerConfiguration, out *componentconfig.SecretBindingControllerConfiguration, s conversion.Scope) error {
	out.ConcurrentSyncs = in.ConcurrentSyncs
	return nil
}

// Convert_v1alpha1_SecretBindingControllerConfiguration_To_componentconfig_SecretBindingControllerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_SecretBindingControllerConfiguration_To_componentconfig_SecretBindingControllerConfiguration(in *SecretBindingControllerConfiguration, out *componentconfig.SecretBindingControllerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_SecretBindingControllerConfiguration_To_componentconfig_SecretBindingControllerConfiguration(in, out, s)
}

func autoConvert_componentconfig_SecretBindingControllerConfiguration_To_v1alpha1_SecretBindingControllerConfiguration(in *componentconfig.SecretBindingControllerConfiguration, out *SecretBindingControllerConfiguration, s conversion.Scope) error {
	out.ConcurrentSyncs = in.ConcurrentSyncs
	return nil
}

// Convert_componentconfig_SecretBindingControllerConfiguration_To_v1alpha1_SecretBindingControllerConfiguration is an autogenerated conversion function.
func Convert_componentconfig_SecretBindingControllerConfiguration_To_v1alpha1_SecretBindingControllerConfiguration(in *componentconfig.SecretBindingControllerConfiguration, out *SecretBindingControllerConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_SecretBindingControllerConfiguration_To_v1alpha1_SecretBindingControllerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_SeedControllerConfiguration_To_componentconfig_SeedControllerConfiguration(in *SeedControllerConfiguration, out *componentconfig.SeedControllerConfiguration, s conversion.Scope) error {
	out.ConcurrentSyncs = in.ConcurrentSyncs
	return nil
}

// Convert_v1alpha1_SeedControllerConfiguration_To_componentconfig_SeedControllerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_SeedControllerConfiguration_To_componentconfig_SeedControllerConfiguration(in *SeedControllerConfiguration, out *componentconfig.SeedControllerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_SeedControllerConfiguration_To_componentconfig_SeedControllerConfiguration(in, out, s)
}

func autoConvert_componentconfig_SeedControllerConfiguration_To_v1alpha1_SeedControllerConfiguration(in *componentconfig.SeedControllerConfiguration, out *SeedControllerConfiguration, s conversion.Scope) error {
	out.ConcurrentSyncs = in.ConcurrentSyncs
	return nil
}

// Convert_componentconfig_SeedControllerConfiguration_To_v1alpha1_SeedControllerConfiguration is an autogenerated conversion function.
func Convert_componentconfig_SeedControllerConfiguration_To_v1alpha1_SeedControllerConfiguration(in *componentconfig.SeedControllerConfiguration, out *SeedControllerConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_SeedControllerConfiguration_To_v1alpha1_SeedControllerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_ServerConfiguration_To_componentconfig_ServerConfiguration(in *ServerConfiguration, out *componentconfig.ServerConfiguration, s conversion.Scope) error {
	out.BindAddress = in.BindAddress
	out.Port = in.Port
	return nil
}

// Convert_v1alpha1_ServerConfiguration_To_componentconfig_ServerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ServerConfiguration_To_componentconfig_ServerConfiguration(in *ServerConfiguration, out *componentconfig.ServerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ServerConfiguration_To_componentconfig_ServerConfiguration(in, out, s)
}

func autoConvert_componentconfig_ServerConfiguration_To_v1alpha1_ServerConfiguration(in *componentconfig.ServerConfiguration, out *ServerConfiguration, s conversion.Scope) error {
	out.BindAddress = in.BindAddress
	out.Port = in.Port
	return nil
}

// Convert_componentconfig_ServerConfiguration_To_v1alpha1_ServerConfiguration is an autogenerated conversion function.
func Convert_componentconfig_ServerConfiguration_To_v1alpha1_ServerConfiguration(in *componentconfig.ServerConfiguration, out *ServerConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_ServerConfiguration_To_v1alpha1_ServerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_ShootCareControllerConfiguration_To_componentconfig_ShootCareControllerConfiguration(in *ShootCareControllerConfiguration, out *componentconfig.ShootCareControllerConfiguration, s conversion.Scope) error {
	out.ConcurrentSyncs = in.ConcurrentSyncs
	out.SyncPeriod = in.SyncPeriod
	return nil
}

// Convert_v1alpha1_ShootCareControllerConfiguration_To_componentconfig_ShootCareControllerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ShootCareControllerConfiguration_To_componentconfig_ShootCareControllerConfiguration(in *ShootCareControllerConfiguration, out *componentconfig.ShootCareControllerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ShootCareControllerConfiguration_To_componentconfig_ShootCareControllerConfiguration(in, out, s)
}

func autoConvert_componentconfig_ShootCareControllerConfiguration_To_v1alpha1_ShootCareControllerConfiguration(in *componentconfig.ShootCareControllerConfiguration, out *ShootCareControllerConfiguration, s conversion.Scope) error {
	out.ConcurrentSyncs = in.ConcurrentSyncs
	out.SyncPeriod = in.SyncPeriod
	return nil
}

// Convert_componentconfig_ShootCareControllerConfiguration_To_v1alpha1_ShootCareControllerConfiguration is an autogenerated conversion function.
func Convert_componentconfig_ShootCareControllerConfiguration_To_v1alpha1_ShootCareControllerConfiguration(in *componentconfig.ShootCareControllerConfiguration, out *ShootCareControllerConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_ShootCareControllerConfiguration_To_v1alpha1_ShootCareControllerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_ShootControllerConfiguration_To_componentconfig_ShootControllerConfiguration(in *ShootControllerConfiguration, out *componentconfig.ShootControllerConfiguration, s conversion.Scope) error {
	out.ConcurrentSyncs = in.ConcurrentSyncs
	out.RespectSyncPeriodOverwrite = (*bool)(unsafe.Pointer(in.RespectSyncPeriodOverwrite))
	out.RetryDuration = in.RetryDuration
	out.RetrySyncPeriod = (*v1.Duration)(unsafe.Pointer(in.RetrySyncPeriod))
	out.SyncPeriod = in.SyncPeriod
	out.WatchNamespace = (*string)(unsafe.Pointer(in.WatchNamespace))
	return nil
}

// Convert_v1alpha1_ShootControllerConfiguration_To_componentconfig_ShootControllerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ShootControllerConfiguration_To_componentconfig_ShootControllerConfiguration(in *ShootControllerConfiguration, out *componentconfig.ShootControllerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ShootControllerConfiguration_To_componentconfig_ShootControllerConfiguration(in, out, s)
}

func autoConvert_componentconfig_ShootControllerConfiguration_To_v1alpha1_ShootControllerConfiguration(in *componentconfig.ShootControllerConfiguration, out *ShootControllerConfiguration, s conversion.Scope) error {
	out.ConcurrentSyncs = in.ConcurrentSyncs
	out.RespectSyncPeriodOverwrite = (*bool)(unsafe.Pointer(in.RespectSyncPeriodOverwrite))
	out.RetryDuration = in.RetryDuration
	out.RetrySyncPeriod = (*v1.Duration)(unsafe.Pointer(in.RetrySyncPeriod))
	out.SyncPeriod = in.SyncPeriod
	out.WatchNamespace = (*string)(unsafe.Pointer(in.WatchNamespace))
	return nil
}

// Convert_componentconfig_ShootControllerConfiguration_To_v1alpha1_ShootControllerConfiguration is an autogenerated conversion function.
func Convert_componentconfig_ShootControllerConfiguration_To_v1alpha1_ShootControllerConfiguration(in *componentconfig.ShootControllerConfiguration, out *ShootControllerConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_ShootControllerConfiguration_To_v1alpha1_ShootControllerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_ShootMaintenanceControllerConfiguration_To_componentconfig_ShootMaintenanceControllerConfiguration(in *ShootMaintenanceControllerConfiguration, out *componentconfig.ShootMaintenanceControllerConfiguration, s conversion.Scope) error {
	out.ConcurrentSyncs = in.ConcurrentSyncs
	out.SyncPeriod = in.SyncPeriod
	return nil
}

// Convert_v1alpha1_ShootMaintenanceControllerConfiguration_To_componentconfig_ShootMaintenanceControllerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ShootMaintenanceControllerConfiguration_To_componentconfig_ShootMaintenanceControllerConfiguration(in *ShootMaintenanceControllerConfiguration, out *componentconfig.ShootMaintenanceControllerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ShootMaintenanceControllerConfiguration_To_componentconfig_ShootMaintenanceControllerConfiguration(in, out, s)
}

func autoConvert_componentconfig_ShootMaintenanceControllerConfiguration_To_v1alpha1_ShootMaintenanceControllerConfiguration(in *componentconfig.ShootMaintenanceControllerConfiguration, out *ShootMaintenanceControllerConfiguration, s conversion.Scope) error {
	out.ConcurrentSyncs = in.ConcurrentSyncs
	out.SyncPeriod = in.SyncPeriod
	return nil
}

// Convert_componentconfig_ShootMaintenanceControllerConfiguration_To_v1alpha1_ShootMaintenanceControllerConfiguration is an autogenerated conversion function.
func Convert_componentconfig_ShootMaintenanceControllerConfiguration_To_v1alpha1_ShootMaintenanceControllerConfiguration(in *componentconfig.ShootMaintenanceControllerConfiguration, out *ShootMaintenanceControllerConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_ShootMaintenanceControllerConfiguration_To_v1alpha1_ShootMaintenanceControllerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_ShootQuotaControllerConfiguration_To_componentconfig_ShootQuotaControllerConfiguration(in *ShootQuotaControllerConfiguration, out *componentconfig.ShootQuotaControllerConfiguration, s conversion.Scope) error {
	out.ConcurrentSyncs = in.ConcurrentSyncs
	out.SyncPeriod = in.SyncPeriod
	return nil
}

// Convert_v1alpha1_ShootQuotaControllerConfiguration_To_componentconfig_ShootQuotaControllerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ShootQuotaControllerConfiguration_To_componentconfig_ShootQuotaControllerConfiguration(in *ShootQuotaControllerConfiguration, out *componentconfig.ShootQuotaControllerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ShootQuotaControllerConfiguration_To_componentconfig_ShootQuotaControllerConfiguration(in, out, s)
}

func autoConvert_componentconfig_ShootQuotaControllerConfiguration_To_v1alpha1_ShootQuotaControllerConfiguration(in *componentconfig.ShootQuotaControllerConfiguration, out *ShootQuotaControllerConfiguration, s conversion.Scope) error {
	out.ConcurrentSyncs = in.ConcurrentSyncs
	out.SyncPeriod = in.SyncPeriod
	return nil
}

// Convert_componentconfig_ShootQuotaControllerConfiguration_To_v1alpha1_ShootQuotaControllerConfiguration is an autogenerated conversion function.
func Convert_componentconfig_ShootQuotaControllerConfiguration_To_v1alpha1_ShootQuotaControllerConfiguration(in *componentconfig.ShootQuotaControllerConfiguration, out *ShootQuotaControllerConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_ShootQuotaControllerConfiguration_To_v1alpha1_ShootQuotaControllerConfiguration(in, out, s)
}
