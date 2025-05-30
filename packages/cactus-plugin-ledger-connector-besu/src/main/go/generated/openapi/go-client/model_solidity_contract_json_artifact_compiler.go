/*
Hyperledger Cactus Plugin - Connector Besu

Can perform basic tasks on a Besu ledger

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-besu

import (
	"encoding/json"
)

// checks if the SolidityContractJsonArtifactCompiler type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SolidityContractJsonArtifactCompiler{}

// SolidityContractJsonArtifactCompiler struct for SolidityContractJsonArtifactCompiler
type SolidityContractJsonArtifactCompiler struct {
	Name *string `json:"name,omitempty"`
	Version *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SolidityContractJsonArtifactCompiler SolidityContractJsonArtifactCompiler

// NewSolidityContractJsonArtifactCompiler instantiates a new SolidityContractJsonArtifactCompiler object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSolidityContractJsonArtifactCompiler() *SolidityContractJsonArtifactCompiler {
	this := SolidityContractJsonArtifactCompiler{}
	return &this
}

// NewSolidityContractJsonArtifactCompilerWithDefaults instantiates a new SolidityContractJsonArtifactCompiler object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSolidityContractJsonArtifactCompilerWithDefaults() *SolidityContractJsonArtifactCompiler {
	this := SolidityContractJsonArtifactCompiler{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SolidityContractJsonArtifactCompiler) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolidityContractJsonArtifactCompiler) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SolidityContractJsonArtifactCompiler) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SolidityContractJsonArtifactCompiler) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SolidityContractJsonArtifactCompiler) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolidityContractJsonArtifactCompiler) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SolidityContractJsonArtifactCompiler) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SolidityContractJsonArtifactCompiler) SetVersion(v string) {
	o.Version = &v
}

func (o SolidityContractJsonArtifactCompiler) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SolidityContractJsonArtifactCompiler) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SolidityContractJsonArtifactCompiler) UnmarshalJSON(bytes []byte) (err error) {
	varSolidityContractJsonArtifactCompiler := _SolidityContractJsonArtifactCompiler{}

	if err = json.Unmarshal(bytes, &varSolidityContractJsonArtifactCompiler); err == nil {
		*o = SolidityContractJsonArtifactCompiler(varSolidityContractJsonArtifactCompiler)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSolidityContractJsonArtifactCompiler struct {
	value *SolidityContractJsonArtifactCompiler
	isSet bool
}

func (v NullableSolidityContractJsonArtifactCompiler) Get() *SolidityContractJsonArtifactCompiler {
	return v.value
}

func (v *NullableSolidityContractJsonArtifactCompiler) Set(val *SolidityContractJsonArtifactCompiler) {
	v.value = val
	v.isSet = true
}

func (v NullableSolidityContractJsonArtifactCompiler) IsSet() bool {
	return v.isSet
}

func (v *NullableSolidityContractJsonArtifactCompiler) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSolidityContractJsonArtifactCompiler(val *SolidityContractJsonArtifactCompiler) *NullableSolidityContractJsonArtifactCompiler {
	return &NullableSolidityContractJsonArtifactCompiler{value: val, isSet: true}
}

func (v NullableSolidityContractJsonArtifactCompiler) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSolidityContractJsonArtifactCompiler) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


