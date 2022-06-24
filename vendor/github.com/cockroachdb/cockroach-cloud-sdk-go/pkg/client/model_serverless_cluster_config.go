// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// CockroachDB Cloud API
// API version: 2022-03-31

package client

import (
	"encoding/json"
)

// ServerlessClusterConfig struct for ServerlessClusterConfig.
type ServerlessClusterConfig struct {
	// Spend limit in US cents.
	SpendLimit int32 `json:"spend_limit"`
	// Used to build a connection string.
	RoutingId            string `json:"routing_id"`
	AdditionalProperties map[string]interface{}
}

type serverlessClusterConfig ServerlessClusterConfig

// NewServerlessClusterConfig instantiates a new ServerlessClusterConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerlessClusterConfig(spendLimit int32, routingId string) *ServerlessClusterConfig {
	p := ServerlessClusterConfig{}
	p.SpendLimit = spendLimit
	p.RoutingId = routingId
	return &p
}

// NewServerlessClusterConfigWithDefaults instantiates a new ServerlessClusterConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerlessClusterConfigWithDefaults() *ServerlessClusterConfig {
	p := ServerlessClusterConfig{}
	return &p
}

// GetSpendLimit returns the SpendLimit field value.
func (o *ServerlessClusterConfig) GetSpendLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SpendLimit
}

// SetSpendLimit sets field value.
func (o *ServerlessClusterConfig) SetSpendLimit(v int32) {
	o.SpendLimit = v
}

// GetRoutingId returns the RoutingId field value.
func (o *ServerlessClusterConfig) GetRoutingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoutingId
}

// SetRoutingId sets field value.
func (o *ServerlessClusterConfig) SetRoutingId(v string) {
	o.RoutingId = v
}

func (o ServerlessClusterConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["spend_limit"] = o.SpendLimit
	}
	if true {
		toSerialize["routing_id"] = o.RoutingId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServerlessClusterConfig) UnmarshalJSON(bytes []byte) (err error) {
	varServerlessClusterConfig := serverlessClusterConfig{}

	if err = json.Unmarshal(bytes, &varServerlessClusterConfig); err == nil {
		*o = ServerlessClusterConfig(varServerlessClusterConfig)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "spend_limit")
		delete(additionalProperties, "routing_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
