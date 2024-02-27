// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AWSClassicLoadBalancerParametersApplyConfiguration represents an declarative configuration of the AWSClassicLoadBalancerParameters type for use
// with apply.
type AWSClassicLoadBalancerParametersApplyConfiguration struct {
	ConnectionIdleTimeout *v1.Duration `json:"connectionIdleTimeout,omitempty"`
}

// AWSClassicLoadBalancerParametersApplyConfiguration constructs an declarative configuration of the AWSClassicLoadBalancerParameters type for use with
// apply.
func AWSClassicLoadBalancerParameters() *AWSClassicLoadBalancerParametersApplyConfiguration {
	return &AWSClassicLoadBalancerParametersApplyConfiguration{}
}

// WithConnectionIdleTimeout sets the ConnectionIdleTimeout field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ConnectionIdleTimeout field is set to the value of the last call.
func (b *AWSClassicLoadBalancerParametersApplyConfiguration) WithConnectionIdleTimeout(value v1.Duration) *AWSClassicLoadBalancerParametersApplyConfiguration {
	b.ConnectionIdleTimeout = &value
	return b
}