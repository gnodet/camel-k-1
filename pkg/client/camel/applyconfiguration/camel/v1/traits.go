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

import (
	trait "github.com/apache/camel-k/pkg/apis/camel/v1/trait"
)

// TraitsApplyConfiguration represents an declarative configuration of the Traits type for use
// with apply.
type TraitsApplyConfiguration struct {
	Affinity       *trait.AffinityTrait                    `json:"affinity,omitempty"`
	Builder        *trait.BuilderTrait                     `json:"builder,omitempty"`
	Camel          *trait.CamelTrait                       `json:"camel,omitempty"`
	Container      *trait.ContainerTrait                   `json:"container,omitempty"`
	Cron           *trait.CronTrait                        `json:"cron,omitempty"`
	Dependencies   *trait.DependenciesTrait                `json:"dependencies,omitempty"`
	Deployer       *trait.DeployerTrait                    `json:"deployer,omitempty"`
	Deployment     *trait.DeploymentTrait                  `json:"deployment,omitempty"`
	Environment    *trait.EnvironmentTrait                 `json:"environment,omitempty"`
	ErrorHandler   *trait.ErrorHandlerTrait                `json:"error-handler,omitempty"`
	GC             *trait.GCTrait                          `json:"gc,omitempty"`
	Health         *trait.HealthTrait                      `json:"health,omitempty"`
	Ingress        *trait.IngressTrait                     `json:"ingress,omitempty"`
	Istio          *trait.IstioTrait                       `json:"istio,omitempty"`
	Jolokia        *trait.JolokiaTrait                     `json:"jolokia,omitempty"`
	JVM            *trait.JVMTrait                         `json:"jvm,omitempty"`
	Kamelets       *trait.KameletsTrait                    `json:"kamelets,omitempty"`
	Knative        *trait.KnativeTrait                     `json:"knative,omitempty"`
	KnativeService *trait.KnativeServiceTrait              `json:"knative-service,omitempty"`
	Logging        *trait.LoggingTrait                     `json:"logging,omitempty"`
	Mount          *trait.MountTrait                       `json:"mount,omitempty"`
	OpenAPI        *trait.OpenAPITrait                     `json:"openapi,omitempty"`
	Owner          *trait.OwnerTrait                       `json:"owner,omitempty"`
	PDB            *trait.PDBTrait                         `json:"pdb,omitempty"`
	Platform       *trait.PlatformTrait                    `json:"platform,omitempty"`
	Pod            *trait.PodTrait                         `json:"pod,omitempty"`
	Prometheus     *trait.PrometheusTrait                  `json:"prometheus,omitempty"`
	PullSecret     *trait.PullSecretTrait                  `json:"pull-secret,omitempty"`
	Quarkus        *trait.QuarkusTrait                     `json:"quarkus,omitempty"`
	Registry       *trait.RegistryTrait                    `json:"registry,omitempty"`
	Route          *trait.RouteTrait                       `json:"route,omitempty"`
	Service        *trait.ServiceTrait                     `json:"service,omitempty"`
	ServiceBinding *trait.ServiceBindingTrait              `json:"service-binding,omitempty"`
	Toleration     *trait.TolerationTrait                  `json:"toleration,omitempty"`
	Addons         map[string]AddonTraitApplyConfiguration `json:"addons,omitempty"`
	Keda           *TraitSpecApplyConfiguration            `json:"keda,omitempty"`
	Master         *TraitSpecApplyConfiguration            `json:"master,omitempty"`
	Strimzi        *TraitSpecApplyConfiguration            `json:"strimzi,omitempty"`
	ThreeScale     *TraitSpecApplyConfiguration            `json:"3scale,omitempty"`
	Tracing        *TraitSpecApplyConfiguration            `json:"tracing,omitempty"`
}

// TraitsApplyConfiguration constructs an declarative configuration of the Traits type for use with
// apply.
func Traits() *TraitsApplyConfiguration {
	return &TraitsApplyConfiguration{}
}

// WithAffinity sets the Affinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Affinity field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithAffinity(value trait.AffinityTrait) *TraitsApplyConfiguration {
	b.Affinity = &value
	return b
}

// WithBuilder sets the Builder field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Builder field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithBuilder(value trait.BuilderTrait) *TraitsApplyConfiguration {
	b.Builder = &value
	return b
}

// WithCamel sets the Camel field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Camel field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithCamel(value trait.CamelTrait) *TraitsApplyConfiguration {
	b.Camel = &value
	return b
}

// WithContainer sets the Container field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Container field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithContainer(value trait.ContainerTrait) *TraitsApplyConfiguration {
	b.Container = &value
	return b
}

// WithCron sets the Cron field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Cron field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithCron(value trait.CronTrait) *TraitsApplyConfiguration {
	b.Cron = &value
	return b
}

// WithDependencies sets the Dependencies field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Dependencies field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithDependencies(value trait.DependenciesTrait) *TraitsApplyConfiguration {
	b.Dependencies = &value
	return b
}

// WithDeployer sets the Deployer field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Deployer field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithDeployer(value trait.DeployerTrait) *TraitsApplyConfiguration {
	b.Deployer = &value
	return b
}

// WithDeployment sets the Deployment field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Deployment field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithDeployment(value trait.DeploymentTrait) *TraitsApplyConfiguration {
	b.Deployment = &value
	return b
}

// WithEnvironment sets the Environment field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Environment field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithEnvironment(value trait.EnvironmentTrait) *TraitsApplyConfiguration {
	b.Environment = &value
	return b
}

// WithErrorHandler sets the ErrorHandler field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ErrorHandler field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithErrorHandler(value trait.ErrorHandlerTrait) *TraitsApplyConfiguration {
	b.ErrorHandler = &value
	return b
}

// WithGC sets the GC field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GC field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithGC(value trait.GCTrait) *TraitsApplyConfiguration {
	b.GC = &value
	return b
}

// WithHealth sets the Health field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Health field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithHealth(value trait.HealthTrait) *TraitsApplyConfiguration {
	b.Health = &value
	return b
}

// WithIngress sets the Ingress field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Ingress field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithIngress(value trait.IngressTrait) *TraitsApplyConfiguration {
	b.Ingress = &value
	return b
}

// WithIstio sets the Istio field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Istio field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithIstio(value trait.IstioTrait) *TraitsApplyConfiguration {
	b.Istio = &value
	return b
}

// WithJolokia sets the Jolokia field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Jolokia field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithJolokia(value trait.JolokiaTrait) *TraitsApplyConfiguration {
	b.Jolokia = &value
	return b
}

// WithJVM sets the JVM field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the JVM field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithJVM(value trait.JVMTrait) *TraitsApplyConfiguration {
	b.JVM = &value
	return b
}

// WithKamelets sets the Kamelets field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kamelets field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithKamelets(value trait.KameletsTrait) *TraitsApplyConfiguration {
	b.Kamelets = &value
	return b
}

// WithKnative sets the Knative field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Knative field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithKnative(value trait.KnativeTrait) *TraitsApplyConfiguration {
	b.Knative = &value
	return b
}

// WithKnativeService sets the KnativeService field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KnativeService field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithKnativeService(value trait.KnativeServiceTrait) *TraitsApplyConfiguration {
	b.KnativeService = &value
	return b
}

// WithLogging sets the Logging field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Logging field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithLogging(value trait.LoggingTrait) *TraitsApplyConfiguration {
	b.Logging = &value
	return b
}

// WithMount sets the Mount field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Mount field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithMount(value trait.MountTrait) *TraitsApplyConfiguration {
	b.Mount = &value
	return b
}

// WithOpenAPI sets the OpenAPI field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OpenAPI field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithOpenAPI(value trait.OpenAPITrait) *TraitsApplyConfiguration {
	b.OpenAPI = &value
	return b
}

// WithOwner sets the Owner field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Owner field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithOwner(value trait.OwnerTrait) *TraitsApplyConfiguration {
	b.Owner = &value
	return b
}

// WithPDB sets the PDB field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PDB field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithPDB(value trait.PDBTrait) *TraitsApplyConfiguration {
	b.PDB = &value
	return b
}

// WithPlatform sets the Platform field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Platform field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithPlatform(value trait.PlatformTrait) *TraitsApplyConfiguration {
	b.Platform = &value
	return b
}

// WithPod sets the Pod field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Pod field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithPod(value trait.PodTrait) *TraitsApplyConfiguration {
	b.Pod = &value
	return b
}

// WithPrometheus sets the Prometheus field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Prometheus field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithPrometheus(value trait.PrometheusTrait) *TraitsApplyConfiguration {
	b.Prometheus = &value
	return b
}

// WithPullSecret sets the PullSecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PullSecret field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithPullSecret(value trait.PullSecretTrait) *TraitsApplyConfiguration {
	b.PullSecret = &value
	return b
}

// WithQuarkus sets the Quarkus field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Quarkus field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithQuarkus(value trait.QuarkusTrait) *TraitsApplyConfiguration {
	b.Quarkus = &value
	return b
}

// WithRegistry sets the Registry field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Registry field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithRegistry(value trait.RegistryTrait) *TraitsApplyConfiguration {
	b.Registry = &value
	return b
}

// WithRoute sets the Route field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Route field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithRoute(value trait.RouteTrait) *TraitsApplyConfiguration {
	b.Route = &value
	return b
}

// WithService sets the Service field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Service field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithService(value trait.ServiceTrait) *TraitsApplyConfiguration {
	b.Service = &value
	return b
}

// WithServiceBinding sets the ServiceBinding field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServiceBinding field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithServiceBinding(value trait.ServiceBindingTrait) *TraitsApplyConfiguration {
	b.ServiceBinding = &value
	return b
}

// WithToleration sets the Toleration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Toleration field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithToleration(value trait.TolerationTrait) *TraitsApplyConfiguration {
	b.Toleration = &value
	return b
}

// WithAddons puts the entries into the Addons field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Addons field,
// overwriting an existing map entries in Addons field with the same key.
func (b *TraitsApplyConfiguration) WithAddons(entries map[string]AddonTraitApplyConfiguration) *TraitsApplyConfiguration {
	if b.Addons == nil && len(entries) > 0 {
		b.Addons = make(map[string]AddonTraitApplyConfiguration, len(entries))
	}
	for k, v := range entries {
		b.Addons[k] = v
	}
	return b
}

// WithKeda sets the Keda field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Keda field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithKeda(value *TraitSpecApplyConfiguration) *TraitsApplyConfiguration {
	b.Keda = value
	return b
}

// WithMaster sets the Master field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Master field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithMaster(value *TraitSpecApplyConfiguration) *TraitsApplyConfiguration {
	b.Master = value
	return b
}

// WithStrimzi sets the Strimzi field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Strimzi field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithStrimzi(value *TraitSpecApplyConfiguration) *TraitsApplyConfiguration {
	b.Strimzi = value
	return b
}

// WithThreeScale sets the ThreeScale field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ThreeScale field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithThreeScale(value *TraitSpecApplyConfiguration) *TraitsApplyConfiguration {
	b.ThreeScale = value
	return b
}

// WithTracing sets the Tracing field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Tracing field is set to the value of the last call.
func (b *TraitsApplyConfiguration) WithTracing(value *TraitSpecApplyConfiguration) *TraitsApplyConfiguration {
	b.Tracing = value
	return b
}
