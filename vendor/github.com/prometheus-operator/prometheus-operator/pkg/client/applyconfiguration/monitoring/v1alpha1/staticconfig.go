// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	monitoringv1alpha1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1alpha1"
)

// StaticConfigApplyConfiguration represents a declarative configuration of the StaticConfig type for use
// with apply.
type StaticConfigApplyConfiguration struct {
	Targets []monitoringv1alpha1.Target `json:"targets,omitempty"`
	Labels  map[string]string           `json:"labels,omitempty"`
}

// StaticConfigApplyConfiguration constructs a declarative configuration of the StaticConfig type for use with
// apply.
func StaticConfig() *StaticConfigApplyConfiguration {
	return &StaticConfigApplyConfiguration{}
}

// WithTargets adds the given value to the Targets field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Targets field.
func (b *StaticConfigApplyConfiguration) WithTargets(values ...monitoringv1alpha1.Target) *StaticConfigApplyConfiguration {
	for i := range values {
		b.Targets = append(b.Targets, values[i])
	}
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *StaticConfigApplyConfiguration) WithLabels(entries map[string]string) *StaticConfigApplyConfiguration {
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}
