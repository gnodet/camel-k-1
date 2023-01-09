/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// KanikoTaskCacheApplyConfiguration represents an declarative configuration of the KanikoTaskCache type for use
// with apply.
type KanikoTaskCacheApplyConfiguration struct {
	Enabled               *bool   `json:"enabled,omitempty"`
	PersistentVolumeClaim *string `json:"persistentVolumeClaim,omitempty"`
}

// KanikoTaskCacheApplyConfiguration constructs an declarative configuration of the KanikoTaskCache type for use with
// apply.
func KanikoTaskCache() *KanikoTaskCacheApplyConfiguration {
	return &KanikoTaskCacheApplyConfiguration{}
}

// WithEnabled sets the Enabled field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Enabled field is set to the value of the last call.
func (b *KanikoTaskCacheApplyConfiguration) WithEnabled(value bool) *KanikoTaskCacheApplyConfiguration {
	b.Enabled = &value
	return b
}

// WithPersistentVolumeClaim sets the PersistentVolumeClaim field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PersistentVolumeClaim field is set to the value of the last call.
func (b *KanikoTaskCacheApplyConfiguration) WithPersistentVolumeClaim(value string) *KanikoTaskCacheApplyConfiguration {
	b.PersistentVolumeClaim = &value
	return b
}
