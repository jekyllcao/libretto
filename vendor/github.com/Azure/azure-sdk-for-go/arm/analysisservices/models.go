package analysisservices

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
)

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Deleting specifies the deleting state for provisioning state.
	Deleting ProvisioningState = "Deleting"
	// Failed specifies the failed state for provisioning state.
	Failed ProvisioningState = "Failed"
	// Paused specifies the paused state for provisioning state.
	Paused ProvisioningState = "Paused"
	// Pausing specifies the pausing state for provisioning state.
	Pausing ProvisioningState = "Pausing"
	// Preparing specifies the preparing state for provisioning state.
	Preparing ProvisioningState = "Preparing"
	// Provisioning specifies the provisioning state for provisioning state.
	Provisioning ProvisioningState = "Provisioning"
	// Resuming specifies the resuming state for provisioning state.
	Resuming ProvisioningState = "Resuming"
	// Scaling specifies the scaling state for provisioning state.
	Scaling ProvisioningState = "Scaling"
	// Succeeded specifies the succeeded state for provisioning state.
	Succeeded ProvisioningState = "Succeeded"
	// Suspended specifies the suspended state for provisioning state.
	Suspended ProvisioningState = "Suspended"
	// Suspending specifies the suspending state for provisioning state.
	Suspending ProvisioningState = "Suspending"
	// Updating specifies the updating state for provisioning state.
	Updating ProvisioningState = "Updating"
)

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// D1 specifies the d1 state for sku name.
	D1 SkuName = "D1"
	// S1 specifies the s1 state for sku name.
	S1 SkuName = "S1"
	// S2 specifies the s2 state for sku name.
	S2 SkuName = "S2"
	// S4 specifies the s4 state for sku name.
	S4 SkuName = "S4"
)

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// Development specifies the development state for sku tier.
	Development SkuTier = "Development"
	// Standard specifies the standard state for sku tier.
	Standard SkuTier = "Standard"
)

// State enumerates the values for state.
type State string

const (
	// StateDeleting specifies the state deleting state for state.
	StateDeleting State = "Deleting"
	// StateFailed specifies the state failed state for state.
	StateFailed State = "Failed"
	// StatePaused specifies the state paused state for state.
	StatePaused State = "Paused"
	// StatePausing specifies the state pausing state for state.
	StatePausing State = "Pausing"
	// StatePreparing specifies the state preparing state for state.
	StatePreparing State = "Preparing"
	// StateProvisioning specifies the state provisioning state for state.
	StateProvisioning State = "Provisioning"
	// StateResuming specifies the state resuming state for state.
	StateResuming State = "Resuming"
	// StateScaling specifies the state scaling state for state.
	StateScaling State = "Scaling"
	// StateSucceeded specifies the state succeeded state for state.
	StateSucceeded State = "Succeeded"
	// StateSuspended specifies the state suspended state for state.
	StateSuspended State = "Suspended"
	// StateSuspending specifies the state suspending state for state.
	StateSuspending State = "Suspending"
	// StateUpdating specifies the state updating state for state.
	StateUpdating State = "Updating"
)

// Resource is represents an instance of an Analysis Services resource.
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Sku      *ResourceSku        `json:"sku,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// ResourceSku is represents the SKU name and Azure pricing tier for Analysis
// Services resource.
type ResourceSku struct {
	Name SkuName `json:"name,omitempty"`
	Tier SkuTier `json:"tier,omitempty"`
}

// Server is represents an instance of an Analysis Services resource.
type Server struct {
	autorest.Response `json:"-"`
	ID                *string             `json:"id,omitempty"`
	Name              *string             `json:"name,omitempty"`
	Type              *string             `json:"type,omitempty"`
	Location          *string             `json:"location,omitempty"`
	Sku               *ResourceSku        `json:"sku,omitempty"`
	Tags              *map[string]*string `json:"tags,omitempty"`
	*ServerProperties `json:"properties,omitempty"`
}

// ServerAdministrators is an array of administrator user identities
type ServerAdministrators struct {
	Members *[]string `json:"members,omitempty"`
}

// ServerMutableProperties is an object that represents a set of mutable
// Analysis Services resource properties.
type ServerMutableProperties struct {
	AsAdministrators *ServerAdministrators `json:"asAdministrators,omitempty"`
}

// ServerProperties is properties of Analysis Services resource.
type ServerProperties struct {
	AsAdministrators  *ServerAdministrators `json:"asAdministrators,omitempty"`
	State             State                 `json:"state,omitempty"`
	ProvisioningState ProvisioningState     `json:"provisioningState,omitempty"`
	ServerFullName    *string               `json:"serverFullName,omitempty"`
}

// Servers is an array of Analysis Services resources.
type Servers struct {
	autorest.Response `json:"-"`
	Value             *[]Server `json:"value,omitempty"`
}

// ServerUpdateParameters is provision request specification
type ServerUpdateParameters struct {
	Sku                      *ResourceSku        `json:"sku,omitempty"`
	Tags                     *map[string]*string `json:"tags,omitempty"`
	*ServerMutableProperties `json:"properties,omitempty"`
}
