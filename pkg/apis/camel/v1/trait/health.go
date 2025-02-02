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

package trait

// The health trait is responsible for configuring the health probes on the integration container.
//
// NOTE: this trait is enabled by default.
//
// +camel-k:trait=health.
type HealthTrait struct {
	Trait `property:",squash" json:",inline"`

	// Configures the liveness probe for the integration container (default `false`).
	LivenessProbeEnabled *bool `property:"liveness-probe-enabled" json:"livenessProbeEnabled,omitempty"`
	// Scheme to use when connecting to the liveness probe (default `HTTP`).
	LivenessScheme string `property:"liveness-scheme" json:"livenessScheme,omitempty"`
	// Number of seconds after the container has started before the liveness probe is initiated.
	LivenessInitialDelay int32 `property:"liveness-initial-delay" json:"livenessInitialDelay,omitempty"`
	// Number of seconds after which the liveness probe times out.
	LivenessTimeout int32 `property:"liveness-timeout" json:"livenessTimeout,omitempty"`
	// How often to perform the liveness probe.
	LivenessPeriod int32 `property:"liveness-period" json:"livenessPeriod,omitempty"`
	// Minimum consecutive successes for the liveness probe to be considered successful after having failed.
	LivenessSuccessThreshold int32 `property:"liveness-success-threshold" json:"livenessSuccessThreshold,omitempty"`
	// Minimum consecutive failures for the liveness probe to be considered failed after having succeeded.
	LivenessFailureThreshold int32 `property:"liveness-failure-threshold" json:"livenessFailureThreshold,omitempty"`

	// Configures the readiness probe for the integration container (default `true`).
	ReadinessProbeEnabled *bool `property:"readiness-probe-enabled" json:"readinessProbeEnabled,omitempty"`
	// Scheme to use when connecting to the readiness probe (default `HTTP`).
	ReadinessScheme string `property:"readiness-scheme" json:"readinessScheme,omitempty"`
	// Number of seconds after the container has started before the readiness probe is initiated.
	ReadinessInitialDelay int32 `property:"readiness-initial-delay" json:"readinessInitialDelay,omitempty"`
	// Number of seconds after which the readiness probe times out.
	ReadinessTimeout int32 `property:"readiness-timeout" json:"readinessTimeout,omitempty"`
	// How often to perform the readiness probe.
	ReadinessPeriod int32 `property:"readiness-period" json:"readinessPeriod,omitempty"`
	// Minimum consecutive successes for the readiness probe to be considered successful after having failed.
	ReadinessSuccessThreshold int32 `property:"readiness-success-threshold" json:"readinessSuccessThreshold,omitempty"`
	// Minimum consecutive failures for the readiness probe to be considered failed after having succeeded.
	ReadinessFailureThreshold int32 `property:"readiness-failure-threshold" json:"readinessFailureThreshold,omitempty"`

	// Configures the startup probe for the integration container (default `false`).
	StartupProbeEnabled *bool `property:"startup-probe-enabled" json:"startupProbeEnabled,omitempty"`
	// Scheme to use when connecting to the startup probe (default `HTTP`).
	StartupScheme string `property:"startup-scheme" json:"startupScheme,omitempty"`
	// Number of seconds after the container has started before the startup probe is initiated.
	StartupInitialDelay int32 `property:"startup-initial-delay" json:"startupInitialDelay,omitempty"`
	// Number of seconds after which the startup probe times out.
	StartupTimeout int32 `property:"startup-timeout" json:"startupTimeout,omitempty"`
	// How often to perform the startup probe.
	StartupPeriod int32 `property:"startup-period" json:"startupPeriod,omitempty"`
	// Minimum consecutive successes for the startup probe to be considered successful after having failed.
	StartupSuccessThreshold int32 `property:"startup-success-threshold" json:"startupSuccessThreshold,omitempty"`
	// Minimum consecutive failures for the startup probe to be considered failed after having succeeded.
	StartupFailureThreshold int32 `property:"startup-failure-threshold" json:"startupFailureThreshold,omitempty"`
}
