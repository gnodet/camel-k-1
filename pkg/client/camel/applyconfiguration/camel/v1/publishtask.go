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

// PublishTaskApplyConfiguration represents an declarative configuration of the PublishTask type for use
// with apply.
type PublishTaskApplyConfiguration struct {
	ContextDir *string                         `json:"contextDir,omitempty"`
	BaseImage  *string                         `json:"baseImage,omitempty"`
	Image      *string                         `json:"image,omitempty"`
	Registry   *RegistrySpecApplyConfiguration `json:"registry,omitempty"`
}

// PublishTaskApplyConfiguration constructs an declarative configuration of the PublishTask type for use with
// apply.
func PublishTask() *PublishTaskApplyConfiguration {
	return &PublishTaskApplyConfiguration{}
}

// WithContextDir sets the ContextDir field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ContextDir field is set to the value of the last call.
func (b *PublishTaskApplyConfiguration) WithContextDir(value string) *PublishTaskApplyConfiguration {
	b.ContextDir = &value
	return b
}

// WithBaseImage sets the BaseImage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BaseImage field is set to the value of the last call.
func (b *PublishTaskApplyConfiguration) WithBaseImage(value string) *PublishTaskApplyConfiguration {
	b.BaseImage = &value
	return b
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *PublishTaskApplyConfiguration) WithImage(value string) *PublishTaskApplyConfiguration {
	b.Image = &value
	return b
}

// WithRegistry sets the Registry field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Registry field is set to the value of the last call.
func (b *PublishTaskApplyConfiguration) WithRegistry(value *RegistrySpecApplyConfiguration) *PublishTaskApplyConfiguration {
	b.Registry = value
	return b
}